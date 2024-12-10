package repository

import (
	"context"
	"project-ta/entity"

	"gorm.io/gorm"
)

type PaymentRepositoryInj interface {
	CreatePayment(ctx context.Context, payment *entity.Payment, db *gorm.DB) (*entity.Payment, error)
}

type PaymentRepository struct{}

func NewPaymentRepository() PaymentRepositoryInj {
	return &PaymentRepository{}
}

func (r *PaymentRepository) CreatePayment(ctx context.Context, payment *entity.Payment, db *gorm.DB) (*entity.Payment, error) {
	if err := db.WithContext(ctx).Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}
