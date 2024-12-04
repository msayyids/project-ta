package service

import (
	"context"
	"fmt"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"

	"github.com/jmoiron/sqlx"
)

type OrderServiceInj interface {
	CreateOrder(ctx context.Context, orderReq entity.OrderRequest) (entity.Order, error)
	EditOrderById(ctx context.Context, id int, orderReq entity.OrderRequest) (entity.Order, error)
	GetOrderById(ctx context.Context, id int) (entity.Order, error)
	GetOrder(ctx context.Context) ([]entity.Order, error)
	DeleteOrder(ctx context.Context, id int) error
}

type OrderServices struct {
	DB          sqlx.DB
	OrderRepo   repository.OrderRepositoryInj
	LayananRepo repository.LayananRepositoryInj
}

func NewOrderService(or repository.OrderRepositoryInj, db sqlx.DB, lp repository.LayananRepositoryInj) OrderServiceInj {
	return OrderServices{
		DB:          db,
		OrderRepo:   or,
		LayananRepo: lp,
	}
}

func (s OrderServices) CreateOrder(ctx context.Context, orderReq entity.OrderRequest) (entity.Order, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// Ambil data layanan berdasarkan ID
	layanan, err := s.LayananRepo.FindById(ctx, orderReq.Layanan_id, *tx)
	if err != nil {
		return entity.Order{}, fmt.Errorf("error dsni")
	}

	// Hitung total
	orderReq.Total = layanan.Harga * orderReq.Jumlah

	// Simpan order ke database
	newOrder, err := s.OrderRepo.AddOrder(ctx, orderReq, tx)
	if err != nil {
		return entity.Order{}, fmt.Errorf("error 1 not found")
	}

	return newOrder, nil
}

func (s OrderServices) EditOrderById(ctx context.Context, id int, orderReq entity.OrderRequest) (entity.Order, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	editedOrder, err := s.OrderRepo.UpdateOrderById(ctx, id, orderReq, tx)

	helper.PanicIfError(err)

	return editedOrder, nil
}

func (s OrderServices) GetOrderById(ctx context.Context, id int) (entity.Order, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	order, err := s.OrderRepo.FindOrderById(ctx, id, tx)
	helper.PanicIfError(err)

	return order, nil
}

func (s OrderServices) GetOrder(ctx context.Context) ([]entity.Order, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	allOrder, err := s.OrderRepo.FindOrder(ctx, tx)
	helper.PanicIfError(err)

	return allOrder, nil
}

func (s OrderServices) DeleteOrder(ctx context.Context, id int) error {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	if err := s.DeleteOrder(ctx, id); err != nil {
		return err
	}

	return nil
}
