package entity

type Transaksi struct {
	Id              int    `json:"id"`
	Order_id        int    `json:"order_id"`
	Jenis_transaksi int    `json:"jenis_transaksi"`
	Subtotal        int    `json:"subtotal"`
	Status          string `json:"status"`
}
