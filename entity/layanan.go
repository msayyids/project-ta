package entity

type Layanan struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Desksripsi string `json:"deskripsi"`
	Harga      int    `json:"harga"`
}

type LayananRequest struct {
	Nama       string `json:"nama"`
	Desksripsi string `json:"deskripsi"`
	Harga      int    `json:"harga"`
}
