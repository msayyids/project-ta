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
	EditPayment(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error)
	DeletePayment(ctx context.Context, id int, tx sqlx.Tx) error
	EditStatusPayment(ctx context.Context, id int, status string, tx sqlx.Tx) (entity.MidtransPayment, error)
	EditPaymentbyOrderId(ctx context.Context, id int, status string, tx sqlx.Tx) (entity.MidtransPayment, error)
}

func NewMidtransPaymentRepository() MidtransPaymentRepositoryInj {
	return MidtransPaymentRepository{}
}

func (repo MidtransPaymentRepository) AddPayment(ctx context.Context, payment entity.MidtransPayment, tx sqlx.Tx) (entity.MidtransPayment, error) {
	query := `WITH order_check AS (
    SELECT payment_type
    FROM orders
    WHERE id = $1
	)
	INSERT INTO payments (order_id, redirect_url, subtotal, payment_date, status, created_at)
	SELECT $1, $2, $3, $4, $5, $6
	FROM order_check
	WHERE payment_type = 'cashless'
	RETURNING id, order_id, redirect_url, subtotal, payment_date, status, created_at;
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
	sqlQuery := `SELECT *
                 FROM payments`

	var payments []entity.MidtransPayment
	err := tx.SelectContext(ctx, &payments, sqlQuery)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (p MidtransPaymentRepository) FindPaymentById(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error) {
	sqlQuery := `SELECT *
                 FROM payments
                 WHERE id = $1`

	var payment entity.MidtransPayment
	err := tx.QueryRowxContext(ctx, sqlQuery, id).StructScan(&payment)
	if err != nil {
		return entity.MidtransPayment{}, err
	}

	return payment, nil
}

//benerin

func (p MidtransPaymentRepository) EditPayment(ctx context.Context, id int, tx sqlx.Tx) (entity.MidtransPayment, error) {
	sqlQuery := `UPDATE payments
				SET order_id = $1
				WHERE id = $2
				RETURNING id, order_id, redirect_url, payment_date, status, created_at;
`

	var payment entity.MidtransPayment
	err := tx.QueryRowxContext(ctx, sqlQuery, id).StructScan(&payment)
	if err != nil {
		return entity.MidtransPayment{}, err
	}

	return payment, nil
}
func (p MidtransPaymentRepository) EditPaymentbyOrderId(ctx context.Context, id int, status string, tx sqlx.Tx) (entity.MidtransPayment, error) {
	sqlQuery := `UPDATE payments
				SET order_id = $1
				SET status = $2
				WHERE order_id = $1
				RETURNING id, order_id, redirect_url, payment_date, status, created_at;
`

	var payment entity.MidtransPayment
	err := tx.QueryRowxContext(ctx, sqlQuery, id, status).StructScan(&payment)
	if err != nil {
		return entity.MidtransPayment{}, err
	}

	return payment, nil
}

func (p MidtransPaymentRepository) EditStatusPayment(ctx context.Context, id int, status string, tx sqlx.Tx) (entity.MidtransPayment, error) {
	sqlQuery := `UPDATE payments
				SET status = $1
				WHERE id = $2
				RETURNING id, order_id, redirect_url, payment_date, status, created_at;
`

	var payment entity.MidtransPayment
	err := tx.QueryRowxContext(ctx, sqlQuery, status, id).StructScan(&payment)
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
