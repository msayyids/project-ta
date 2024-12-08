package service

import (
	"context"
	"project-ta/entity"
	"project-ta/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OrderServiceInj interface {
	CreateOrder(ctx context.Context, order entity.OrderReq) (entity.Order, error)
}

type OrderService struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Repo     repository.OrderRepositoryInj
}

func NewOrderService(db *gorm.DB, v *validator.Validate, repo repository.OrderRepositoryInj) OrderServiceInj {
	return &OrderService{
		DB:       db,
		Validate: v,
		Repo:     repo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order entity.OrderReq) (entity.Order, error) {
	// Validasi input
	err := s.Validate.Struct(order)
	if err != nil {
		return entity.Order{}, err
	}

	tx := s.DB.Begin()

	// Pastikan transaksi di-rollback jika terjadi error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// Buat order baru
	newOrder, err := s.Repo.AddOrder(ctx, order, tx)
	if err != nil {
		tx.Rollback() // Rollback jika ada error
		return entity.Order{}, err
	}

	// Commit transaksi setelah order berhasil ditambahkan
	if err := tx.Commit().Error; err != nil {
		return entity.Order{}, err
	}

	return newOrder, nil
}
