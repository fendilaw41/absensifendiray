package divisi

import (
	"gorm.io/gorm"
)

type Divisi struct {
	ID        int
	Name      string `gorm:"size:100"`
	CreatedBy int
	UpdatedBy int
	gorm.Model
}
