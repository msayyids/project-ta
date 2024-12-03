package repository

type MidtransPaymentRepository struct{}

type MidtransPaymentRepositoryInj interface {
}

func NewMidtransPaymentRepository() MidtransPaymentRepositoryInj {
	return MidtransPaymentRepository{}
}
