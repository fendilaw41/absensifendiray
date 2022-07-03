package aktifitas

import (
	"path/filepath"
	"strconv"

	"github.com/fendilaw41/absensifendiray/app/src/aktifitasImage"
	"github.com/fendilaw41/absensifendiray/app/src/user"
	"github.com/fendilaw41/absensifendiray/config/action"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAll() ([]Aktifitas, error)
	FindAllPaginate(res *gin.Context) ([]Aktifitas, int, int, int64, error)
	RiwayatAktifitas(res *gin.Context) ([]user.User, int, int, int64, error)
	FindId(Id int) (Aktifitas, error)
	Save(res *gin.Context, userId int, req AktifitasRequest) (Aktifitas, error)
	UpdateAktifitasService(res *gin.Context, Id int, req AktifitasRequestPUT, userId int) (Aktifitas, error)
	SrcDeleteAktifitas(Id int) (Aktifitas, error)
	CountCheckIN(tgl string, userId int) int64
}

type service struct {
	repository Repository
}

func AktifitasService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Aktifitas, error) {
	var Aktifitass []Aktifitas
	Aktifitass, err := s.repository.GetAll()
	return Aktifitass, err
}

func (s *service) FindAllPaginate(res *gin.Context) ([]Aktifitas, int, int, int64, error) {
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

func (s *service) RiwayatAktifitas(res *gin.Context) ([]user.User, int, int, int64, error) {
	query := map[string]string{
		"page_size":    res.Query("page_size"),
		"page":         res.Query("page"),
		"tanggal_dari": res.Query("tanggal_dari"),
		"tanggal_ke":   res.Query("tanggal_ke"),
	}
	pageSize, err := strconv.Atoi(query["page_size"])
	if err != nil {
		pageSize = 5
	}
	page, err := strconv.Atoi(query["page"])
	if err != nil {
		page = 1
	}
	tanggal_dari := action.FormatDate(query["tanggal_dari"])
	tanggal_ke := action.FormatDate(query["tanggal_ke"])
	aktifitas, count, err := s.repository.RiwayatAktifitas(page, pageSize, tanggal_dari, tanggal_ke)

	return aktifitas, page, pageSize, count, err
}

func (s *service) FindId(Id int) (Aktifitas, error) {
	Aktifitas, err := s.repository.GetById(Id)
	return Aktifitas, err
}

func (s *service) Save(res *gin.Context, userId int, req AktifitasRequest) (Aktifitas, error) {
	Aktifitas := Aktifitas{
		Name:        req.Name,
		AssignedId:  req.AssignId,
		Subject:     req.Subject,
		Description: req.Description,
		UserId:      userId,
		CreatedBy:   action.AuthId(res),
	}
	AktifitasIdx, err := s.repository.Store(Aktifitas)
	form, _ := res.MultipartForm()
	files := form.File["picture[]"]
	var aktifitasImages = make([]aktifitasImage.AktifitasImage, len(files))
	for index, file := range files {
		fileName := action.RandomString(16) + ".png"
		filePath, _ := action.HandleImageAktifitas(res, fileName, file)
		fileSize := (uint)(file.Size)
		aktifitasImages[index] = aktifitasImage.AktifitasImage{AktifitasId: AktifitasIdx.ID, Filename: fileName, FilePath: string(filepath.Separator) + filePath, OriginalName: file.Filename, FileSize: fileSize}
	}
	s.repository.StoreImages(AktifitasIdx.ID, aktifitasImages)
	return AktifitasIdx, err
}

func (s *service) UpdateAktifitasService(res *gin.Context, Id int, req AktifitasRequestPUT, userId int) (Aktifitas, error) {
	Aktifitas, _ := s.repository.GetDetail(Id)
	Aktifitas.Name = req.Name
	Aktifitas.AssignedId = req.AssignId
	Aktifitas.UserId = userId
	Aktifitas.Subject = req.Subject
	Aktifitas.Description = req.Description
	Aktifitas.UpdatedBy = action.AuthId(res)
	updateAktifitas, err := s.repository.UpdateAktifitas(Aktifitas)
	form, _ := res.MultipartForm()
	files := form.File["picture[]"]
	var productImages = make([]aktifitasImage.AktifitasImage, len(files))
	for index, file := range files {
		s.repository.DeleteImages(updateAktifitas.ID)
		fileName := action.RandomString(16) + ".png"
		filePath, _ := action.HandleImageAktifitas(res, fileName, file)
		fileSize := (uint)(file.Size)
		productImages[index] = aktifitasImage.AktifitasImage{AktifitasId: updateAktifitas.ID, Filename: fileName, FilePath: string(filepath.Separator) + filePath, OriginalName: file.Filename, FileSize: fileSize}
	}
	s.repository.StoreImages(updateAktifitas.ID, productImages)

	return updateAktifitas, err
}

func (s *service) SrcDeleteAktifitas(Id int) (Aktifitas, error) {
	Aktifitas, err := s.repository.DeleteAktifitas(Id)
	return Aktifitas, err
}

func (s *service) CountCheckIN(tgl string, userId int) int64 {
	count := s.repository.CountCheckIN(tgl, userId)
	return count
}
