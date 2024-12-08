package entity

import "time"

// Model Pengeluaran menggunakan GORM
type Pengeluaran struct {
	Id                int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Nama_pengeluaran  string    `json:"nama_pengeluaran" gorm:"column:nama_pengeluaran" validate:"required"`
	Keterangan        string    `json:"keterangan" gorm:"column:keterangan" validate:"required"`
	Users_id          int       `json:"users_id" gorm:"column:users_id" validate:"required"`
	Total             int       `json:"total" gorm:"column:total" validate:"required"`
	Bukti_pengeluaran string    `json:"bukti_pengeluaran" gorm:"column:bukti_pengeluaran" validate:"required"`
	Created_at        time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

type PengeluaranRequest struct {
	Id                int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Nama_pengeluaran  string    `json:"nama_pengeluaran" gorm:"column:nama_pengeluaran" validate:"required"`
	Keterangan        string    `json:"keterangan" gorm:"column:keterangan" validate:"required"`
	Users_id          int       `json:"users_id" gorm:"column:users_id" validate:"required"`
	Total             int       `json:"total" gorm:"column:total" validate:"required"`
	Bukti_pengeluaran string    `json:"bukti_pengeluaran" gorm:"column:bukti_pengeluaran" validate:"required"`
	Created_at        time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
