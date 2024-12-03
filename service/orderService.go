package service

import (
	"context"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"

	"github.com/jmoiron/sqlx"
)

type OrderServiceInj interface {
	CreateOrder(ctx context.Context, orderReq entity.OrderRequest) (entity.Order, error)
}

type OrderServices struct {
	DB        *sqlx.DB
	OrderRepo repository.OrderRepositoryInj
}

func NewOrderService(or repository.OrderRepositoryInj, db *sqlx.DB) OrderServiceInj {
	return OrderServices{
		DB:        db,
		OrderRepo: or}
}

func (s OrderServices) CreateOrder(ctx context.Context, orderReq entity.OrderRequest) (entity.Order, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	newUsers, err := s.OrderRepo.AddOrder(ctx, orderReq, *tx)
	helper.PanicIfError(err)

	return newUsers, nil

}
