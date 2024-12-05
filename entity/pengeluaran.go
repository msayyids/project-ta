package entity

import "time"

type Pengeluaran struct {
	Id                int       `json:"id" db:"id"`
	Nama_pengeluaran  string    `json:"nama_pengeluaran" db:"nama_pengeluaran"`
	Keterangan        string    `json:"keterangan" db:"keterangan"`
	Users_id          int       `json:"users_id" db:"users_id"`
	Total             int       `json:"total" db:"total"`
	Bukti_pengeluaran string    `json:"bukti_pengeluaran" db:"bukti_pengeluaran"`
	Created_at        time.Time `json:"created_at" db:"created_at"`
}

type PengeluaranRequest struct {
	Nama_pengeluaran  string    `json:"nama_pengeluaran"`
	Keterangan        string    `json:"keterangan"`
	Users_id          int       `json:"users_id"`
	Total             int       `json:"total"`
	Bukti_pengeluaran string    `json:"bukti_pengeluaran"`
	Created_at        time.Time `json:"created_at"`
}
