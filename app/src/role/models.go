package role

import (
	"github.com/fendilaw41/absensifendiray/app/src/permission"

	"gorm.io/gorm"
)

type Role struct {
	Id              uint                    `gorm:"size:50"`
	Name            string                  `gorm:"size:50"`
	Description     string                  `gorm:"size:200"`
	Permission      []permission.Permission `gorm:"many2many:role_permissions;"` // TODO: Untuk mengambil data permission
	RolePermissions []RolePermission        `gorm:"foreignkey:RoleId"`           // TODO: Ambil struct nya lalu kirimkan ke table pivot Role permission
	gorm.Model
}

type RolePermission struct { // TODO: Untuk Pivot Table
	Role         Role `gorm:"association_foreignkey:RoleId"` // TODO: Ambil struct Role nya
	RoleId       uint
	Permission   Role `gorm:"association_foreignkey:PermissionId"` // TODO: Type data role untuk menghubungkan dengan permission melalui role
	PermissionId uint
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

func Any(roles []Role, f func(Role) bool) bool {
	for _, role := range roles {
		if f(role) {
			return true
		}
	}
	return false
}
