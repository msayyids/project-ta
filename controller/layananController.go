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

type LayananControllerInj interface {
	CreateLayanan(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindLayananById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteLayananById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindAllLayanan(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	EditLayananById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type LayananController struct {
	LayananService service.LayananServiceInj
	V              validator.Validate
}

func NewLayananController(layananService service.LayananServiceInj, v validator.Validate) LayananControllerInj {
	return LayananController{LayananService: layananService, V: v}
}

func (lc LayananController) CreateLayanan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var layananRequest entity.LayananRequest

	helper.RequestBody(r, &layananRequest)

	err := lc.V.Struct(layananRequest)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    "INVALID INPUT",
		}, http.StatusBadRequest)
		return
	}

	newLayanan, err := lc.LayananService.AddLayanan(r.Context(), layananRequest)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    500,
			Message: "INTERNAL SERVER ERROR",
			Data:    "FAILED CREATE LAYANAN",
		}, http.StatusInternalServerError)
		return
	}

	response := entity.WebResponse{
		Code:    201,
		Message: "SUCCESS CREATE LAYANAN",
		Data:    newLayanan,
	}

	helper.ResponseBody(w, response, http.StatusCreated)
}

func (lc LayananController) FindLayananById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	layanan, err := lc.LayananService.FindLayananById(r.Context(), ids) // Use ids here
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    nil,
		}, http.StatusNotFound)
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    layanan,
	}, http.StatusOK)
}

func (lc LayananController) DeleteLayananById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = lc.LayananService.DeleteLayananById(r.Context(), ids) // Use ids here
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "INTERNAL SERVER ERROR",
			Data:    nil,
		}, http.StatusNotFound)
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    200,
		Message: "SUCCESS DELETE LAYANAN",
		Data:    nil,
	}, http.StatusOK)
}

func (lc LayananController) FindAllLayanan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	layanan, err := lc.LayananService.FindAllLayanan(r.Context())
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    500,
			Message: "INTERNAL SERVER ERROR",
			Data:    "FAILED FETCH LAYANAN",
		}, http.StatusInternalServerError)
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    layanan,
	}, http.StatusOK)
}

func (lc LayananController) EditLayananById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var layananRequest entity.LayananRequest
	helper.RequestBody(r, &layananRequest)

	layanan, err := lc.LayananService.EditLayananById(r.Context(), ids, layananRequest) // Use ids here
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    404,
			Message: "NOT FOUND",
			Data:    nil,
		}, http.StatusNotFound)
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    200,
		Message: "SUCCESS",
		Data:    layanan,
	}, http.StatusOK)
}
