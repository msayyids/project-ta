package controller

import (
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"net/http"
	"project-ta/config"
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
}

type PengeluaranControllerInj interface {
	CreatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetPengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetPengeluaranById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	UpdatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	DeletePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindPengeluaranByDate(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewPengeluaranController(p service.PengeluaranServiceInj) PengeluaranControllerInj {
	return PengeluaranController{
		P: p,
	}
}

func (pc PengeluaranController) CreatePengeluaran(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	cld := config.InitializeCloudinary()

	file, _, err := r.FormFile("bukti_pengeluaran")
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    "Bukti_pengeluaran is required and must be a valid file",
		}, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload file to Cloudinary
	resp, err := cld.Upload.Upload(r.Context(), file, uploader.UploadParams{})
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Data:    "Failed to upload file",
		}, http.StatusInternalServerError)
		return
	}

	// Convert form values
	usersID, err := strconv.Atoi(r.FormValue("users_id"))
	if err != nil {
		http.Error(w, "Invalid users_id", http.StatusBadRequest)
		return
	}

	total, err := strconv.Atoi(r.FormValue("total"))
	if err != nil {
		http.Error(w, "Invalid total", http.StatusBadRequest)
		return
	}

	// Create request object
	request := entity.PengeluaranRequest{
		Nama_pengeluaran:  r.FormValue("nama_pengeluaran"),
		Keterangan:        r.FormValue("keterangan"),
		Users_id:          usersID,
		Total:             total,
		Bukti_pengeluaran: resp.SecureURL, // Use SecureURL for HTTPS links
		Tipe_pengeluaran:  r.FormValue("tipe_pengeluaran"),
	}

	// Validate request object
	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Create Pengeluaran
	newPengeluaran, err := pc.P.CreatePengeluaran(r.Context(), request)
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
