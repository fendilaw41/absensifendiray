package permission

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAllPaginate(res *gin.Context) ([]Permission, int, int, int64, error)
	FindId(Id int) (Permission, error)
	Save(pr PermissionRequest, res *gin.Context) error
	Update(Id int, pr PermissionUpdateRequest, res *gin.Context) error
	Delete(Id int) error
}

type service struct {
	repository Repository
}

func PermissionService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllPaginate(res *gin.Context) ([]Permission, int, int, int64, error) {
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
	p, count, err := s.repository.GetAllPaginate(page, pageSize, query["name"])
	return p, page, pageSize, count, err
}

func (s *service) FindId(Id int) (Permission, error) {
	p, err := s.repository.GetDetail(Id)
	return p, err
}

func (s *service) Save(pr PermissionRequest, res *gin.Context) error {
	p := Permission{
		Name:        pr.Name,
		Description: pr.Description,
	}
	_, err := s.repository.Store(p)

	return err
}

func (s *service) Update(Id int, pr PermissionUpdateRequest, res *gin.Context) error {
	p, _ := s.repository.GetDetail(Id)
	p.Name = pr.Name
	p.Description = pr.Description
	_, err := s.repository.UpdatePermission(p)
	return err
}

func (s *service) Delete(Id int) error {
	_, err := s.repository.DeletePermission(Id)
	return err
}
