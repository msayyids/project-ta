package repository

import (
	"context"
	"project-ta/entity"

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
        $1, $2, $3, $4, $5, CURRENT_TIMESTAMP, $6, $7
    ) RETURNING *;
    `

	var newOrder entity.Order
	err := tx.QueryRowxContext(ctx, sqlQuery,
		orderReq.Nama_pelanggaan,
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
	sqlQuery := `SELECT * FROM orders WHERE id = $1;`

	var order entity.Order
	err := tx.GetContext(ctx, &order, sqlQuery, id)
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (o *OrderRepository) UpdateOrderById(ctx context.Context, id int, orderReq entity.OrderRequest, tx *sqlx.Tx) (entity.Order, error) {
	sqlQuery := `
    UPDATE orders 
    SET 
        nama_pelanggan = $1, 
        no_telepon_pelanggan = $2, 
        layanan_id = $3, 
        user_id = $4, 
        jumlah = $5, 
        total = $6, 
        status = $7
    WHERE id = $8 
    RETURNING *;
    `

	var updatedOrder entity.Order
	err := tx.QueryRowxContext(ctx, sqlQuery,
		orderReq.Nama_pelanggaan,
		orderReq.No_Telepon_Pelanggan,
		orderReq.Layanan_id,
		orderReq.User_id,
		orderReq.Jumlah,
		orderReq.Total,
		orderReq.Status,
		id,
	).StructScan(&updatedOrder)

	if err != nil {
		return entity.Order{}, err
	}

	return updatedOrder, nil
}

func (o *OrderRepository) DeleteOrderById(ctx context.Context, id int, tx *sqlx.Tx) error {
	sqlQuery := `DELETE FROM orders WHERE id = $1;`

	_, err := tx.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}
