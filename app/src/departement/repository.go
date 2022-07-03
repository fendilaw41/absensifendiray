package departement

import "gorm.io/gorm"

type RepositoryDepartement interface {
	GetAll() ([]Departement, error)
	GetAllPaginate(page, page_size int, name string) ([]Departement, int64, error)
	GetDetail(Id int) (Departement, error)
	StoreDepartementRepo(Departement Departement) (Departement, error)
	UpdateDepartementRepo(Departement Departement) (Departement, error)
	DeleteDepartementRepo(Id int) (Departement, error)
}

type repositoryDepartement struct {
	db *gorm.DB
}

func DepartementRepository(db *gorm.DB) *repositoryDepartement {
	return &repositoryDepartement{db}
}

func (r *repositoryDepartement) GetAll() ([]Departement, error) {
	var categories []Departement
	err := r.db.Order("id asc").Find(&categories).Error
	return categories, err
}

func (r *repositoryDepartement) GetAllPaginate(page, page_size int, name string) ([]Departement, int64, error) {
	var categories []Departement
	var count int64
	r.db.Model(&categories).Count(&count)
	r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&categories)
	// filter data string
	tx := r.db.Model(&categories).Where("name LIKE ?", "%"+name+"%")

	err := tx.Model(&categories).Order("id asc").Offset((page - 1) * page_size).Limit(page_size).Find(&categories).Error

	return categories, count, err
}

func (r *repositoryDepartement) GetDetail(Id int) (Departement, error) {
	var Departement Departement
	err := r.db.Where("id = ?", Id).First(&Departement).Error
	return Departement, err
}

func (r *repositoryDepartement) StoreDepartementRepo(Departement Departement) (Departement, error) {
	err := r.db.Create(&Departement).Error
	return Departement, err
}

func (r *repositoryDepartement) UpdateDepartementRepo(Departement Departement) (Departement, error) {
	err := r.db.Save(&Departement).Error
	return Departement, err
}

func (r *repositoryDepartement) DeleteDepartementRepo(Id int) (Departement, error) {
	var Departement Departement
	err := r.db.Where("id = ?", Id).First(&Departement).Error
	r.db.Delete(&Departement)
	return Departement, err
}
