package service

import (
	"context"

	"project-ta/entity"
	"project-ta/repository"

	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type PaymentServiceInj interface {
	CreatePayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
}

type PaymentService struct {
	DB          *gorm.DB
	Validate    *validator.Validate
	Repo        repository.PaymentRepositoryInj
	OrderRepo   repository.OrderRepositoryInj
	LayananRepo repository.LayananRepositoryInj
	CLient      snap.Client
}

func NewPaymentService(db *gorm.DB, v *validator.Validate, repo repository.PaymentRepositoryInj, or repository.OrderRepositoryInj, lr repository.LayananRepositoryInj) PaymentServiceInj {
	return &PaymentService{
		DB:          db,
		Validate:    v,
		Repo:        repo,
		OrderRepo:   or,
		LayananRepo: lr,
	}
}

func (s *PaymentService) CreatePayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {

	return s.Repo.CreatePayment(ctx, payment, s.DB)
}
