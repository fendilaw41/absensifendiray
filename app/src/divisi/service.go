package divisi

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAll() ([]Divisi, error)
	FindAllPaginate(res *gin.Context) ([]Divisi, int, int, int64, error)
	FindId(Id int) (Divisi, error)
	Save(DivisiRequest DivisiRequest) (Divisi, error)
	UpdateDivisiService(Id int, DivisiRequest DivisiRequest) (Divisi, error)
	SrcDeleteDivisi(Id int) (Divisi, error)
}

type service struct {
	repository Repository
}

func DivisiService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Divisi, error) {
	var Divisis []Divisi
	Divisis, err := s.repository.GetAll()
	return Divisis, err
}

func (s *service) FindAllPaginate(res *gin.Context) ([]Divisi, int, int, int64, error) {
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

func (s *service) FindId(Id int) (Divisi, error) {
	Divisi, err := s.repository.GetDetail(Id)
	return Divisi, err
}

func (s *service) Save(dr DivisiRequest) (Divisi, error) {
	Divisi := Divisi{
		Name:      dr.Name,
		CreatedBy: 1,
	}
	saveDivisi, err := s.repository.Store(Divisi)
	return saveDivisi, err
}

func (s *service) UpdateDivisiService(Id int, DivisiRequest DivisiRequest) (Divisi, error) {
	Divisi, _ := s.repository.GetDetail(Id)
	Divisi.Name = DivisiRequest.Name

	updateDivisi, err := s.repository.UpdateDivisi(Divisi)
	return updateDivisi, err
}

func (s *service) SrcDeleteDivisi(Id int) (Divisi, error) {
	Divisi, err := s.repository.DeleteDivisi(Id)
	return Divisi, err
}
