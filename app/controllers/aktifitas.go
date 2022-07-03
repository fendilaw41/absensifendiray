package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fendilaw41/absensifendiray/app/src/aktifitas"
	"github.com/fendilaw41/absensifendiray/config/action"
	"github.com/fendilaw41/absensifendiray/config/database"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type aktifitasController struct {
	aktifitasService aktifitas.Service
}

func constaktifitas() *aktifitasController {
	db, _ := database.DbSetup()
	aktifitasRepository := aktifitas.AktifitasRepository(db)
	aktifitasService := aktifitas.AktifitasService(aktifitasRepository)
	return &aktifitasController{aktifitasService}
}

func GetAllAktifitas(res *gin.Context) {
	aktifitass, err := constaktifitas().aktifitasService.FindAll()
	if err != nil {
		action.BadRequest(err, res)
	}

	var aktifitassRes []aktifitas.AktifitasResponse

	for _, u := range aktifitass {
		aktifitasRes := resultaktifitas(u)
		aktifitassRes = append(aktifitassRes, aktifitasRes)
	}
	action.AccepCount(action.All("aktifitas"), len(aktifitassRes), aktifitassRes, res)
}

func RiwayatAktifitas(res *gin.Context) {
	users, page, pageSize, count, err := constaktifitas().aktifitasService.RiwayatAktifitas(res)
	if err != nil {
		action.BadRequest(err, res)
		return
	}

	var userModel []aktifitas.UserForAktifitasResponse
	for _, u := range users {
		var aktifitasRes []aktifitas.AktifitasForUserResponse
		for _, v := range u.Aktifitas {
			newAktifitas := aktifitas.ResultAllAbsensiByUser(v)
			aktifitasRes = append(aktifitasRes, newAktifitas)
		}
		newUser := aktifitas.ResultUserForAktifitas(u, aktifitasRes)
		userModel = append(userModel, newUser)
	}
	if userModel == nil {
		action.NotFound(res)
		return
	}
	action.AcceptPaginate(res.Request, res, action.All("Riwayat Aktifitas"), userModel, len(userModel), page, pageSize, count)
}

func GetDetailAktifitas(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, errAktifitasId := constaktifitas().aktifitasService.FindId(int(id))

	if errAktifitasId != nil {
		action.CustomError("ID tidak ditemukan", res)
		return
	}

	aktifitasResponse := resultaktifitas(u)
	action.AccepData(action.Detail("aktifitas"), aktifitasResponse, res)
}

func StoreAktifitas(res *gin.Context) {
	var aktifitasReq aktifitas.AktifitasRequest
	err := res.ShouldBind(&aktifitasReq)

	if err != nil {
		fmt.Println(err)
		action.BadRequest(err, res)
		return
	}

	authId := action.AuthId(res)
	today := action.FormatDateToString(datatypes.Date(time.Now()))
	countDate := constaktifitas().aktifitasService.CountCheckIN(today, authId)
	if countDate == 1 {
		_, err := constaktifitas().aktifitasService.Save(res, authId, aktifitasReq)
		if err != nil {
			action.BadRequest(err, res)
			return
		}
	} else {
		action.CustomError("Anda Belum Check-IN Hari Ini! Silahkan Check-IN terlebih dahulu", res)
		return
	}

	action.AccepMsg(action.Save("Aktifitas"), res)
}

func UpdateAktifitas(res *gin.Context) {
	var aktifitasReq aktifitas.AktifitasRequestPUT
	err := res.ShouldBind(&aktifitasReq)
	if err != nil {
		fmt.Println(err)
		action.BadRequest(err, res)
		return
	}

	authId := action.AuthId(res)
	today := action.FormatDateToString(datatypes.Date(time.Now()))
	countDate := constaktifitas().aktifitasService.CountCheckIN(today, authId)
	if countDate == 1 {
		idString := res.Param("id")
		id, _ := strconv.Atoi(idString)
		_, err := constaktifitas().aktifitasService.UpdateAktifitasService(res, id, aktifitasReq, authId)
		if err != nil {
			action.BadRequest(err, res)
		}
	} else {
		action.CustomError("Anda Belum Check-IN Hari Ini! Silahkan Check-IN terlebih dahulu", res)
		return
	}

	action.AccepMsg(action.Update("Aktifitas"), res)
}

func DeleteAktifitas(res *gin.Context) {
	authId := action.AuthId(res)
	today := action.FormatDateToString(datatypes.Date(time.Now()))
	countDate := constaktifitas().aktifitasService.CountCheckIN(today, authId)
	if countDate == 1 {
		idString := res.Param("id")
		id, _ := strconv.Atoi(idString)
		aktifitasId, errAktifitasId := constaktifitas().aktifitasService.FindId(int(id))
		if errAktifitasId != nil {
			action.CustomError("ID tidak ditemukan", res)
			return
		}
		_, err := constaktifitas().aktifitasService.SrcDeleteAktifitas(aktifitasId.ID)
		if err != nil {
			action.BadRequest(err, res)
		}
	} else {
		action.CustomError("Anda Belum Check-IN Hari Ini! Silahkan Check-IN terlebih dahulu", res)
		return
	}

	action.AccepMsg(action.Delete("Aktifitas"), res)
}

func resultaktifitas(u aktifitas.Aktifitas) aktifitas.AktifitasResponse {
	return aktifitas.AktifitasResponse{
		ID:          u.ID,
		UserId:      u.User.ID,
		Fullname:    u.User.Fullname,
		Name:        u.Name,
		Subject:     u.Subject,
		Description: u.Description,
		Assigned:    u.User.Fullname,
		CreatedAt:   action.FormatDateToString(datatypes.Date(u.CreatedAt)),
	}
}
