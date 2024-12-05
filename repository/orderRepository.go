package repository

import (
	"context"
	"database/sql"
	"fmt"
	"project-ta/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct{}

type OrderRepositoryInj interface {
	AddOrder(ctx context.Context, orderReq entity.OrderRequest, tx *sqlx.Tx) (entity.Order, error)
	FindOrderById(ctx context.Context, id int, tx *sqlx.Tx) (entity.Order, error)
	UpdateOrderById(ctx context.Context, id int, orderReq entity.OrderRequest, tx *sqlx.Tx) (entity.Order, error)
	DeleteOrderById(ctx context.Context, id int, tx *sqlx.Tx) error
	FindOrder(ctx context.Context, tx *sqlx.Tx) ([]entity.Order, error)
}

func NewOrderRepository() OrderRepositoryInj {
	return &OrderRepository{}
}

func (o *OrderRepository) AddOrder(ctx context.Context, orderReq entity.OrderRequest, tx *sqlx.Tx) (entity.Order, error) {
	sqlQuery := `
    INSERT INTO orders (
        nama_pelanggan, 
        no_telepon_pelanggan, 
        layanan_id, 
        user_id, 
        jumlah, 
        tanggal_order, 
        total, 
        status
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8
    ) RETURNING *;
`

	var newOrder entity.Order
	err := tx.QueryRowxContext(ctx, sqlQuery,
		orderReq.Nama_pelanggan,
		orderReq.No_Telepon_Pelanggan,
		orderReq.Layanan_id,
		orderReq.User_id,
		orderReq.Jumlah,
		time.Now(),
		orderReq.Total,
		orderReq.Status,
	).StructScan(&newOrder)

	if err != nil {
		return entity.Order{}, fmt.Errorf("error repo")
	}

	return newOrder, nil
}

func (o *OrderRepository) FindOrder(ctx context.Context, tx *sqlx.Tx) ([]entity.Order, error) {
	sqlQuery := `SELECT * FROM orders ORDER BY tanggal_order DESC;`
	var allOrders []entity.Order

	err := tx.SelectContext(ctx, &allOrders, sqlQuery)
	if err != nil {
		return nil, err
	}

	return allOrders, nil
}

func (o *OrderRepository) FindOrderById(ctx context.Context, id int, tx *sqlx.Tx) (entity.Order, error) {
	if tx == nil {
		return entity.Order{}, fmt.Errorf("transaction is nil")
	}

	sqlQuery := `SELECT * FROM orders WHERE id = $1;`

	var order entity.Order
	err := tx.GetContext(ctx, &order, sqlQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Order{}, fmt.Errorf("order with ID %d not found", id)
		}
		return entity.Order{}, fmt.Errorf("error executing query: %v", err)
	}

	return order, nil
}

func (o *OrderRepository) UpdateOrderById(ctx context.Context, id int, orderReq entity.OrderRequest, tx *sqlx.Tx) (entity.Order, error) {
	sqlQuery := `
    INSERT INTO orders (
        nama_pelanggan, 
        no_telepon_pelanggan, 
        layanan_id, 
        user_id, 
        jumlah, 
        tanggal_order, 
        total, 
        status
    ) VALUES (
        $1, $2, $3, $4, $5, CURRENT_TIMESTAMP, $6, $7
    ) RETURNING *;
`

	var newOrder entity.Order
	err := tx.QueryRowxContext(ctx, sqlQuery,
		orderReq.Nama_pelanggan,
		orderReq.No_Telepon_Pelanggan,
		orderReq.Layanan_id,
		orderReq.User_id,
		orderReq.Jumlah,
		orderReq.Total,
		orderReq.Status,
	).StructScan(&newOrder)

	if err != nil {
		return entity.Order{}, err
	}

	return newOrder, nil
}

func (o *OrderRepository) DeleteOrderById(ctx context.Context, id int, tx *sqlx.Tx) error {
	sqlQuery := `DELETE FROM orders WHERE id = $1;`

	_, err := tx.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}
