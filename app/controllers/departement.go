package controllers

import (
	"fmt"
	"strconv"

	departement "github.com/fendilaw41/absensifendiray/app/src/departement"
	"github.com/fendilaw41/absensifendiray/config/action"
	"github.com/fendilaw41/absensifendiray/config/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DepartementController struct {
	DepartementService departement.Service
}

func constDepartement() *DepartementController {
	db, _ := database.DbSetup()
	DepartementRepository := departement.DepartementRepository(db)
	DepartementService := departement.DepartementService(DepartementRepository)
	return &DepartementController{DepartementService}
}

func GetAllDepartement(res *gin.Context) {
	Departement, err := constDepartement().DepartementService.FindAll()
	if err != nil {
		action.BadRequest(err, res)
		return
	}
	var cats []departement.DepartementResponse
	for _, u := range Departement {
		cat := departement.DepartementResponse{
			Id:   u.Id,
			Name: u.Name,
		}
		cats = append(cats, cat)
	}
	if cats == nil {
		action.NotFound(res)
		return
	}
	action.AccepCount(action.All("Departement"), len(cats), cats, res)
}

func GetAllDepartementPaginate(res *gin.Context) {
	departements, page, pageSize, count, err := constDepartement().DepartementService.FindAllPaginate(res)
	if err != nil {
		action.BadRequest(err, res)
		return
	}
	var cats []departement.DepartementResponse
	for _, u := range departements {

		cat := departement.DepartementResponse{
			Id:   u.Id,
			Name: u.Name,
		}
		cats = append(cats, cat)
	}
	if cats == nil {
		action.NotFound(res)
		return
	}
	action.AcceptPaginate(res.Request, res, action.All("Paginate Departement"), cats, len(cats), page, pageSize, count)
}

func GetDetailDepartement(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)

	u, err := constDepartement().DepartementService.FindId(int(id))

	if err != nil {
		action.NotFound(res)
		return
	}

	DepartementRes := resultDepartement(u)
	action.AccepData(action.Detail("Departement"), DepartementRes, res)
}

func StoreDepartement(res *gin.Context) {
	var DepartementReq departement.DepartementRequest
	err := res.ShouldBindJSON(&DepartementReq)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error: %s, %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		action.BadRequest(errorMessages, res)
		return
	}

	Departement, err := constDepartement().DepartementService.SaveDepartementSrv(DepartementReq)
	if err != nil {
		action.BadRequest(err, res)
		return
	}

	action.AccepData(action.Save("Departement"), resultDepartement(Departement), res)
}

func UpdateDepartement(res *gin.Context) {
	var DepartementReq departement.DepartementRequest
	err := res.ShouldBindJSON(&DepartementReq)
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
	Departement, err := constDepartement().DepartementService.UpdateDepartementSrv(id, DepartementReq)
	if err != nil {
		action.BadRequest(err, res)
		return
	}
	action.AccepData(action.Update("Departement"), resultDepartement(Departement), res)
}

func DeleteDepartement(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, err := constDepartement().DepartementService.SrcDeleteDepartement(int(id))

	if err != nil {
		action.BadRequest(err, res)
	}

	DepartementResponse := resultDepartement(u)
	action.AccepData(action.Delete("Departement"), DepartementResponse, res)
}

func resultDepartement(c departement.Departement) departement.DepartementResponse {
	return departement.DepartementResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}
