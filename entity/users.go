package entity

import (
	"time"
)

type Users struct {
	Id            int       `json:"id"`
	Nama_depan    string    `json:"nama_depan" validate:"required,min=2,max=50"`
	Nama_belakang string    `json:"nama_belakang" validate:"required,min=2,max=50"`
	Role          string    `json:"role" validate:"required,oneof=admin karyawan"`
	Email         string    `json:"email" validate:"required,email"`
	Password      string    `json:"password" validate:"required,min=8"`
	No_telepon    string    `json:"no_telepon" validate:"required,e164"`
	Alamat        string    `json:"alamat" validate:"required"`
	Gaji          int       `json:"gaji" validate:"required,gte=0"`
	No_rekening   string    `json:"no_rekening" validate:"required"`
	Bank_id       int       `json:"bank_id" validate:"required"`
	Created_At    time.Time `json:"created_at"`
}

type UserRequest struct {
	Nama_depan    string `json:"nama_depan" validate:"required,min=2,max=50"`
	Nama_belakang string `json:"nama_belakang" validate:"required,min=2,max=50"`
	Role          string `json:"role" validate:"required,oneof=admin karyawan"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required,min=8"`
	No_telepon    string `json:"no_telepon" validate:"required,e164"`
	Alamat        string `json:"alamat" validate:"required"`
	Gaji          int    `json:"gaji" validate:"required,gte=0"`
	No_rekening   string `json:"no_rekening" validate:"required"`
	Bank_id       int    `json:"bank_id" validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
