package controller

import (
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type LayananControllerInj interface {
	CreateLayanan(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindLayananById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	DeleteLayananById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindAllLayanan(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	EditLayananById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

type LayananController struct {
	Ls service.LayananServiceInj
}

func NewLayananController(ls service.LayananServiceInj) LayananControllerInj {
	return LayananController{Ls: ls}
}

func (lc LayananController) CreateLayanan(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	layananRequest := entity.LayananRequest{}
	helper.RequestBody(r, layananRequest)

	newLayanan, err := lc.Ls.AddLayanan(r.Context(), layananRequest)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   "FAILED CREATE LAYANAN",
		})
		return
	}

	response := entity.WebResponse{
		Code:   201,
		Status: "SUCCESS CREATE LAYANAN",
		Data:   newLayanan,
	}

	helper.ResponseBody(w, response)

}

func (lc LayananController) FindLayananById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	idParam := param.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   404,
			Status: "NOT FOUND",
			Data:   nil,
		})
		return
	}

	layanan, err := lc.Ls.FindLayananById(r.Context(), id)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   nil,
		})

		return
	}

	layananByIdResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   layanan,
	}

	helper.ResponseBody(w, layananByIdResponse)
}

func (lc LayananController) DeleteLayananById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	idParam := param.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   404,
			Status: "NOT FOUND",
			Data:   nil,
		})
		return
	}

	err = lc.Ls.DeleteLayananById(r.Context(), id)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   nil,
		})
		return
	}

	response := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "success delete layanan",
	}

	helper.ResponseBody(w, response)

}

func (lc LayananController) FindAllLayanan(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	allLayanan, err := lc.Ls.FindAllLayanan(r.Context())
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   404,
			Status: "NOT FOUND",
			Data:   nil,
		})
		return
	}

	helper.ResponseBody(w, allLayanan)

}

func (lc LayananController) EditLayananById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	idParam := param.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   404,
			Status: "NOT FOUND",
			Data:   nil,
		})
		return
	}

	layananRequest := entity.LayananRequest{}
	helper.RequestBody(r, layananRequest)

	editedLayanan, err := lc.Ls.EditLayananById(r.Context(), id, layananRequest)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   404,
			Status: "NOT FOUND",
			Data:   nil,
		})
	}

	response := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   editedLayanan,
	}

	helper.ResponseBody(w, response)
}
