package repository

import (
	"context"
	"log"
	"project-ta/entity"

	"gorm.io/gorm"
)

type OrderRepositoryInj interface {
	AddOrder(ctx context.Context, order entity.OrderReq, db *gorm.DB) (entity.Order, error)
}

type OrderRepository struct {
}

func NewOrderRepository() OrderRepositoryInj {
	return &OrderRepository{}
}

func (r *OrderRepository) AddOrder(ctx context.Context, order entity.OrderReq, db *gorm.DB) (entity.Order, error) {
	orders := entity.Order{
		NamaPelanggan:      order.NamaPelanggan,
		NoTeleponPelanggan: order.NoTeleponPelanggan,
		LayananID:          order.LayananID,
		UserID:             order.UserID,
		Jumlah:             order.Jumlah,
		TanggalOrder:       order.TanggalOrder,
		Total:              order.Total,
		Status:             order.Status,
		PaymentType:        order.PaymentType,
	}

	err := db.Create(&orders).Error
	if err != nil {
		log.Println("Error creating order:", err)
		return entity.Order{}, err
	}

	return orders, nil
}
