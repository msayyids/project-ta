package repository

import (
	"context"
	"project-ta/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct{}

type OrderRepositoryInj interface {
	AddOrder(ctx context.Context, orderReq entity.OrderRequest, db sqlx.Tx) (entity.Order, error)
	FindOrderById(ctx context.Context, id int, db sqlx.Tx) (entity.Order, error)
	UpdateOrderById(ctx context.Context, id int, orderReq entity.OrderRequest, db sqlx.Tx) (entity.Order, error)
	DeleteOrderByid(ctx context.Context, id int, db sqlx.Tx) error
}

func NewOrderRepository() OrderRepositoryInj {
	return OrderRepository{}
}

func (o OrderRepository) AddOrder(ctx context.Context, orderReq entity.OrderRequest, tx sqlx.Tx) (entity.Order, error) {
	sqlQuery := `INSERT INTO orders (
    nama_pelanggan,
    no_telepon_pelanggan,
    layanan_id,
    user_id,
    jumlah,
    tanggal_order,
    total,
    status,
    tanggal_update
	) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP
	) RETURNING id;`

	var newOrder entity.Order
	err := tx.QueryRowxContext(ctx, sqlQuery,
		orderReq.Nama_pelanggaan,
		orderReq.No_Telepon_Pelanggan,
		orderReq.Layanan_id,
		orderReq.User_id,
		orderReq.Jumlah,
		time.Now(),
		orderReq.Total,
		orderReq.Status,
		time.Now(),
	).StructScan(&newOrder)

	if err != nil {
		return entity.Order{}, err
	}

	return newOrder, nil

}

func (o OrderRepository) FindOrderById(ctx context.Context, id int, db sqlx.Tx) (entity.Order, error) {
	return entity.Order{}, nil
}

func (o OrderRepository) UpdateOrderById(ctx context.Context, id int, orderReq entity.OrderRequest, db sqlx.Tx) (entity.Order, error) {

	return entity.Order{}, nil
}

func (o OrderRepository) DeleteOrderByid(ctx context.Context, id int, db sqlx.Tx) error {
	return nil
}
