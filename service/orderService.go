package service

import (
	"context"
	"fmt"
	"log"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"

	"github.com/jmoiron/sqlx"
	"github.com/midtrans/midtrans-go/coreapi"
)

type OrderServiceInj interface {
	CreateOrder(ctx context.Context, orderReq entity.OrderRequest) (entity.Order, error)
	EditOrderById(ctx context.Context, id int, orderReq entity.OrderRequest) (entity.Order, error)
	GetOrderById(ctx context.Context, id int) (entity.Order, error)
	GetOrder(ctx context.Context) ([]entity.Order, error)
	DeleteOrder(ctx context.Context, id int) error
}

type OrderServices struct {
	Client      coreapi.Client
	DB          sqlx.DB
	OrderRepo   repository.OrderRepositoryInj
	LayananRepo repository.LayananRepositoryInj
}

func NewOrderService(client coreapi.Client, or repository.OrderRepositoryInj, db sqlx.DB, lp repository.LayananRepositoryInj) OrderServiceInj {
	return OrderServices{
		Client:      client,
		DB:          db,
		OrderRepo:   or,
		LayananRepo: lp,
	}
}

func (s OrderServices) CreateOrder(ctx context.Context, orderReq entity.OrderRequest) (entity.Order, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// Validate input
	if orderReq.Layanan_id <= 0 || orderReq.Jumlah <= 0 {
		log.Printf("Invalid input: %+v", orderReq)
		return entity.Order{}, fmt.Errorf("invalid input")
	}

	layanan, err := s.LayananRepo.FindById(ctx, orderReq.Layanan_id, *tx)
	if err != nil {
		log.Printf("Failed to find layanan: %v", err)
		return entity.Order{}, fmt.Errorf("error dsni")
	}

	log.Printf("Harga layanan ditemukan: %d", layanan.Harga)
	orderReq.Total = layanan.Harga * orderReq.Jumlah

	newOrder, err := s.OrderRepo.AddOrder(ctx, orderReq, tx)
	if err != nil {
		log.Printf("Error adding order: %v", err)
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
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Order{}, fmt.Errorf("error 1")
	}

	defer helper.CommitOrRollback(tx)

	order, err := s.OrderRepo.FindOrderById(ctx, id, tx)
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Order{}, fmt.Errorf("error 2")
	}

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

	err = s.DeleteOrder(ctx, id)
	helper.PanicIfError(err)

	return nil
}
