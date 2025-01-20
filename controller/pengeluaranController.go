package controller

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"net/http"
	"os"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type PengeluaranController struct {
	P service.PengeluaranServiceInj
	C *cloudinary.Cloudinary
}

type PengeluaranControllerInj interface {
	CreatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetPengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetPengeluaranById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	UpdatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	DeletePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindPengeluaranByDate(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewPengeluaranController(p service.PengeluaranServiceInj, c *cloudinary.Cloudinary) PengeluaranController {
	return PengeluaranController{
		P: p,
		C: c,
	}
}

func (pc PengeluaranController) CreatePengeluaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req entity.PengeluaranRequest

	helper.RequestBody(r, &req)

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    "INVALID INPUT",
		}, http.StatusBadRequest)
		return
	}

	fileBytes, err := os.Open(req.Bukti_pengeluaran)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	resp, err := pc.C.Upload.Upload(r.Context(), fileBytes, uploader.UploadParams{})
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Data:    "Failed to upload file",
		}, http.StatusInternalServerError)
		return
	}

	req.Bukti_pengeluaran = resp.SecureURL

	// Create Pengeluaran
	newPengeluaran, err := pc.P.CreatePengeluaran(r.Context(), req)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Data:    "Failed to create pengeluaran",
		}, http.StatusInternalServerError)
		return
	}

	// Send success response
	response := entity.WebResponse{
		Code:    http.StatusCreated,
		Message: "created",
		Data:    newPengeluaran,
	}
	helper.ResponseBody(w, response, http.StatusCreated)
}

func (pc PengeluaranController) GetPengeluaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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

func (pc *PengeluaranController) UpdatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var requestBody entity.PengeluaranRequest

	helper.RequestBody(r, &requestBody)

	// Ambil ID dari parameter URL
	id := param.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID parameter",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	if requestBody.Bukti_pengeluaran != "" {
		filePath := requestBody.Bukti_pengeluaran
		fileBytes, err := os.Open(filePath)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			}, http.StatusBadRequest)
			return
		}

		uploadResult, err := pc.C.Upload.Upload(r.Context(), fileBytes, uploader.UploadParams{})
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			}, http.StatusBadRequest)
			return
		}
		requestBody.Bukti_pengeluaran = uploadResult.SecureURL
	}

	// Panggil service untuk memperbarui pengeluaran
	editedPengeluaran, err := pc.P.EditPengeluaran(r.Context(), requestBody, idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// Kirim response sukses
	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "Success edit pengeluaran",
		Data:    editedPengeluaran,
	}

	helper.ResponseBody(w, response, http.StatusOK)
}

func (pc PengeluaranController) FindPengeluaranByDate(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	dateStr := param.ByName("tanggal")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	pengeluarans, err := pc.P.GetPengeluaranByDate(r.Context(), date)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "not found",
			Data:    nil,
		}, http.StatusNotFound)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "success find pengeluaran",
		Data:    pengeluarans,
	}

	helper.ResponseBody(w, response, http.StatusOK)

}
