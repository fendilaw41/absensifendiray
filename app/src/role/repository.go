package role

import (
	"github.com/fendilaw41/absensifendiray/app/src/permission"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllPaginate(page, page_size int, name string) ([]Role, int64, error)
	GetAll() ([]Role, error)
	GetDetail(Id int) (Role, error)
	Store(role Role) (Role, error)
	StoreRolePermission(rp RolePermission) (RolePermission, error)
	FirstPermissionName(name string) (permission.Permission, error)
	UpdateRole(role Role) (Role, error)
	DeleteRole(Id int) (Role, error)
	DeleteRolePermission(Id int) (RolePermission, error)
}

type repository struct {
	db *gorm.DB
}

func RoleRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllPaginate(page, page_size int, name string) ([]Role, int64, error) {
	var roles []Role
	var count int64
	r.db.Model(&roles).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&roles)
	tx := r.db.Model(&roles).Where("name LIKE ?", "%"+name+"%")
	err := tx.Model(&roles).Preload("Permission").Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&roles).Error

	return roles, count, err
}

func (r *repository) GetAll() ([]Role, error) {
	var roles []Role
	tx := r.db.Begin()
	tx.Model(&roles).Preload("Permission").Order("id asc").Find(&roles)
	err := tx.Commit().Error
	return roles, err
}

func (r *repository) GetDetail(Id int) (Role, error) {
	var role Role
	err := r.db.Preload("Permission").Find(&role, Id).Error
	return role, err
}

func (r *repository) Store(role Role) (Role, error) {
	err := r.db.Create(&role).Error
	return role, err
}

func (r *repository) UpdateRole(role Role) (Role, error) {
	err := r.db.Save(&role).Error
	return role, err
}

func (r *repository) DeleteRole(Id int) (Role, error) {
	var role Role
	err := r.db.Where("id = ?", Id).First(&role).Error
	r.db.Delete(&role)
	return role, err
}

func (r *repository) StoreRolePermission(rp RolePermission) (RolePermission, error) {
	err := r.db.Create(&rp).Error
	return rp, err
}

func (r *repository) FirstPermissionName(name string) (permission.Permission, error) {
	var permissionModel permission.Permission
	err := r.db.First(&permissionModel, "name = ?", name).Error
	return permissionModel, err
}

func (r *repository) DeleteRolePermission(Id int) (RolePermission, error) {
	var rp RolePermission
	err := r.db.Where("role_id = ?", Id).Delete(&rp).Error
	return rp, err
}
