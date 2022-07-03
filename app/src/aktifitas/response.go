package aktifitas

import (
	"github.com/fendilaw41/absensifendiray/app/src/user"
	"github.com/fendilaw41/absensifendiray/config/action"

	"gorm.io/datatypes"
)

type AktifitasResponse struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Fullname    string `json:"fullname"`
	Name        string `json:"name"`
	Assigned    string `json:"di_assign_oleh"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	CreatedAt   string `json:"dibuat_pada"`
}

type AktifitasForUserResponse struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"nama_aktifitas"`
	Subject     string `json:"judul"`
	Description string `json:"keterangan"`
	CreatedAt   string `json:"dibuat_pada"`
}

type UserForAktifitasResponse struct {
	DepartementName string      `json:"departement_name" form:"departement_name"`
	DivisiName      string      `json:"divisi_name" form:"divisi_name"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	Aktifitas       interface{} `json:"aktifitas"`
}

func ResultAllAbsensiByUser(c user.Aktifitas) AktifitasForUserResponse {
	return AktifitasForUserResponse{
		UserId:      c.UserId,
		Name:        c.Name,
		Subject:     c.Subject,
		Description: c.Description,
		CreatedAt:   action.FormatDateToString(datatypes.Date(c.CreatedAt)),
	}
}

func ResultUserForAktifitas(c user.User, aktifitas interface{}) UserForAktifitasResponse {
	return UserForAktifitasResponse{
		DepartementName: c.Departement.Name,
		DivisiName:      c.Divisi.Name,
		FirstName:       c.FirstName,
		LastName:        c.LastName,
		Aktifitas:       aktifitas,
	}
}
