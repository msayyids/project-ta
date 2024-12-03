package entity

import "time"

type Order struct {
	Id                   int       `json:"id"`
	Nama_pelanggaan      string    `json:"nama_pelanggan"`
	No_Telepon_Pelanggan string    `json:"no_telepon_pelanggan"`
	Layanan_id           int       `json:"layanan_id"`
	User_id              int       `json:"user_id"`
	Jumlah               int       `json:"jumlah"`
	Tanggal_order        time.Time `json:"tanggal_order"`
	Total                int       `json:"total"`
	Status               string    `json:"status"`
	Tanggal_update       time.Time `json:"Tanggal_update"`
}

type OrderRequest struct {
	Nama_pelanggaan      string    `json:"nama_pelanggan"`
	No_Telepon_Pelanggan string    `json:"no_telepon_pelanggan"`
	Layanan_id           int       `json:"layanan_id"`
	User_id              int       `json:"user_id"`
	Jumlah               int       `json:"jumlah"`
	Tanggal_order        time.Time `json:"-"`
	Total                int       `json:"total"`
	Status               string    `json:"status"`
	Tanggal_update       time.Time `json:"-"`
}
