package aktifitas

import (
	"github.com/fendilaw41/absensifendiray/app/src/absensi"
	"github.com/fendilaw41/absensifendiray/app/src/aktifitasImage"
	"github.com/fendilaw41/absensifendiray/app/src/user"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Aktifitas, error)
	GetAllPaginate(page, page_size int, name string) ([]Aktifitas, int64, error)
	RiwayatAktifitas(page, page_size int, tanggal_dari, tanggal_ke datatypes.Date) ([]user.User, int64, error)
	GetDetail(Id int) (Aktifitas, error)
	GetById(Id int) (Aktifitas, error)
	Store(Aktifitas Aktifitas) (Aktifitas, error)
	StoreImages(aktifitasId int, images []aktifitasImage.AktifitasImage) ([]aktifitasImage.AktifitasImage, error)
	UpdateAktifitas(Aktifitas Aktifitas) (Aktifitas, error)
	DeleteImages(Id int) (aktifitasImage.AktifitasImage, error)
	DeleteAktifitas(Id int) (Aktifitas, error)
	CountCheckIN(tgl string, userId int) int64
}

type repository struct {
	db *gorm.DB
}

func AktifitasRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Aktifitas, error) {
	var Aktifitas []Aktifitas
	err := r.db.Preload("User").Preload("Assigned").Find(&Aktifitas).Error
	return Aktifitas, err
}

func (r *repository) RiwayatAktifitas(page, page_size int, tanggal_dari, tanggal_ke datatypes.Date) ([]user.User, int64, error) {
	var users []user.User
	var akt []Aktifitas
	var aktId []int
	var count int64
	r.db.Model(&users).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&users)
	// filter data string
	r.db.Where("created_at BETWEEN ? AND ?", tanggal_dari, tanggal_ke).Find(&akt)
	for _, v := range akt {
		aktId = append(aktId, v.ID)
	}
	tx := r.db.Where("id IN ?", aktId)

	err := tx.Model(&users).Preload("Departement").Preload("Divisi").Preload("Aktifitas").Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&users).Error

	return users, count, err
}

func (r *repository) GetAllPaginate(page, page_size int, name string) ([]Aktifitas, int64, error) {
	var Aktifitas []Aktifitas
	var count int64
	r.db.Model(&Aktifitas).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&Aktifitas)
	// filter data string
	tx := r.db.Model(&Aktifitas).Where("name LIKE ?", "%"+name+"%")

	err := tx.Model(&Aktifitas).Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&Aktifitas).Error

	return Aktifitas, count, err
}

func (r *repository) GetDetail(Id int) (Aktifitas, error) {
	var Aktifitas Aktifitas
	err := r.db.Find(&Aktifitas, Id).Error
	return Aktifitas, err
}

func (r *repository) GetById(Id int) (Aktifitas, error) {
	var akt Aktifitas
	err := r.db.Where("id = ?", Id).First(&akt).Error
	return akt, err
}

func (r *repository) Store(Aktifitas Aktifitas) (Aktifitas, error) {
	err := r.db.Create(&Aktifitas).Error
	return Aktifitas, err
}

func (r *repository) StoreImages(aktifitasId int, images []aktifitasImage.AktifitasImage) ([]aktifitasImage.AktifitasImage, error) {
	err := r.db.Where("aktifitas_id = ?", aktifitasId).Create(&images).Error
	return images, err
}

func (r *repository) DeleteImages(Id int) (aktifitasImage.AktifitasImage, error) {
	var image aktifitasImage.AktifitasImage
	err := r.db.Where("aktifitas_id = ?", Id).Delete(&image).Error
	return image, err
}

func (r *repository) UpdateAktifitas(Aktifitas Aktifitas) (Aktifitas, error) {
	err := r.db.Save(&Aktifitas).Error
	return Aktifitas, err
}

func (r *repository) DeleteAktifitas(Id int) (Aktifitas, error) {
	var d Aktifitas
	err := r.db.Where("id = ?", Id).First(&d).Error
	r.db.Delete(&d)
	return d, err
}

func (r *repository) CountCheckIN(tgl string, userId int) int64 {
	var count int64
	var p absensi.Absensi
	r.db.Model(&p).Where("tanggal_absen = ?", tgl).Where("check_absen = ?", "Check-IN").Where("user_id = ?", userId).Count(&count)
	return count
}
