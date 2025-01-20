package entity

import (
	"time"
)

type Users struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama_depan    string    `gorm:"type:varchar(255);not null" json:"nama_depan"`
	Nama_belakang string    `gorm:"type:varchar(255);not null" json:"nama_belakang"`
	Role          string    `gorm:"type:role_type;not null" json:"role"`
	Email         string    `gorm:"type:varchar(255);unique;not null" json:"email" validate:"required,min=2,max=50"`
	Password      string    `gorm:"type:varchar(255);not null" json:"-"`
	No_telepon    string    `gorm:"type:varchar(12);not null" json:"no_telepon"`
	Alamat        string    `gorm:"type:varchar(255);not null" json:"alamat"`
	Gaji          int       `gorm:"not null" json:"gaji"`
	No_rekening   string    `gorm:"type:varchar(255);default:null" json:"no_rekening"`
	Bank_id       int       `gorm:"default:null" json:"bank_id"`
	Created_At    time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

type UserRequest struct {
	Nama_depan    string `json:"nama_depan" validate:"required,min=2,max=50"`
	Nama_belakang string `json:"nama_belakang" validate:"required,min=2,max=50"`
	Role          string `json:"role" validate:"required,oneof=admin karyawan"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required"`
	No_telepon    string `json:"no_telepon" validate:"required"`
	Alamat        string `json:"alamat" validate:"required"`
	Gaji          int    `json:"gaji" validate:"required,gte=0"`
	No_rekening   string `json:"no_rekening" validate:"required"`
	Bank_id       int    `json:"bank_id" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
