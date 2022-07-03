package seeds

import (
	"github.com/fendilaw41/absensifendiray/app/src/permission"
	"github.com/fendilaw41/absensifendiray/app/src/role"
	"github.com/fendilaw41/absensifendiray/app/src/user"

	"golang.org/x/crypto/bcrypt"
)

// TODO: Untuk Dummy User
func (s Seed) UserSeed() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)

	user := []user.User{
		{
			Name:          "Fendi",
			DepartementId: 1,
			DivisiId:      1,
			FirstName:     "Muhamad",
			LastName:      "Efendy",
			Fullname:      "Muhamad Efendy Ray",
			Email:         "fendi@gmail.com",
			Password:      string(bytes),
		},
		{
			DepartementId: 1,
			DivisiId:      1,
			Name:          "Admin",
			FirstName:     "Admin",
			LastName:      "istrator",
			Fullname:      "Administrator",
			Email:         "admin@gmail.com",
			Password:      string(bytes),
		},
		{
			DepartementId: 2,
			DivisiId:      2,
			Name:          "Master",
			FirstName:     "Muhamad",
			LastName:      "Keenan",
			Fullname:      "Muhamad Keenan Athariz",
			Email:         "master@gmail.com",
			Password:      string(bytes),
		},
	}

	s.db.Create(&user)
}

// TODO: Untuk Dummy Role
func (s Seed) RoleSeed() {

	role := []role.Role{
		{
			Name:        "superadmin",
			Description: "superadmin",
		},
		{
			Name:        "admin",
			Description: "administrator",
		},
		{
			Name:        "master",
			Description: "master",
		},
	}

	s.db.Create(&role)
}

// TODO: Untuk Dummy Pivot User & Role
func (s Seed) UserRoleSeed() { // untuk pivot

	usersRole := []user.UserRole{
		{
			UserId: 1,
			RoleId: 1,
		},
		{
			UserId: 2,
			RoleId: 2,
		},
		{
			UserId: 3,
			RoleId: 3,
		},
	}

	s.db.Create(&usersRole)
}

// TODO: Untuk Dummy Permission
func (s Seed) PermissionSeed() {

	permission := []permission.Permission{
		{
			Name:        "View Absen",
			Description: "Akses Melihat Data Absen",
		},
		{
			Name:        "Edit Absen",
			Description: "Akses Mengubah Data Absen",
		},
		{
			Name:        "Save Absen",
			Description: "Akses Menyimpan Data Absen",
		},
		{
			Name:        "Delete Absen",
			Description: "Akses Menghapus Data Absen",
		},
	}

	s.db.Create(&permission)
}

// TODO: Untuk Dummy Pivot Role & Permission
func (s Seed) RolePermissionSeed() { // untuk pivot

	rolePermission := []role.RolePermission{
		{
			RoleId:       1,
			PermissionId: 1,
		},
		{
			RoleId:       1,
			PermissionId: 2,
		},
		{
			RoleId:       1,
			PermissionId: 3,
		},
		{
			RoleId:       2,
			PermissionId: 1,
		},
	}

	s.db.Create(&rolePermission)
}
