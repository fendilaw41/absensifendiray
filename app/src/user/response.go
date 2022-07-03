package user

import (
	"github.com/fendilaw41/absensifendiray/app/src/permission"
	"github.com/fendilaw41/absensifendiray/app/src/role"
	"github.com/fendilaw41/absensifendiray/config/action"
)

type UserResponse struct {
	ID              int    `json:"id"`
	DepartementId   int    `json:"departement_id" form:"departement_id"`
	DepartementName string `json:"departement_name" form:"departement_name"`
	DivisiId        int    `json:"divisi_id" form:"divisi_id"`
	DivisiName      string `json:"divisi_name" form:"divisi_name"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Photo           string `json:"photo"`
	FullName        string `json:"full_name"`
	Phone           string `json:"phone"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	BirthDate       string `json:"birth_date"`
	Gender          string `json:"gender"`

	Roles interface{} `json:"roles"`
}
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token    string      `json:"token"`
	Name     string      `json:"name"`
	Fullname string      `json:"fullname"`
	Email    string      `json:"email"`
	Roles    interface{} `json:"roles"`
}

type RolesResponse struct {
	Name       string      `json:"role_name"`
	Permission interface{} `json:"permission"`
}

type PermissionResponse struct {
	Name string `json:"name"`
}

func ResultUser(u User, roles interface{}) UserResponse {
	return UserResponse{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		FullName:        u.Fullname,
		BirthDate:       action.FormatDateToString(u.BirthDate),
		Gender:          u.Gender,
		DepartementId:   u.DepartementId,
		DepartementName: u.Departement.Name,
		DivisiId:        u.DivisiId,
		DivisiName:      u.Divisi.Name,
		Photo:           u.Photo,
		Roles:           roles,
	}
}

func ResultRole(u role.Role, permission interface{}) RolesResponse {
	return RolesResponse{
		Name:       u.Name,
		Permission: permission,
	}
}

func ResultPermission(u permission.Permission) PermissionResponse {
	return PermissionResponse{
		Name: u.Name,
	}
}
