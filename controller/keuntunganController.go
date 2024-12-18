package controller

import (
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type KeuntunganController struct {
	Ks service.KeuntunganServiceInj
}

func NewKeuntunganCntroller(ks service.KeuntunganServiceInj) KeuntunganController {
	return KeuntunganController{
		Ks: ks,
	}
}

func (c KeuntunganController) GetKeuntunganByLast7DaysEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	dateStr := ps.ByName("tanggal")

	// Parse string tanggal menjadi time.Time
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	keuntungan, err := c.Ks.GetKeuntunganByLast7Days(r.Context(), date)
	if err != nil {
		http.Error(w, "Error fetching keuntungan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    keuntungan,
	}, http.StatusOK)
}

func (c KeuntunganController) GetKeuntunganByMonthEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	yearStr := ps.ByName("tahun")
	monthStr := ps.ByName("bulan")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 1900 || year > 9999 {
		http.Error(w, "Invalid year format", http.StatusBadRequest)
		return
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		http.Error(w, "Invalid month format", http.StatusBadRequest)
		return
	}

	keuntungan, err := c.Ks.GetKeuntunganByMonth(r.Context(), year, month)
	if err != nil {
		http.Error(w, "Error fetching keuntungan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    keuntungan,
	}, http.StatusOK)
}

func (c KeuntunganController) GetKeuntunganByDateEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dateStr := ps.ByName("tanggal")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	keuntungan, err := c.Ks.GetKeuntunganByDate(r.Context(), date)
	if err != nil {
		http.Error(w, "Error fetching keuntungan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if keuntungan == nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    "DATA TIDAK DITEMUKAN",
		}, http.StatusNotFound)

		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    keuntungan,
	}, http.StatusOK)
}
