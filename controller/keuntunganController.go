package controller

import (
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type KeuntunganController struct {
	DB *gorm.DB
}

func NewKeuntunganCntroller(db gorm.DB) KeuntunganController {
	return KeuntunganController{
		DB: &db,
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

	// Ambil data keuntungan dalam rentang 7 hari terakhir
	keuntungan, err := GetKeuntunganByLast7Days(c.DB, date)
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

	keuntungan, err := GetKeuntunganByMonth(c.DB, year, month)
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

	keuntungan, err := GetKeuntunganByDate(c.DB, date)
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
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    keuntungan,
	}, http.StatusOK)
}

func GetKeuntunganByDate(db *gorm.DB, date time.Time) ([]entity.KeuntunganResponse, error) {
	var keuntungan []entity.KeuntunganResponse

	startOfDay := date.Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := db.Table("orders o").
		Select("DATE(o.tanggal_order) AS tanggal, SUM(o.total) AS total_pemasukan, SUM(COALESCE(ex.total, 0)) AS total_pengeluaran, SUM(o.total) - SUM(COALESCE(ex.total, 0)) AS keuntungan, CASE WHEN SUM(o.total) - SUM(COALESCE(ex.total, 0)) < 0 THEN 'MINUS' ELSE 'PLUS' END AS status_keuntungan").
		Joins("JOIN payments p ON o.id = p.order_id").
		Joins("LEFT JOIN pengeluaran ex ON 1 = 1").
		Where("o.status = ? AND p.status = ? AND o.tanggal_order BETWEEN ? AND ?", "PAID", "PAID", startOfDay, endOfDay).
		Group("DATE(o.tanggal_order)").
		Order("tanggal DESC").
		Scan(&keuntungan).Error

	if err != nil {
		return nil, err
	}

	return keuntungan, nil
}

func GetKeuntunganByMonth(db *gorm.DB, year, month int) ([]entity.KeuntunganResponseMonthly, error) {
	var keuntungan []entity.KeuntunganResponseMonthly

	// Menentukan start dan end date untuk bulan yang diminta
	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second) // Menghitung akhir bulan

	err := db.Table("orders o").
		Select("EXTRACT(YEAR FROM o.tanggal_order) AS tahun, EXTRACT(MONTH FROM o.tanggal_order) AS bulan, "+
			"SUM(o.total) AS total_pemasukan, SUM(COALESCE(ex.total, 0)) AS total_pengeluaran, "+
			"SUM(o.total) - SUM(COALESCE(ex.total, 0)) AS keuntungan, "+
			"CASE WHEN SUM(o.total) - SUM(COALESCE(ex.total, 0)) < 0 THEN 'MINUS' ELSE 'PLUS' END AS status_keuntungan").
		Joins("JOIN payments p ON o.id = p.order_id").
		Joins("LEFT JOIN pengeluaran ex ON 1 = 1"). // Ganti LeftJoin dengan Joins dan tambahkan LEFT
		Where("o.status = ? AND p.status = ? AND o.tanggal_order BETWEEN ? AND ?", "PAID", "PAID", startOfMonth, endOfMonth).
		Group("EXTRACT(YEAR FROM o.tanggal_order), EXTRACT(MONTH FROM o.tanggal_order)").
		Order("tahun DESC, bulan DESC").
		Scan(&keuntungan).Error

	if err != nil {
		return nil, err
	}

	return keuntungan, nil
}

func GetKeuntunganByLast7Days(db *gorm.DB, date time.Time) ([]entity.KeuntunganPer7HariResponse, error) {
	var keuntungan []entity.KeuntunganPer7HariResponse

	// Menghitung tanggal 7 hari yang lalu
	sevenDaysAgo := date.AddDate(0, 0, -7)

	// Query GORM untuk menghitung total pemasukan, total pengeluaran, dan keuntungan dalam rentang 7 hari
	err := db.Table("orders o").
		Select("DATE(o.tanggal_order) AS tanggal, "+
			"SUM(o.total) AS total_pemasukan, "+
			"SUM(COALESCE(ex.total, 0)) AS total_pengeluaran, "+
			"SUM(o.total) - SUM(COALESCE(ex.total, 0)) AS keuntungan, "+
			"CASE WHEN SUM(o.total) - SUM(COALESCE(ex.total, 0)) < 0 THEN 'MINUS' ELSE 'PLUS' END AS status_keuntungan").
		Joins("JOIN payments p ON o.id = p.order_id").
		Joins("LEFT JOIN pengeluaran ex ON 1 = 1").
		Where("o.status = ? AND p.status = ? AND o.tanggal_order BETWEEN ? AND ?", "PAID", "PAID", sevenDaysAgo, date).
		Group("DATE(o.tanggal_order)").
		Order("tanggal DESC").
		Scan(&keuntungan).Error

	if err != nil {
		return nil, err
	}

	return keuntungan, nil
}
