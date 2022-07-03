package user

import (
	"github.com/fendilaw41/absensifendiray/app/src/role"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllPaginate(page, page_size int, name string) ([]User, int64, error)
	GetAll() ([]User, error)
	GetDetail(Id int) (User, error)
	Store(user User) (User, error)
	StoreUserRole(role UserRole) error
	FirstRoleName(name string) (role.Role, error)
	UpdateUser(user User) (User, error)
	DeleteUser(Id int) (User, error)
	DeleteUserRole(Id int) (UserRole, error)
}

type repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetAllPaginate(page, page_size int, name string) ([]User, int64, error) {
	var users []User
	var count int64
	r.db.Model(&users).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&users)
	// filter data string
	tx := r.db.Model(&users).Where("name LIKE ?", "%"+name+"%")

	err := tx.Model(&users).Preload("Aktifitas").Preload("Absensi").Preload("Divisi").Preload("Departement").Preload("Roles.Permission").Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&users).Error

	return users, count, err
}

func (r *repository) GetAllCustomerPaginate(page, page_size int, name string) ([]User, int64, error) {
	var users []User
	var count int64
	r.db.Model(&users).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&users)
	// filter data string
	tx := r.db.Model(&users).Where("name LIKE ?", "%"+name+"%")

	err := tx.Model(&users).Preload("Divisi").Preload("Departement").Preload("Roles.Permission").Preload("Roles", "name = (?)", "customer").Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&users).Error

	return users, count, err
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	tx := r.db.Begin()
	tx.Model(&users).Preload("Divisi").Preload("Departement").Preload("Roles.Permission").Order("id asc").Find(&users)
	err := tx.Commit().Error
	return users, err
}

func (r *repository) GetDetail(Id int) (User, error) {
	var user User
	err := r.db.Preload("Divisi").Preload("Departement").Preload("Roles.Permission").Find(&user, Id).Error
	return user, err
}

func (r *repository) Store(user User) (User, error) {
	err := r.db.Preload("Roles.Permission").Create(&user).Error
	return user, err
}

func (r *repository) StoreUserRole(role UserRole) error {
	err := r.db.Create(&role).Error
	return err
}

func (r *repository) FirstRoleName(name string) (role.Role, error) {
	var roleModel role.Role
	err := r.db.Where("name = ?", name).First(&roleModel).Error
	return roleModel, err
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Preload("Roles.Permission").Save(&user).Error
	return user, err
}

func (r *repository) DeleteUser(Id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", Id).First(&user).Error
	r.db.Delete(&user)
	return user, err
}

func (r *repository) DeleteUserRole(Id int) (UserRole, error) {
	var userRole UserRole
	err := r.db.Where("user_id = ?", Id).Delete(&userRole).Error
	return userRole, err
}
