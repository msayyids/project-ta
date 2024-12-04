package repository

import (
	"context"
	"project-ta/entity"

	"github.com/jmoiron/sqlx"
)

type MidtransPaymentRepository struct{}

type MidtransPaymentRepositoryInj interface {
	AddPayment(ctx context.Context, paymentRequest entity.MidtransPaymentRequest, tx sqlx.Tx) (entity.MidtransPayment, error)
	// FindPayment(ctx context.Context, tx sqlx.Tx) ([]entity.MidtransPayment, error)
	// FindPaymentById(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error)
	// EditPayment(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error)
	// Deletepayment(ctx context.Context, id int, tx sqlx.Tx) error
}

func NewMidtransPaymentRepository() MidtransPaymentRepositoryInj {
	return MidtransPaymentRepository{}
}

func (p MidtransPaymentRepository) AddPayment(ctx context.Context, paymentRequest entity.MidtransPaymentRequest, tx sqlx.Tx) (entity.MidtransPayment, error) {
	sqlQuery := `INSERT INTO payment () VALUES()`

	var payment entity.MidtransPayment

	err := tx.QueryRowxContext(ctx, sqlQuery, paymentRequest).StructScan(payment)
	if err != nil {
		return entity.MidtransPayment{}, err
	}

	return entity.MidtransPayment{}, nil

}

// func (p MidtransPaymentRepository) FindPayment(ctx context.Context, tx sqlx.Tx) ([]entity.MidtransPayment, error) {
// 	sqlQuery := `INSERT INTO payment () VALUES()`

// }

// func (p MidtransPaymentRepository) FindPaymentById(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error) {
// 	sqlQuery := `INSERT INTO payment () VALUES()`

// }

// func (p MidtransPaymentRepository) EditPayment(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error) {
// 	sqlQuery := `INSERT INTO payment () VALUES()`

// }

// func (p MidtransPaymentRepository) Deletepayment(ctx context.Context, id int, tx sqlx.Tx) error {
// 	sqlQuery := `INSERT INTO payment () VALUES()`

// }
