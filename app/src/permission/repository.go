package permission

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllPaginate(page, page_size int, name string) ([]Permission, int64, error)
	GetDetail(Id int) (Permission, error)
	Store(p Permission) (Permission, error)
	UpdatePermission(p Permission) (Permission, error)
	DeletePermission(Id int) (Permission, error)
}

type repository struct {
	db *gorm.DB
}

func PermissionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllPaginate(page, page_size int, name string) ([]Permission, int64, error) {
	var p []Permission
	var count int64
	r.db.Model(&p).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&p)
	tx := r.db.Model(&p).Where("name LIKE ?", "%"+name+"%")
	err := tx.Model(&p).Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&p).Error

	return p, count, err
}

func (r *repository) GetDetail(Id int) (Permission, error) {
	var p Permission
	err := r.db.Find(&p, Id).Error
	return p, err
}

func (r *repository) Store(p Permission) (Permission, error) {
	err := r.db.Create(&p).Error
	return p, err
}

func (r *repository) UpdatePermission(p Permission) (Permission, error) {
	err := r.db.Save(&p).Error
	return p, err
}

func (r *repository) DeletePermission(Id int) (Permission, error) {
	var p Permission
	err := r.db.Where("id = ?", Id).First(&p).Error
	r.db.Delete(&p)
	return p, err
}
