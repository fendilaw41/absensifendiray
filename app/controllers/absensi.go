package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fendilaw41/absensifendiray/app/src/absensi"
	"github.com/fendilaw41/absensifendiray/app/src/user"
	"github.com/fendilaw41/absensifendiray/config/action"
	"github.com/fendilaw41/absensifendiray/config/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/datatypes"
)

type absensiController struct {
	absenservice absensi.Service
}

func constAbsensi() *absensiController {
	db, _ := database.DbSetup()
	absensiRepository := absensi.AbsensiRepository(db)
	absenservice := absensi.AbsensiService(absensiRepository)
	return &absensiController{absenservice}
}

func RiwayatAbsensi(res *gin.Context) {
	users, page, pageSize, count, err := constUser().userService.FindAllPaginate(res)
	if err != nil {
		action.BadRequest(err, res)
		return
	}

	var userModel []absensi.UserForAbsenResponse
	for _, u := range users {
		var roles []user.RolesResponse
		var absensiRes []absensi.AbsensiForUserResponse
		for _, v := range u.Roles {
			var permission []user.PermissionResponse
			for _, x := range v.Permission {
				newPermission := user.ResultPermission(x)
				permission = append(permission, newPermission)
			}
			newRoles := user.ResultRole(v, permission)
			roles = append(roles, newRoles)
		}
		for _, v := range u.Absensi {
			newAbsensi := absensi.ResultAllAbsensiByUser(v)
			absensiRes = append(absensiRes, newAbsensi)
		}
		newUser := absensi.ResultUserForAbsen(u, roles, absensiRes)
		userModel = append(userModel, newUser)
	}
	if userModel == nil {
		action.NotFound(res)
		return
	}
	action.AcceptPaginate(res.Request, res, action.All("Riwayat Absensi"), userModel, len(userModel), page, pageSize, count)
}

func GetAllAbsensi(res *gin.Context) {
	absensis, err := constAbsensi().absenservice.FindAll()
	if err != nil {
		action.BadRequest(err, res)
		return
	}
	var absenRes []absensi.AbsensiResponse

	for _, u := range absensis {
		absensiRes := absensi.ResultAbsensiWithJoin(u)
		absenRes = append(absenRes, absensiRes)
	}

	action.AccepCount(action.All("absensi"), len(absenRes), absenRes, res)
}

func GetByIdAbsensi(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, err := constAbsensi().absenservice.FindId(int(id))

	if err != nil {
		action.NotFound(res)
		return
	}
	absensiRes := absensi.ResultAbsensiWithJoin(u)
	action.AccepData(action.Detail("absensi"), absensiRes, res)
}

func CheckIn(res *gin.Context) {
	var req absensi.AbsensiRequest
	errabsensiReq := res.ShouldBind(&req)
	if errabsensiReq != nil {
		fmt.Println(errabsensiReq)
		action.BadRequest(errabsensiReq, res)
		return
	}

	if file, _ := res.FormFile("picture"); file != nil { // jika file di isi
		filename := file.Filename
		path := "storage/absensi/" + filename
		saveUpload := res.SaveUploadedFile(file, path)
		if saveUpload != nil {
			action.CustomError("Save File absensi Salah", res)
			return
		}
	}
	authId := action.AuthId(res)
	today := action.FormatDateToString(datatypes.Date(time.Now()))
	countDate := constAbsensi().absenservice.CountCheckIN(today, authId)
	if countDate > 0 {
		action.CustomError("Anda Sudah Check-IN Hari Ini! Silahkan Checkout terlebih dahulu", res)
		return
	} else {
		_, err := constAbsensi().absenservice.CheckIn(req, res)
		if err != nil {
			action.BadRequest(err, res)
			return
		}
	}

	action.AccepMsg("Success CheckIn Hari ini!", res)
}

func CheckOut(res *gin.Context) {
	var req absensi.AbsensiRequest
	errabsensiReq := res.ShouldBind(&req)
	if errabsensiReq != nil {
		fmt.Println(errabsensiReq)
		action.BadRequest(errabsensiReq, res)
		return
	}

	if file, _ := res.FormFile("picture"); file != nil { // jika file di isi
		filename := file.Filename
		path := "storage/absensi/" + filename
		saveUpload := res.SaveUploadedFile(file, path)
		if saveUpload != nil {
			action.CustomError("Save File absensi Salah", res)
			return
		}
	}
	authId := action.AuthId(res)
	today := action.FormatDateToString(datatypes.Date(time.Now()))
	countDate := constAbsensi().absenservice.CountCheckOut(today, authId)
	if countDate > 0 {
		action.CustomError("Anda Sudah Check-OUT Hari Ini! Silahkan Kembali Lagi Esok Hari...", res)
		return
	} else {
		_, err := constAbsensi().absenservice.CheckOut(req, res)
		if err != nil {
			action.BadRequest(err, res)
			return
		}
	}

	action.AccepMsg("Success Check-OUT Hari ini!", res)
}

func UpdateAbsensi(res *gin.Context) {
	var req absensi.AbsensiRequest
	if errReq := res.ShouldBind(&req); errReq != nil {
		errorMessages := []string{}
		for _, e := range errReq.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Gagal Validasi Form Request: %s, %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		action.BadRequest(errorMessages, res)
		return
	}
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	if file, _ := res.FormFile("picture"); file != nil { // jika file di isi
		filename := file.Filename
		path := "storage/absensi/" + filename
		saveUpload := res.SaveUploadedFile(file, path)
		if saveUpload != nil {
			action.CustomError("Update File absensi Salah", res)
			return
		}
	}
	result, err := constAbsensi().absenservice.UpdateAbsensiService(id, req, res)
	if err != nil {
		action.BadRequest(err, res)
	}
	action.AccepData(action.Update("absensi"), absensi.ResultAllAbsensi(result), res)
}

func DeleteAbsensi(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, err := constAbsensi().absenservice.SrcDeleteAbsensi(int(id))

	if err != nil {
		action.BadRequest(err, res)
		return
	}

	AbsensiResponse := absensi.ResultAbsensiWithJoin(u)
	action.AccepData(action.Delete("absensi"), AbsensiResponse, res)
}
