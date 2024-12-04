package entity

type Layanan struct {
	ID        int    `json:"id" db:"id"`
	Nama      string `json:"nama" db:"nama"`
	Deskripsi string `json:"deskripsi" db:"deskripsi"`
	Harga     int    `json:"harga" db:"harga"`
}

type LayananRequest struct {
	Nama       string `json:"nama"`
	Desksripsi string `json:"deskripsi"`
	Harga      int    `json:"harga"`
}
