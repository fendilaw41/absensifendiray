package absensi

import (
	"strconv"
	"time"

	"github.com/fendilaw41/absensifendiray/app/src/user"
	"github.com/fendilaw41/absensifendiray/config/action"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Service interface {
	FindAllPaginate(res *gin.Context) ([]Absensi, int, int, int64, error)
	FindId(Id int) (Absensi, error)
	FindAll() ([]Absensi, error)
	CheckIn(req AbsensiRequest, res *gin.Context) (Absensi, error)
	CheckOut(req AbsensiRequest, res *gin.Context) (Absensi, error)
	UpdateAbsensiService(Id int, AbsensiRequest AbsensiRequest, res *gin.Context) (Absensi, error)
	SrcDeleteAbsensi(Id int) (Absensi, error)
	CountCheckIN(tgl string, userId int) int64
	CountCheckOut(tgl string, userId int) int64
}

type service struct {
	repository Repository
}

func AbsensiService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllPaginate(res *gin.Context) ([]Absensi, int, int, int64, error) {
	query := map[string]string{
		"page_size": res.Query("page_size"),
		"page":      res.Query("page"),
		"status":    res.Query("status"),
	}

	pageSize, err := strconv.Atoi(query["page_size"])
	if err != nil {
		pageSize = 5
	}
	page, err := strconv.Atoi(query["page"])
	if err != nil {
		page = 1
	}
	absensis, count, errAbsensi := s.repository.GetAllPaginate(page, pageSize, query["status"])

	return absensis, page, pageSize, count, errAbsensi
}

func (s *service) FindAll() ([]Absensi, error) {
	Absensis, err := s.repository.GetAll()
	return Absensis, err
}

func (s *service) FindId(Id int) (Absensi, error) {
	Absensi, err := s.repository.GetDetail(Id)
	return Absensi, err
}

func (s *service) CountCheckIN(tgl string, userId int) int64 {
	count := s.repository.CountCheckIN(tgl, userId)
	return count
}
func (s *service) CountCheckOut(tgl string, userId int) int64 {
	count := s.repository.CountCheckOut(tgl, userId)
	return count
}

func (s *service) CheckIn(pr AbsensiRequest, res *gin.Context) (Absensi, error) {
	authId := action.AuthId(res)
	user, _ := s.repository.GetByUserId(authId)
	if file, _ := res.FormFile("picture"); file != nil { // jika file di isi
		user, _ := s.repository.GetByUserId(authId)
		reqFile := Absensi{
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			FullName:     user.Fullname,
			TanggalAbsen: datatypes.Date(time.Now()),
			JamAbsen:     datatypes.NewTime(int(pr.Hours), int(pr.Minutes), int(pr.Seconds), 0),
			CheckAbsen:   "Check-IN",
			Status:       pr.Status,
			Picture:      file.Filename,
			UserId:       action.AuthId(res),
			CreatedBy:    authId,
		}
		saveAbsensi, err := s.repository.Store(reqFile)
		return saveAbsensi, err
	}
	// jika file di kosongkan
	req := RequestCheckIn(pr, authId, user, res)
	saveAbsensi, err := s.repository.Store(req)
	return saveAbsensi, err
}

func (s *service) CheckOut(pr AbsensiRequest, res *gin.Context) (Absensi, error) {
	authId := action.AuthId(res)
	user, _ := s.repository.GetByUserId(authId)
	if file, _ := res.FormFile("picture"); file != nil { // jika file di isi
		user, _ := s.repository.GetByUserId(authId)
		reqFile := Absensi{
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			FullName:     user.Fullname,
			TanggalAbsen: datatypes.Date(time.Now()),
			JamAbsen:     datatypes.NewTime(int(pr.Hours), int(pr.Minutes), int(pr.Seconds), 0),
			CheckAbsen:   "Check-OUT",
			Status:       pr.Status,
			Picture:      file.Filename,
			UserId:       action.AuthId(res),
			CreatedBy:    authId,
		}
		saveAbsensi, err := s.repository.Store(reqFile)
		return saveAbsensi, err
	}
	// jika file di kosongkan
	req := RequestCheckOut(pr, authId, user, res)
	saveAbsensi, err := s.repository.Store(req)
	return saveAbsensi, err
}

func (s *service) UpdateAbsensiService(Id int, pr AbsensiRequest, res *gin.Context) (Absensi, error) {
	if file, _ := res.FormFile("picture"); file != nil { // jika file di isi
		authId := action.AuthId(res)
		user, _ := s.repository.GetByUserId(authId)
		reqPicture, _ := s.repository.GetDetail(Id)
		reqPicture.FirstName = user.FirstName
		reqPicture.LastName = user.LastName
		reqPicture.FullName = user.Fullname
		reqPicture.TanggalAbsen = datatypes.Date(time.Now())
		reqPicture.JamAbsen = datatypes.NewTime(int(pr.Hours), int(pr.Minutes), int(pr.Seconds), 0)
		reqPicture.Status = pr.Status
		reqPicture.Picture = file.Filename
		reqPicture.UserId = authId
		updateAbsensi, err := s.repository.UpdateAbsensi(Id, reqPicture)
		return updateAbsensi, err
	}
	// jika file dikosongkan
	authId := action.AuthId(res)
	user, _ := s.repository.GetByUserId(authId)
	Absensi, _ := s.repository.GetDetail(Id)
	Absensi.FirstName = user.FirstName
	Absensi.LastName = user.LastName
	Absensi.FullName = user.Fullname
	Absensi.TanggalAbsen = datatypes.Date(time.Now())
	Absensi.JamAbsen = datatypes.NewTime(int(pr.Hours), int(pr.Minutes), int(pr.Seconds), 0)
	Absensi.Status = pr.Status
	Absensi.UserId = authId

	updateAbsensi, err := s.repository.UpdateAbsensi(Id, Absensi)
	return updateAbsensi, err
}

func (s *service) SrcDeleteAbsensi(Id int) (Absensi, error) {
	Absensi, err := s.repository.DeleteAbsensiRepo(Id)
	return Absensi, err
}

func RequestCheckIn(c AbsensiRequest, authId int, u user.User, res *gin.Context) Absensi {
	return Absensi{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		FullName:     u.Fullname,
		UserId:       action.AuthId(res),
		TanggalAbsen: datatypes.Date(time.Now()),
		JamAbsen:     datatypes.NewTime(int(c.Hours), int(c.Minutes), int(c.Seconds), 0),
		CheckAbsen:   "Check-IN",
		CreatedBy:    authId,
	}
}

func RequestCheckOut(c AbsensiRequest, authId int, u user.User, res *gin.Context) Absensi {
	return Absensi{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		FullName:     u.Fullname,
		UserId:       action.AuthId(res),
		TanggalAbsen: datatypes.Date(time.Now()),
		JamAbsen:     datatypes.NewTime(int(c.Hours), int(c.Minutes), int(c.Seconds), 0),
		CheckAbsen:   "Check-Out",
		CreatedBy:    authId,
	}
}
