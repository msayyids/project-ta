package repository

import (
	"context"
	"project-ta/entity"

	"gorm.io/gorm"
)

type OrderRepositoryInj interface {
	AddOrder(ctx context.Context, order entity.OrderReq, db *gorm.DB) (entity.Order, error)
	FindById(ctx context.Context, id int, db *gorm.DB) (entity.Order, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]entity.Order, error)
	UpdateOrder(ctx context.Context, id int, order entity.OrderReq, db *gorm.DB) (entity.Order, error)
	DeleteOrder(ctx context.Context, id int, db *gorm.DB) error
	SaveOrder(ctx context.Context, order entity.Order, tx *gorm.DB) error
	UpdatePaymentURL(ctx context.Context, orderID int, paymentURL string, db *gorm.DB) error
	UpdateStatus(ctx context.Context, orderId int, status string, db *gorm.DB) (*entity.Order, error)
	FindByStatus(ctx context.Context, status string, db *gorm.DB) ([]entity.Order, error)
}

type OrderRepository struct{}

func NewOrderRepository() OrderRepositoryInj {
	return &OrderRepository{}
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, orderId int, status string, db *gorm.DB) (*entity.Order, error) {
	var order entity.Order

	// Update kolom status dan ambil data yang diperbarui
	if err := db.WithContext(ctx).Model(&order).
		Where("id = ?", orderId).
		Update("status", status).Error; err != nil {
		return nil, err
	}

	// Ambil data order yang sudah diperbarui
	if err := db.WithContext(ctx).First(&order, orderId).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) UpdatePaymentURL(ctx context.Context, orderID int, paymentURL string, db *gorm.DB) error {
	// Find the existing order by ID
	var existingOrder entity.Order
	if err := db.WithContext(ctx).First(&existingOrder, orderID).Error; err != nil {
		return err
	}

	// Update only the payment_url field
	existingOrder.Payment_url = paymentURL

	// Save the updated order with the new payment_url
	if err := db.WithContext(ctx).Save(&existingOrder).Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) SaveOrder(ctx context.Context, order entity.Order, tx *gorm.DB) error {
	if err := tx.Save(&order).Error; err != nil {
		return err
	}
	return nil
}

// AddOrder to insert new order into database
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
		return entity.Order{}, err
	}

	return orders, nil
}

// FindById to get order by ID
func (r *OrderRepository) FindById(ctx context.Context, id int, db *gorm.DB) (entity.Order, error) {
	var order entity.Order
	if err := db.WithContext(ctx).First(&order, id).Error; err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

// FindAll to get all orders
func (r *OrderRepository) FindAll(ctx context.Context, db *gorm.DB) ([]entity.Order, error) {
	var orders []entity.Order
	if err := db.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrder to update existing order by ID
func (r *OrderRepository) UpdateOrder(ctx context.Context, id int, order entity.OrderReq, db *gorm.DB) (entity.Order, error) {
	var existingOrder entity.Order
	if err := db.WithContext(ctx).First(&existingOrder, id).Error; err != nil {
		return entity.Order{}, err
	}

	// Update the fields
	existingOrder.NamaPelanggan = order.NamaPelanggan
	existingOrder.NoTeleponPelanggan = order.NoTeleponPelanggan
	existingOrder.LayananID = order.LayananID
	existingOrder.UserID = order.UserID
	existingOrder.Jumlah = order.Jumlah
	existingOrder.Total = order.Total
	existingOrder.Status = order.Status
	existingOrder.PaymentType = order.PaymentType

	if err := db.WithContext(ctx).Save(&existingOrder).Error; err != nil {
		return entity.Order{}, err
	}

	return existingOrder, nil
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, id int, db *gorm.DB) error {
	if err := db.WithContext(ctx).Delete(&entity.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindByStatus(ctx context.Context, status string, db *gorm.DB) ([]entity.Order, error) {
	var orders []entity.Order
	if err := db.WithContext(ctx).Where("status=?", status).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil

}
