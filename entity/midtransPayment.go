package entity

type MidtransPayment struct {
	Id          int    `json:"id"`
	Order_id    int    `json:"order_id"`
	RedirectUrl string `json:"redirect_url"`
}
