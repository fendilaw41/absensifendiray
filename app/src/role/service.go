package role

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAllPaginate(res *gin.Context) ([]Role, int, int, int64, error)
	FindAll() ([]Role, error)
	FindId(Id int) (Role, error)
	SaveRole(roleRequest RoleRequest, res *gin.Context) error
	UpdateRoleService(Id int, roleRequest RoleUpdateRequest, res *gin.Context) error
	SrcDeleteRole(Id int) error
}

type service struct {
	repository Repository
}

func RoleService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Role, error) {
	users, err := s.repository.GetAll()
	return users, err
}

func (s *service) FindAllPaginate(res *gin.Context) ([]Role, int, int, int64, error) {
	query := map[string]string{
		"page_size": res.Query("page_size"),
		"page":      res.Query("page"),
		"name":      res.Query("name"),
	}
	pageSize, err := strconv.Atoi(query["page_size"])
	if err != nil {
		pageSize = 5
	}
	page, err := strconv.Atoi(query["page"])
	if err != nil {
		page = 1
	}
	role, count, err := s.repository.GetAllPaginate(page, pageSize, query["name"])
	return role, page, pageSize, count, err
}

func (s *service) FindId(Id int) (Role, error) {
	role, err := s.repository.GetDetail(Id)
	return role, err
}

func (s *service) SaveRole(roleRequest RoleRequest, res *gin.Context) error {
	role := Role{
		Name:        roleRequest.Name,
		Description: "test",
	}
	roles, err := s.repository.Store(role)
	form, _ := res.MultipartForm()
	for k, v := range form.Value {
		if strings.HasPrefix(k, "permission[") {
			name := v[0]
			newPermission, _ := s.repository.FirstPermissionName(name)
			permission := RolePermission{
				RoleId:       roles.Id,
				PermissionId: uint(newPermission.Id),
			}
			s.repository.StoreRolePermission(permission)
		}
	}
	return err
}

func (s *service) UpdateRoleService(Id int, roleRequest RoleUpdateRequest, res *gin.Context) error {
	role, _ := s.repository.GetDetail(Id)
	role.Name = roleRequest.Name
	role.Description = roleRequest.Name
	roles, err := s.repository.UpdateRole(role)
	s.repository.DeleteRolePermission(int(roles.Id))
	form, _ := res.MultipartForm()
	for k, v := range form.Value {
		if strings.HasPrefix(k, "permission[") {
			name := v[0]
			newPermission, _ := s.repository.FirstPermissionName(name)
			permission := RolePermission{
				RoleId:       roles.Id,
				PermissionId: uint(newPermission.Id),
			}
			s.repository.StoreRolePermission(permission)
		}
	}
	return err
}

func (s *service) SrcDeleteRole(Id int) error {
	roles, err := s.repository.DeleteRole(Id)
	s.repository.DeleteRolePermission(int(roles.ID))
	return err
}
