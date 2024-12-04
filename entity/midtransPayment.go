package entity

import "time"

type MidtransPayment struct {
	ID          int       `db:"id"`
	OrderID     int       `db:"order_id"`
	RedirectURL string    `db:"redirect_url"`
	SubTotal    int       `db:"subtotal"`
	PaymentDate time.Time `db:"payment_date"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
}

type MidtransPaymentRequest struct {
	OrderID     int       `json:"order_id"`
	RedirectURL string    `json:"redirect_url"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
}
