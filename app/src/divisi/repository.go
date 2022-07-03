package divisi

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Divisi, error)
	GetAllPaginate(page, page_size int, name string) ([]Divisi, int64, error)
	GetDetail(Id int) (Divisi, error)
	Store(Divisi Divisi) (Divisi, error)
	UpdateDivisi(Divisi Divisi) (Divisi, error)
	DeleteDivisi(Id int) (Divisi, error)
}

type repository struct {
	db *gorm.DB
}

func DivisiRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Divisi, error) {
	var Divisi []Divisi
	err := r.db.Find(&Divisi).Error
	return Divisi, err
}

func (r *repository) GetAllPaginate(page, page_size int, name string) ([]Divisi, int64, error) {
	var Divisi []Divisi
	var count int64
	r.db.Model(&Divisi).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&Divisi)
	// filter data string
	tx := r.db.Model(&Divisi).Where("name LIKE ?", "%"+name+"%")

	err := tx.Model(&Divisi).Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&Divisi).Error

	return Divisi, count, err
}

func (r *repository) GetDetail(Id int) (Divisi, error) {
	var Divisi Divisi
	err := r.db.Find(&Divisi, Id).Error
	return Divisi, err
}

func (r *repository) Store(Divisi Divisi) (Divisi, error) {
	err := r.db.Create(&Divisi).Error
	return Divisi, err
}

func (r *repository) UpdateDivisi(Divisi Divisi) (Divisi, error) {
	err := r.db.Save(&Divisi).Error
	return Divisi, err
}

func (r *repository) DeleteDivisi(Id int) (Divisi, error) {
	var d Divisi
	err := r.db.Where("id = ?", Id).First(&d).Error
	r.db.Delete(&d)
	return d, err
}
