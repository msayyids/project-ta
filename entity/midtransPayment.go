package entity

import "time"

type MidtransPayment struct {
	ID          int       `db:"id"`
	OrderID     int       `db:"order_id" json:"order_id"`
	RedirectURL string    `db:"redirect_url" json:"redirect_url"`
	SubTotal    int       `db:"subtotal" json:"sub_total"`
	PaymentDate time.Time `db:"payment_date" json:"payment_date"`
	Status      string    `db:"status" json:"status" `
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type MidtransPaymentRequest struct {
	OrderID     int       `json:"order_id"`
	SubTotal    int       `json:"sub_total"`
	RedirectURL string    `json:"redirect_url"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
}

type ReqId struct {
	OrderID int `json:"order_id"`
}
