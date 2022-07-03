package absensi

import (
	"time"

	"github.com/fendilaw41/absensifendiray/app/src/user"
	"github.com/fendilaw41/absensifendiray/config/action"
)

type AbsensiResponse struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id"`
	FirstName    string `json:"firstname"  form:"firstname"`
	LastName     string `json:"lastname"  form:"lastname"`
	FullName     string `json:"fullname"  form:"fullname"`
	TanggalAbsen string `json:"tanggal_absen" form:"tanggal_absen" binding:"required"`
	JamAbsen     string `json:"jam_absen"  form:"jam_absen"`
	Picture      string `json:"picture"  form:"picture"`
	CheckAbsen   string `json:"cek_absen"  form:"cek_absen"`
	Status       string `json:"status"  form:"status"`
	CreatedBy    int    `json:"created_by"`
}

type AbsensiForUserResponse struct {
	TanggalAbsen string `json:"tanggal_absen" form:"tanggal_absen" binding:"required"`
	JamAbsen     string `json:"jam_absen"  form:"jam_absen"`
	Picture      string `json:"picture"  form:"picture"`
	CheckAbsen   string `json:"cek_absen"  form:"cek_absen"`
	Status       string `json:"status"  form:"status"`
}

type UserForAbsenResponse struct {
	ID              int    `json:"id"`
	DepartementName string `json:"departement_name" form:"departement_name"`
	DivisiName      string `json:"divisi_name" form:"divisi_name"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`

	Roles   interface{} `json:"roles"`
	Absensi interface{} `json:"absensi"`
}

func ResultAbsensiWithJoin(c Absensi) AbsensiResponse {
	return AbsensiResponse{
		ID:           c.ID,
		UserId:       c.UserId,
		FirstName:    c.User.FirstName,
		LastName:     c.User.LastName,
		FullName:     c.User.Fullname,
		TanggalAbsen: action.FormatDateToString(c.TanggalAbsen),
		JamAbsen:     action.FormatTime(time.Duration(c.JamAbsen)),
		CheckAbsen:   c.CheckAbsen,
		Status:       c.Status,
		Picture:      c.Picture,
		CreatedBy:    c.CreatedBy,
	}
}

func ResultAllAbsensi(c Absensi) AbsensiResponse {
	return AbsensiResponse{
		ID:           c.ID,
		UserId:       c.UserId,
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		FullName:     c.FullName,
		TanggalAbsen: action.FormatDateToString(c.TanggalAbsen),
		JamAbsen:     action.FormatTime(time.Duration(c.JamAbsen)),
		Picture:      c.Picture,
		CheckAbsen:   c.CheckAbsen,
		CreatedBy:    c.CreatedBy,
	}
}

func ResultUserForAbsen(u user.User, roles interface{}, absensi interface{}) UserForAbsenResponse {
	return UserForAbsenResponse{
		ID:              u.ID,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		DepartementName: u.Departement.Name,
		DivisiName:      u.Divisi.Name,
		Roles:           roles,
		Absensi:         absensi,
	}
}

func ResultAllAbsensiByUser(c user.Absensi) AbsensiForUserResponse {
	return AbsensiForUserResponse{
		TanggalAbsen: action.FormatDateToString(c.TanggalAbsen),
		JamAbsen:     action.FormatTime(time.Duration(c.JamAbsen)),
		Picture:      c.Picture,
		CheckAbsen:   c.CheckAbsen,
		Status:       c.Status,
	}
}
