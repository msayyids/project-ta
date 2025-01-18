package entity

import "time"

type KeuntunganResponse struct {
	Tanggal          time.Time `json:"tanggal"`
	TotalPemasukan   int       `json:"total_pemasukan"`
	TotalPengeluaran int       `json:"total_pengeluaran"`
	Keuntungan       int       `json:"keuntungan"`
	StatusKeuntungan string    `json:"status_keuntungan"`
}

type KeuntunganResponseMonthly struct {
	Tahun            int    `json:"tahun"`
	Bulan            int    `json:"bulan"`
	TotalPemasukan   int    `json:"total_pemasukan"`
	TotalPengeluaran int    `json:"total_pengeluaran"`
	Keuntungan       int    `json:"keuntungan"`
	StatusKeuntungan string `json:"status_keuntungan"`
}

type KeuntunganPer7HariResponse struct {
	Tanggal          string `json:"tanggal"`
	TotalPemasukan   int    `json:"total_pemasukan"`
	TotalPengeluaran int    `json:"total_pengeluaran"`
	Keuntungan       int    `json:"keuntungan"`
	StatusKeuntungan string `json:"status_keuntungan"`
}
