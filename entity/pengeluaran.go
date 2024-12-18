package entity

import "time"

type Pengeluaran struct {
	Id                int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Nama_pengeluaran  string    `json:"nama_pengeluaran" gorm:"column:nama_pengeluaran"`
	Keterangan        string    `json:"keterangan" gorm:"column:keterangan"`
	Users_id          int       `json:"users_id" gorm:"column:users_id"`
	Total             int       `json:"total" gorm:"column:total"`
	Bukti_pengeluaran string    `json:"bukti_pengeluaran" gorm:"column:bukti_pengeluaran"`
	Tipe_pengeluaran  string    `json:"tipe_pengeluaran" gorm:"column:tipe_pengeluaran"`
	Created_at        time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

type PengeluaranRequest struct {
	Nama_pengeluaran  string `json:"nama_pengeluaran" form:"nama_pengeluaran" validate:"required"`
	Keterangan        string `json:"keterangan" form:"keterangan" validate:"required"`
	Users_id          int    `json:"users_id" form:"users_id" validate:"required"`
	Total             int    `json:"total" form:"total" validate:"required"`
	Bukti_pengeluaran string `json:"bukti_pengeluaran" form:"bukti_pengeluaran" validate:"required"`
	Tipe_pengeluaran  string `json:"tipe_pengeluaran" form:"tipe_pengeluaran" validate:"required"`
}

func (Pengeluaran) TableName() string {
	return "pengeluaran"
}
