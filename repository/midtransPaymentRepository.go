package repository

import (
	"context"
	"fmt"
	"project-ta/entity"

	"github.com/jmoiron/sqlx"
)

type MidtransPaymentRepository struct{}

type MidtransPaymentRepositoryInj interface {
	AddPayment(ctx context.Context, payment entity.MidtransPayment, tx sqlx.Tx) (entity.MidtransPayment, error)
	FindPayment(ctx context.Context, tx sqlx.Tx) ([]entity.MidtransPayment, error)
	FindPaymentById(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error)
	EditPayment(ctx context.Context, id int, paymentRequest entity.MidtransPaymentRequest, tx sqlx.Tx) (entity.MidtransPayment, error)
	DeletePayment(ctx context.Context, id int, tx sqlx.Tx) error
}

func NewMidtransPaymentRepository() MidtransPaymentRepositoryInj {
	return MidtransPaymentRepository{}
}

func (repo MidtransPaymentRepository) AddPayment(ctx context.Context, payment entity.MidtransPayment, tx sqlx.Tx) (entity.MidtransPayment, error) {
	query := `
        INSERT INTO payments (order_id, redirect_url, subtotal, payment_date, status, created_at) 
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, order_id, redirect_url, subtotal, payment_date, status, created_at
    `

	var result entity.MidtransPayment
	err := tx.QueryRowxContext(ctx, query,
		payment.OrderID, payment.RedirectURL, payment.SubTotal, payment.PaymentDate, payment.Status, payment.CreatedAt,
	).StructScan(&result)

	if err != nil {
		return entity.MidtransPayment{}, fmt.Errorf("failed to insert payment: %v", err)
	}

	return result, nil
}

func (p MidtransPaymentRepository) FindPayment(ctx context.Context, tx sqlx.Tx) ([]entity.MidtransPayment, error) {
	sqlQuery := `SELECT id, order_id, redirect_url, payment_date, status, created_at
                 FROM payments`

	var payments []entity.MidtransPayment
	err := tx.SelectContext(ctx, &payments, sqlQuery)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (p MidtransPaymentRepository) FindPaymentById(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error) {
	sqlQuery := `SELECT id, order_id, redirect_url, payment_date, status, created_at
                 FROM payments
                 WHERE id = $1`

	var payment entity.MidtransPayment
	err := tx.QueryRowxContext(ctx, sqlQuery, id).StructScan(&payment)
	if err != nil {
		return entity.MidtransPayment{}, err
	}

	return payment, nil
}

func (p MidtransPaymentRepository) EditPayment(ctx context.Context, id int, paymentRequest entity.MidtransPaymentRequest, tx sqlx.Tx) (entity.MidtransPayment, error) {
	sqlQuery := `UPDATE payments
                 SET order_id = $1, redirect_url = $2, payment_date = $3, status = $4
                 WHERE id = $5
                 RETURNING id, order_id, redirect_url, payment_date, status, created_at`

	var payment entity.MidtransPayment
	err := tx.QueryRowxContext(ctx, sqlQuery,
		paymentRequest.OrderID,
		paymentRequest.RedirectURL,
		paymentRequest.PaymentDate,
		paymentRequest.Status,
		id,
	).StructScan(&payment)
	if err != nil {
		return entity.MidtransPayment{}, err
	}

	return payment, nil
}

func (p MidtransPaymentRepository) DeletePayment(ctx context.Context, id int, tx sqlx.Tx) error {
	sqlQuery := `DELETE FROM payments WHERE id = $1`

	_, err := tx.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}
