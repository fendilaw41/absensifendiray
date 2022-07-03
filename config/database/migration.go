package database

import (
	"github.com/fendilaw41/absensifendiray/app/src/absensi"
	"github.com/fendilaw41/absensifendiray/app/src/aktifitas"
	"github.com/fendilaw41/absensifendiray/app/src/aktifitasImage"
	"github.com/fendilaw41/absensifendiray/app/src/departement"
	"github.com/fendilaw41/absensifendiray/app/src/divisi"
	"github.com/fendilaw41/absensifendiray/app/src/permission"
	"github.com/fendilaw41/absensifendiray/app/src/role"
	"github.com/fendilaw41/absensifendiray/app/src/user"
)

func DbMigration() {
	db.AutoMigrate(&divisi.Divisi{})
	db.AutoMigrate(&departement.Departement{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&absensi.Absensi{})
	db.AutoMigrate(&role.Role{})
	db.AutoMigrate(&user.UserRole{}) // TODO: Untuk Pivot Table
	db.AutoMigrate(&permission.Permission{})
	db.AutoMigrate(&role.RolePermission{})
	db.AutoMigrate(&aktifitas.Aktifitas{})
	db.AutoMigrate(&aktifitasImage.AktifitasImage{})
}

func Drop() {
	db.Migrator().DropTable(&departement.Departement{})
	db.Migrator().DropTable(&divisi.Divisi{})
	db.Migrator().DropTable(&user.User{})
	db.Migrator().DropTable(&absensi.Absensi{})
	db.Migrator().DropTable(&role.Role{})
	db.Migrator().DropTable(&user.UserRole{}) // TODO: Untuk Pivot Table
	db.Migrator().DropTable(&permission.Permission{})
	db.Migrator().DropTable(&role.RolePermission{})
	db.Migrator().DropTable(&aktifitas.Aktifitas{})
	db.Migrator().DropTable(&aktifitasImage.AktifitasImage{})
}
