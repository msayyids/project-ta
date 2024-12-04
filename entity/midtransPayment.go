package entity

import "time"

type MidtransPayment struct {
	Id          int       `json:"id"`
	Order_id    int       `json:"order_id"`
	Amount      int       `json:"amount"`
	RedirectUrl string    `json:"redirect_url"`
	PaymentTime time.Time `json:"payment_time"`
	Status      string    `json:"status"`
}

type MidtransPaymentRequest struct {
	Order_id int `json:"order_id"`
	Amount   int `json:"amount"`
}
