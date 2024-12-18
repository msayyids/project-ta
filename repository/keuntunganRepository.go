package repository

import (
	"context"
	"gorm.io/gorm"
	"project-ta/entity"
	"time"
)

type KeuntunganRepository struct{}

// Pastikan interface ini sesuai dengan metode yang ada di struct
type KeuntunganRepositoryInj interface {
	GetKeuntunganByDate(ctx context.Context, db *gorm.DB, date time.Time) ([]entity.KeuntunganResponse, error)
	GetKeuntunganByMonth(ctx context.Context, db *gorm.DB, year, month int) ([]entity.KeuntunganResponseMonthly, error)
	GetKeuntunganByLast7Days(ctx context.Context, db *gorm.DB, date time.Time) ([]entity.KeuntunganPer7HariResponse, error)
}

// Fungsi konstruktor
func NewKeuntunganRepository() KeuntunganRepositoryInj {
	return &KeuntunganRepository{}
}

// Implementasi metode GetKeuntunganByDate
func (r *KeuntunganRepository) GetKeuntunganByDate(ctx context.Context, db *gorm.DB, date time.Time) ([]entity.KeuntunganResponse, error) {
	var keuntungan []entity.KeuntunganResponse

	// Truncate waktu untuk memastikan hanya tanggal yang diperhitungkan, tanpa waktu
	startOfDay := date.Truncate(24 * time.Hour) // Jam diatur ke 00:00:00
	endOfDay := startOfDay.Add(24 * time.Hour)  // Jam diatur ke 23:59:59

	// Query untuk mendapatkan keuntungan pada satu hari tertentu
	err := db.Table("orders o").
		Select("DATE(o.tanggal_order) AS tanggal, "+
			"SUM(o.total) AS total_pemasukan, "+
			"SUM(COALESCE(ex.total, 0)) AS total_pengeluaran, "+
			"SUM(o.total) - SUM(COALESCE(ex.total, 0)) AS keuntungan, "+
			"CASE WHEN SUM(o.total) - SUM(COALESCE(ex.total, 0)) < 0 THEN 'MINUS' ELSE 'PLUS' END AS status_keuntungan").
		Joins("JOIN payments p ON o.id = p.order_id").
		// Gabungkan dengan pengeluaran berdasarkan tanggal
		Joins("LEFT JOIN pengeluaran ex ON DATE(o.tanggal_order) = DATE(ex.created_at)").
		Where("o.status = ? AND p.status = ? AND o.tanggal_order BETWEEN ? AND ?", "PAID", "PAID", startOfDay, endOfDay).
		Group("DATE(o.tanggal_order)").
		Order("tanggal DESC").
		Scan(&keuntungan).Error

	if err != nil {
		return nil, err
	}

	return keuntungan, nil
}

// Implementasi metode GetKeuntunganByMonth
func (r *KeuntunganRepository) GetKeuntunganByMonth(ctx context.Context, db *gorm.DB, year, month int) ([]entity.KeuntunganResponseMonthly, error) {
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
		// Gabungkan dengan pengeluaran berdasarkan bulan
		Joins("LEFT JOIN pengeluaran ex ON EXTRACT(MONTH FROM o.tanggal_order) = EXTRACT(MONTH FROM ex.created_at) AND EXTRACT(YEAR FROM o.tanggal_order) = EXTRACT(YEAR FROM ex.created_at)").
		Where("o.status = ? AND p.status = ? AND o.tanggal_order BETWEEN ? AND ?", "PAID", "PAID", startOfMonth, endOfMonth).
		Group("EXTRACT(YEAR FROM o.tanggal_order), EXTRACT(MONTH FROM o.tanggal_order)").
		Order("tahun DESC, bulan DESC").
		Scan(&keuntungan).Error

	if err != nil {
		return nil, err
	}

	// Pastikan hasil yang dikembalikan sesuai dengan bulan yang diminta
	return keuntungan, nil
}

// Implementasi metode GetKeuntunganByLast7Days
func (r *KeuntunganRepository) GetKeuntunganByLast7Days(ctx context.Context, db *gorm.DB, date time.Time) ([]entity.KeuntunganPer7HariResponse, error) {
	var keuntungan []entity.KeuntunganPer7HariResponse

	// Menghitung tanggal 7 hari yang lalu
	sevenDaysAgo := date.AddDate(0, 0, -7)

	err := db.Table("orders o").
		Select("DATE(o.tanggal_order) AS tanggal, "+
			"SUM(o.total) AS total_pemasukan, "+
			"SUM(COALESCE(ex.total, 0)) AS total_pengeluaran, "+
			"SUM(o.total) - SUM(COALESCE(ex.total, 0)) AS keuntungan, "+
			"CASE WHEN SUM(o.total) - SUM(COALESCE(ex.total, 0)) < 0 THEN 'MINUS' ELSE 'PLUS' END AS status_keuntungan").
		Joins("JOIN payments p ON o.id = p.order_id").
		// Gabungkan dengan pengeluaran berdasarkan tanggal
		Joins("LEFT JOIN pengeluaran ex ON DATE(o.tanggal_order) = DATE(ex.created_at)").
		Where("o.status = ? AND p.status = ? AND o.tanggal_order BETWEEN ? AND ?", "PAID", "PAID", sevenDaysAgo, date).
		Group("DATE(o.tanggal_order)").
		Order("tanggal DESC").
		Scan(&keuntungan).Error

	if err != nil {
		return nil, err
	}

	return keuntungan, nil
}
