package entity

import "time"

type Order struct {
	ID                 int       `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaPelanggan      string    `json:"nama_pelanggan" gorm:"type:varchar(255);not null "`
	NoTeleponPelanggan string    `json:"no_telepon_pelanggan,omitempty" gorm:"type:varchar(15)"`
	LayananID          int       `json:"layanan_id" gorm:"not null"`
	UserID             int       `json:"user_id" gorm:"not null"`
	Jumlah             int       `json:"jumlah" gorm:"not null" validate:"required"`
	TanggalOrder       time.Time `json:"tanggal_order" gorm:"default:CURRENT_TIMESTAMP"`
	Total              int       `json:"total" gorm:"not null"`
	Status             string    `json:"status" gorm:"type:varchar(50);default:'UNPAID'"`
	PaymentType        string    `json:"payment_type" gorm:"type:varchar(50);not null"`
	Payment_url        string    `json:"payment_url" gorm:"type:varchar"`
}

type OrderReq struct {
	NamaPelanggan      string    `json:"nama_pelanggan" gorm:"type:varchar(255);not null" validate:"required"`
	NoTeleponPelanggan string    `json:"no_telepon_pelanggan,omitempty" gorm:"type:varchar(15)"`
	LayananID          int       `json:"layanan_id" gorm:"not null"`
	UserID             int       `json:"user_id" gorm:"not null"`
	Jumlah             int       `json:"jumlah" gorm:"not null"`
	TanggalOrder       time.Time `json:"tanggal_order" gorm:"default:CURRENT_TIMESTAMP"`
	Total              int       `json:"total" gorm:"not null"`
	Status             string    `json:"status" gorm:"type:varchar(50);default:'UNPAID'"`
	PaymentType        string    `json:"payment_type" gorm:"type:varchar(50);not null"`
}
