package aktifitas

import (
	"github.com/fendilaw41/absensifendiray/app/src/user"

	"gorm.io/gorm"
)

type Aktifitas struct {
	ID          int
	UserId      int
	User        user.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AssignedId  int
	Assigned    user.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string    `gorm:"size:100"`
	Subject     string    `gorm:"size:200"`
	Description string    `gorm:"size:200"`
	CreatedBy   int
	UpdatedBy   int
	// CreatedAt   datatypes.Date
	// UpdatedAt   datatypes.Date
	gorm.Model
}
