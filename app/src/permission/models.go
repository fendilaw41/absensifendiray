package permission

import "gorm.io/gorm"

type Permission struct {
	Id          int
	Name        string `gorm:"size:50"`
	Description string `gorm:"size:200"`
	gorm.Model
}
