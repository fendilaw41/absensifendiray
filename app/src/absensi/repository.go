package absensi

import (
	"github.com/fendilaw41/absensifendiray/app/src/user"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllPaginate(page, page_size int, status string) ([]Absensi, int64, error)
	GetDetail(Id int) (Absensi, error)
	GetById(Id int) (Absensi, error)
	GetByUserId(Id int) (user.User, error)
	GetAll() ([]Absensi, error)
	Store(Absensi Absensi) (Absensi, error)
	UpdateAbsensi(Id int, Absensi Absensi) (Absensi, error)
	DeleteAbsensiRepo(Id int) (Absensi, error)
	CountCheckIN(tgl string, userId int) int64
	CountCheckOut(tgl string, userId int) int64
}

type repository struct {
	db *gorm.DB
}

func AbsensiRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllPaginate(page, page_size int, status string) ([]Absensi, int64, error) {
	var Absensi []Absensi
	var count int64
	r.db.Model(&Absensi).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&Absensi)
	// filter data string
	tx := r.db.Model(&Absensi).Where("status LIKE ?", "%"+status+"%")

	err := tx.Model(&Absensi).Preload("User").Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&Absensi).Error

	return Absensi, count, err
}

func (r *repository) GetAll() ([]Absensi, error) {
	var Absensi []Absensi
	tx := r.db.Begin()
	tx.Model(&Absensi).Preload("User").Order("id asc").Find(&Absensi)
	err := tx.Commit().Error
	return Absensi, err
}

func (r *repository) GetDetail(Id int) (Absensi, error) {
	var Absensi Absensi
	err := r.db.Where("id = ?", Id).Preload("User").First(&Absensi).Error
	return Absensi, err
}

func (r *repository) GetById(Id int) (Absensi, error) {
	var Absensi Absensi
	err := r.db.Where("id = ?", Id).Preload("User").First(&Absensi).Error
	return Absensi, err
}

func (r *repository) GetByUserId(Id int) (user.User, error) {
	var user user.User
	err := r.db.Where("id = ?", Id).First(&user).Error
	return user, err
}

func (r *repository) Store(Absensi Absensi) (Absensi, error) {
	err := r.db.Create(&Absensi).Error
	return Absensi, err
}

func (r *repository) UpdateAbsensi(Id int, Absensi Absensi) (Absensi, error) {
	err := r.db.Where("id = ?", Id).Save(&Absensi).Error
	return Absensi, err
}

func (r *repository) DeleteAbsensiRepo(Id int) (Absensi, error) {
	var Absensi Absensi
	err := r.db.Where("id = ?", Id).Preload("User").First(&Absensi).Error
	r.db.Delete(&Absensi)
	return Absensi, err
}

func (r *repository) CountCheckIN(tgl string, userId int) int64 {
	var count int64
	var p Absensi
	r.db.Model(&p).Where("tanggal_absen = ?", tgl).Where("check_absen = ?", "Check-IN").Where("user_id = ?", userId).Count(&count)
	return count
}

func (r *repository) CountCheckOut(tgl string, userId int) int64 {
	var count int64
	var p Absensi
	r.db.Model(&p).Where("tanggal_absen = ?", tgl).Where("check_absen = ?", "Check-OUT").Where("user_id = ?", userId).Count(&count)
	return count
}
