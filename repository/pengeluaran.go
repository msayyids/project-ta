package repository

type PengeluaranRepository struct{}

type PengeluaranRepositoryInj interface{}

func NewPengeluaranRepository() PengeluaranRepositoryInj {
	return PengeluaranRepository{}
}
