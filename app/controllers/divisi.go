package controllers

import (
	"fmt"
	"strconv"

	"github.com/fendilaw41/absensifendiray/app/src/divisi"
	"github.com/fendilaw41/absensifendiray/config/action"
	"github.com/fendilaw41/absensifendiray/config/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type divisiController struct {
	divisiService divisi.Service
}

func constdivisi() *divisiController {
	db, _ := database.DbSetup()
	divisiRepository := divisi.DivisiRepository(db)
	divisiService := divisi.DivisiService(divisiRepository)
	return &divisiController{divisiService}
}

func GetAllDivisi(res *gin.Context) {
	divisis, err := constdivisi().divisiService.FindAll()
	if err != nil {
		action.BadRequest(err, res)
	}

	var divisisRes []divisi.DivisiResponse

	for _, u := range divisis {
		divisiRes := resultdivisi(u)
		divisisRes = append(divisisRes, divisiRes)
	}
	action.AccepCount(action.All("divisi"), len(divisisRes), divisisRes, res)
}

func GetAllDivisiPaginate(res *gin.Context) {
	divisis, page, pageSize, count, err := constdivisi().divisiService.FindAllPaginate(res)
	if err != nil {
		action.BadRequest(err, res)
		return
	}
	var newdivisis []divisi.DivisiResponse
	for _, u := range divisis {
		cat := divisi.DivisiResponse{
			ID:   u.ID,
			Name: u.Name,
		}
		newdivisis = append(newdivisis, cat)
	}
	if newdivisis == nil {
		action.NotFound(res)
		return
	}
	action.AcceptPaginate(res.Request, res, action.All("Paginate divisi"), newdivisis, len(newdivisis), page, pageSize, count)
}

func GetDetailDivisi(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, err := constdivisi().divisiService.FindId(int(id))

	if err != nil {
		action.BadRequest(err, res)
	}

	divisiResponse := resultdivisi(u)
	action.AccepData(action.Detail("divisi"), divisiResponse, res)
}

func StoreDivisi(res *gin.Context) {
	var divisiReq divisi.DivisiRequest
	err := res.ShouldBindJSON(&divisiReq)

	if err != nil {
		fmt.Println(err)
		action.BadRequest(err, res)
		return
	}

	divisi, err := constdivisi().divisiService.Save(divisiReq)
	if err != nil {
		action.BadRequest(err, res)
		return
	}

	action.AccepData(action.Save("divisi"), resultdivisi(divisi), res)
}

func UpdateDivisi(res *gin.Context) {
	var divisiReq divisi.DivisiRequest
	err := res.ShouldBindJSON(&divisiReq)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error: %s, %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		action.BadRequest(errorMessages, res)
		return
	}
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	divisi, err := constdivisi().divisiService.UpdateDivisiService(id, divisiReq)
	if err != nil {
		action.BadRequest(err, res)
	}
	action.AccepData(action.Update("divisi"), resultdivisi(divisi), res)
}

func DeleteDivisi(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, err := constdivisi().divisiService.SrcDeleteDivisi(int(id))

	if err != nil {
		action.BadRequest(err, res)
	}

	divisiResponse := resultdivisi(u)
	action.AccepData(action.Delete("divisi"), divisiResponse, res)
}

func resultdivisi(u divisi.Divisi) divisi.DivisiResponse {
	return divisi.DivisiResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
