package entity

import "time"

type Pengeluaran struct {
	Id                int       `json:"id" db:"id"`
	Nama_pengeluaran  string    `json:"nama_pengeluaran" db:"nama_pengeluaran" validate:"required"`
	Keterangan        string    `json:"keterangan" db:"keterangan" validate:"required"`
	Users_id          int       `json:"users_id" db:"users_id" validate:"required"`
	Total             int       `json:"total" db:"total" validate:"required"`
	Bukti_pengeluaran string    `json:"bukti_pengeluaran" db:"bukti_pengeluaran" validate:"required"`
	Created_at        time.Time `json:"created_at" db:"created_at"`
}

type PengeluaranRequest struct {
	Nama_pengeluaran  string    `json:"nama_pengeluaran" validate:"required"`
	Keterangan        string    `json:"keterangan" validate:"required"`
	Users_id          int       `json:"users_id" validate:"required"`
	Total             int       `json:"total" validate:"required,gte=0"`
	Bukti_pengeluaran string    `json:"bukti_pengeluaran" validate:"required"`
	Created_at        time.Time `json:"created_at"`
}
