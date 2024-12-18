package controller

import (
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type PengeluaranController struct {
	P service.PengeluaranServiceInj
}

type PengeluaranControllerInj interface {
	CreatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetPengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetPengeluaranById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	UpdatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	DeletePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewPengeluaranController(p service.PengeluaranServiceInj) PengeluaranControllerInj {
	return PengeluaranController{
		P: p,
	}
}

func (pc PengeluaranController) CreatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var requestBody entity.PengeluaranRequest

	helper.RequestBody(r, &requestBody)
	validate := validator.New()
	if err := validate.Struct(&requestBody); err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    err.Error(),
		}, http.StatusBadRequest)
		return
	}

	newPengeluaran, err := pc.P.CreatePengeluaran(r.Context(), requestBody)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    "Failed to create user",
		}, http.StatusInternalServerError)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusCreated,
		Message: "created",
		Data:    newPengeluaran,
	}

	helper.ResponseBody(w, response, http.StatusCreated)

}

func (pc PengeluaranController) GetPengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	allPengeluaran, err := pc.P.FindAllPengeluaran(r.Context())
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "not found",
			Data:    "Failed to create user",
		}, http.StatusNotFound)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    allPengeluaran,
	}

	helper.ResponseBody(w, response, http.StatusOK)

}

func (pc PengeluaranController) GetPengeluaranById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	pengeluaran, err := pc.P.FindPengeluaranById(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "not found",
			Data:    "Failed to create user",
		}, http.StatusNotFound)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    pengeluaran,
	}

	helper.ResponseBody(w, response, http.StatusOK)
}

func (pc PengeluaranController) DeletePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	err := pc.P.DeletePengeluaran(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "not found",
			Data:    "Failed to create user",
		}, http.StatusNotFound)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    "succes delete pengeluaran",
	}

	helper.ResponseBody(w, response, http.StatusOK)
}

func (pc PengeluaranController) UpdatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var requestBody entity.PengeluaranRequest

	helper.RequestBody(r, &requestBody)

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	editedPengeluaran, err := pc.P.EditPengeluaran(r.Context(), requestBody, idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    "Failed to create user",
		}, http.StatusInternalServerError)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "succses edit data pengeluaran",
		Data:    editedPengeluaran,
	}

	helper.ResponseBody(w, response, http.StatusOK)

}
