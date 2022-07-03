package user

import (
	"github.com/fendilaw41/absensifendiray/app/src/departement"
	"github.com/fendilaw41/absensifendiray/app/src/divisi"
	"github.com/fendilaw41/absensifendiray/app/src/role"

	"gorm.io/datatypes"
)

type User struct {
	ID            int
	DepartementId int
	Departement   departement.Departement
	DivisiId      int
	Divisi        divisi.Divisi
	Name          string         `gorm:"size:50"`
	Fullname      string         `gorm:"size:50"`
	FirstName     string         `gorm:"size:50"`
	LastName      string         `gorm:"size:50"`
	Email         string         `gorm:"size:50;unique"`
	Password      string         `gorm:"size:100" json:"password"`
	Phone         string         `gorm:"size:20"`
	BirthDate     datatypes.Date `gorm:"default:NULL"`
	Gender        string         `gorm:"size:20"`
	Photo         string
	Status        string `gorm:"default: 1"`
	CreatedBy     int
	UpdatedBy     int
	Roles         []role.Role `gorm:"many2many:user_roles; constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	UserRoles     []UserRole  `gorm:"foreignkey:UserId"`
	Absensi       []Absensi   `gorm:"foreignkey:UserId"`
	Aktifitas     []Aktifitas `gorm:"foreignkey:UserId"`
}

type Aktifitas struct {
	UserId      int
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AssignedId  int
	Assigned    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string `gorm:"size:100"`
	Subject     string `gorm:"size:200"`
	Description string `gorm:"size:200"`
	CreatedAt   datatypes.Date
}

type Absensi struct {
	User         User
	UserId       int
	TanggalAbsen datatypes.Date
	JamAbsen     datatypes.Time
	Picture      string
	CheckAbsen   string // check in or checkout
	Status       string //
}

type UserRole struct { // TODO: Untuk Pivot Table
	User   User
	UserId int
	Role   User
	RoleId int
}

func (UserRole) TableName() string {
	return "user_roles"
}
