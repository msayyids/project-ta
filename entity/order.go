package entity

import (
	"time"
)

type Order struct {
	Id                   int       `json:"id"`
	Nama_pelanggan       string    `json:"nama_pelanggan"`
	No_Telepon_Pelanggan string    `json:"no_telepon_pelanggan"`
	Layanan_id           int       `json:"layanan_id"`
	User_id              int       `json:"user_id"`
	Jumlah               int       `json:"jumlah"`
	Tanggal_order        time.Time `json:"tanggal_order"`
	Total                int       `json:"total"`
	Status               string    `json:"status"`
	Payment_type         string    `json:"payment_status"`
	Payment_url          string    `json:"payment_url"`
}

type OrderRequest struct {
	Nama_pelanggan       string `json:"nama_pelanggan"`
	No_Telepon_Pelanggan string `json:"no_telepon_pelanggan"`
	Layanan_id           int    `json:"layanan_id"`
	User_id              int    `json:"user_id"`
	Jumlah               int    `json:"jumlah"`
	Total                int    `json:"total"`
	Status               string `json:"status"`
	Payment_type         string `json:"payment_status"`
	Payment_url          string `json:"payment_url"`
}
