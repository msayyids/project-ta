package entity

type Order struct {
	Id              int    `json:"id"`
	Nama_pelanggaan string `json:"nama_pelanggan"`
	Layanan_id      int    `json:"layanan_id"`
	User_id         int    `json:"user_id"`
	Jumlah          int    `json:"jumlah"`
	Tanggal_order   int    `json:"tanggal_order"`
	Total           int    `json:"total"`
	Status          string `json:"status"`
}
