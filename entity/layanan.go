package entity

type Layanan struct {
	ID        int    `json:"id" db:"id"`
	Nama      string `json:"nama" db:"nama" validate:"required"`
	Deskripsi string `json:"deskripsi" db:"deskripsi"`
	Harga     int    `json:"harga" db:"harga"`
}

type LayananRequest struct {
	Nama       string `json:"nama" validate:"required"`
	Desksripsi string `json:"deskripsi"`
	Harga      int    `json:"harga"`
}

func (Layanan) TableName() string {
	return "layanan"
}
