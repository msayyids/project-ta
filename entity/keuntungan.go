package entity

import "time"

type KeuntunganResponse struct {
	Tanggal          time.Time `json:"tanggal"`           // Tanggal transaksi
	TotalPemasukan   int       `json:"total_pemasukan"`   // Total pemasukan per hari
	TotalPengeluaran int       `json:"total_pengeluaran"` // Total pengeluaran per hari
	Keuntungan       int       `json:"keuntungan"`        // Keuntungan per hari
	StatusKeuntungan string    `json:"status_keuntungan"` // Status keuntungan: 'PLUS' atau 'MINUS'
}

type KeuntunganResponseMonthly struct {
	Tahun            int     `json:"tahun"`             // Tahun dari tanggal order
	Bulan            int     `json:"bulan"`             // Bulan dari tanggal order
	TotalPemasukan   float64 `json:"total_pemasukan"`   // Total pemasukan per bulan
	TotalPengeluaran float64 `json:"total_pengeluaran"` // Total pengeluaran per bulan
	Keuntungan       float64 `json:"keuntungan"`        // Keuntungan (total pemasukan - total pengeluaran)
	StatusKeuntungan string  `json:"status_keuntungan"` // Status keuntungan ("PLUS" atau "MINUS")
}

type KeuntunganPer7HariResponse struct {
	Tanggal          string  `json:"tanggal"`           // Tanggal dari order
	TotalPemasukan   float64 `json:"total_pemasukan"`   // Total pemasukan dalam periode
	TotalPengeluaran float64 `json:"total_pengeluaran"` // Total pengeluaran dalam periode
	Keuntungan       float64 `json:"keuntungan"`        // Keuntungan (total pemasukan - total pengeluaran)
	StatusKeuntungan string  `json:"status_keuntungan"` // Status keuntungan ("PLUS" atau "MINUS")
}
