package departement

import (
	"gorm.io/gorm"
)

type Departement struct {
	Id        int
	Name      string `gorm:"size:100"`
	CreatedBy int
	UpdatedBy int
	gorm.Model
}
