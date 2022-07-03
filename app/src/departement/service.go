package departement

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAll() ([]Departement, error)
	FindAllPaginate(res *gin.Context) ([]Departement, int, int, int64, error)
	FindId(Id int) (Departement, error)
	SaveDepartementSrv(dp DepartementRequest) (Departement, error)
	UpdateDepartementSrv(Id int, dp DepartementRequest) (Departement, error)
	SrcDeleteDepartement(Id int) (Departement, error)
}

type service struct {
	repository RepositoryDepartement
}

func DepartementService(repository RepositoryDepartement) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Departement, error) {
	var categories []Departement
	categories, err := s.repository.GetAll()
	return categories, err
}

func (s *service) FindAllPaginate(res *gin.Context) ([]Departement, int, int, int64, error) {
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
	categories, count, err := s.repository.GetAllPaginate(page, pageSize, query["name"])

	return categories, page, pageSize, count, err
}

func (s *service) FindId(Id int) (Departement, error) {
	Departement, err := s.repository.GetDetail(Id)
	return Departement, err
}

func (s *service) SaveDepartementSrv(dp DepartementRequest) (Departement, error) {
	Departement := Departement{
		Name: dp.Name,
	}
	saveDepartement, err := s.repository.StoreDepartementRepo(Departement)
	return saveDepartement, err
}

func (s *service) UpdateDepartementSrv(Id int, DepartementRequest DepartementRequest) (Departement, error) {
	Departement, _ := s.repository.GetDetail(Id)
	Departement.Name = DepartementRequest.Name

	updateDepartement, err := s.repository.UpdateDepartementRepo(Departement)
	return updateDepartement, err
}

func (s *service) SrcDeleteDepartement(Id int) (Departement, error) {
	Departement, err := s.repository.DeleteDepartementRepo(Id)
	return Departement, err
}
