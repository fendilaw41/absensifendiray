package absensi

import (
	"github.com/fendilaw41/absensifendiray/app/src/user"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Absensi struct {
	ID           int
	UserId       int
	User         user.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FirstName    string    `gorm:"size: 50"`
	LastName     string
	FullName     string `gorm:"size: 150"`
	TanggalAbsen datatypes.Date
	JamAbsen     datatypes.Time
	Picture      string
	CheckAbsen   string // check in or checkout
	Status       string // Terlambat or belum terlambat
	CreatedBy    int
	UpdatedBy    int
	gorm.Model
}
