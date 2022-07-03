package user

import (
	"strconv"
	"strings"
	"time"

	"github.com/fendilaw41/absensifendiray/config/action"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

type Service interface {
	FindAllPaginate(res *gin.Context) ([]User, int, int, int64, error)
	FindAll() ([]User, error)
	FindId(Id int) (User, error)
	Save(userRequest UserRequest, res *gin.Context) error
	UpdateUserService(Id int, req UserRequestPUT, res *gin.Context) error
	SrcDeleteUser(Id int, res *gin.Context) error
}

type service struct {
	repository Repository
}

func UserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllPaginate(res *gin.Context) ([]User, int, int, int64, error) {
	query := map[string]string{
		"page_size": res.Query("page_size"),
		"page":      res.Query("page"),
		"name":      res.Query("name"),
	}
	pageSize, err := strconv.Atoi(query["page_size"])
	if err != nil {
		pageSize = 10
	}
	page, err := strconv.Atoi(query["page"])
	if err != nil {
		page = 1
	}
	categories, count, err := s.repository.GetAllPaginate(page, pageSize, query["name"])

	return categories, page, pageSize, count, err
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.GetAll()
	return users, err
}

func (s *service) FindId(Id int) (User, error) {
	user, err := s.repository.GetDetail(Id)
	return user, err
}

func (s *service) Save(req UserRequest, res *gin.Context) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	authId := action.AuthId(res)
	user := User{
		Name:          req.Name,
		Email:         req.Email,
		Password:      string(bytes),
		BirthDate:     datatypes.Date(time.Now()),
		FirstName:     req.Name,
		LastName:      req.LastName,
		Fullname:      req.Name + " " + req.LastName,
		DepartementId: req.DepartementId,
		DivisiId:      req.DivisiId,
		CreatedBy:     authId,
	}
	users, err := s.repository.Store(user)
	form, _ := res.MultipartForm()
	for k, v := range form.Value {
		if strings.HasPrefix(k, "roles[") {
			name := v[0]
			rolesName, _ := s.repository.FirstRoleName(name)
			role := UserRole{
				UserId: users.ID,
				RoleId: int(rolesName.Id),
			}
			s.repository.StoreUserRole(role)
		}
	}

	return err
}

func (s *service) UpdateUserService(Id int, req UserRequestPUT, res *gin.Context) error {
	authId := action.AuthId(res)
	user, _ := s.repository.GetDetail(Id)
	user.Name = req.Name
	user.FirstName = req.Name
	user.LastName = req.LastName
	user.Fullname = req.Name + " " + req.LastName
	user.DepartementId = req.DepartementId
	user.DivisiId = req.DivisiId
	user.UpdatedBy = authId

	users, err := s.repository.UpdateUser(user)
	s.repository.DeleteUserRole(Id)
	form, _ := res.MultipartForm()
	for k, v := range form.Value {
		if strings.HasPrefix(k, "roles[") {
			name := v[0]
			rolesName, _ := s.repository.FirstRoleName(name)
			role := UserRole{
				UserId: users.ID,
				RoleId: int(rolesName.Id),
			}
			s.repository.StoreUserRole(role)
		}
	}
	return err
}

func (s *service) DeleteUserRole(Id int, res *gin.Context) error {
	userId, _ := s.repository.GetDetail(Id)
	_, err := s.repository.DeleteUserRole(userId.ID)
	return err
}

func (s *service) SrcDeleteUser(Id int, res *gin.Context) error {
	userId, _ := s.repository.GetDetail(Id)
	authId := action.AuthId(res)
	userId.UpdatedBy = authId
	_, err := s.repository.DeleteUser(userId.ID)
	return err
}
