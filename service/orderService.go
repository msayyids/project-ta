package service

import (
	"context"
	"project-ta/entity"
	"project-ta/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OrderServiceInj interface {
	CreateOrder(ctx context.Context, order entity.OrderReq) (entity.Order, error)
	FindById(ctx context.Context, id int) (entity.Order, error)
	FindAll(ctx context.Context) ([]entity.Order, error)
	UpdateOrder(ctx context.Context, id int, order entity.OrderReq) (entity.Order, error)
	DeleteOrder(ctx context.Context, id int) error
	UpdatePaymentURL(ctx context.Context, orderID int, paymentURL string) error
	UpdateOrderStatus(ctx context.Context, orderId int, status string) (*entity.Order, error)
}

type OrderService struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Repo     repository.OrderRepositoryInj
	LR       repository.LayananRepositoryInj
}

func NewOrderService(db *gorm.DB, v *validator.Validate, repo repository.OrderRepositoryInj, lr repository.LayananRepositoryInj) OrderServiceInj {
	return &OrderService{
		DB:       db,
		Validate: v,
		Repo:     repo,
		LR:       lr,
	}
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderId int, status string) (*entity.Order, error) {
	// Panggil repository untuk update status
	updatedOrder, err := s.Repo.UpdateStatus(ctx, orderId, status, s.DB)
	if err != nil {
		return nil, err
	}

	return updatedOrder, nil
}

func (s *OrderService) UpdatePaymentURL(ctx context.Context, orderID int, paymentURL string) error {
	// Call the repository to update the payment URL in the database
	err := s.Repo.UpdatePaymentURL(ctx, orderID, paymentURL, s.DB)
	if err != nil {
		return err
	}

	return nil
}

// CreateOrder to add new order
func (s *OrderService) CreateOrder(ctx context.Context, order entity.OrderReq) (entity.Order, error) {

	// if err := s.Validate.Struct(order); err != nil {
	// 	return entity.Order{}, err
	// }

	// Mulai transaksi
	tx := s.DB.Begin()

	findLayanan, err := s.LR.FindById(ctx, order.LayananID, tx)
	if err != nil {
		tx.Rollback()
		return entity.Order{}, err
	}

	order.Total = findLayanan.Harga * order.Jumlah

	inputOrder := entity.OrderReq{
		NamaPelanggan:      order.NamaPelanggan,
		NoTeleponPelanggan: order.NoTeleponPelanggan,
		LayananID:          order.LayananID,
		UserID:             order.UserID,
		Jumlah:             order.Jumlah,
		TanggalOrder:       time.Now(),
		Total:              order.Total,
		Status:             order.Status,
		PaymentType:        order.PaymentType,
	}

	// Menambahkan order ke database
	newOrder, err := s.Repo.AddOrder(ctx, inputOrder, tx)
	if err != nil {
		tx.Rollback()
		return entity.Order{}, err
	}

	// Cari layanan berdasarkan LayananID

	// Menghitung total berdasarkan harga layanan dan jumlah order

	// Membuat objek order baru dengan nilai total yang dihitung

	// Simpan order baru ke database
	// if err := s.Repo.SaveOrder(ctx, newOrder, tx); err != nil {
	// 	tx.Rollback()
	// 	return entity.Order{}, err
	// }

	// Commit transaksi setelah semua perubahan berhasil
	if err := tx.Commit().Error; err != nil {
		return entity.Order{}, err
	}

	// Kembalikan order yang sudah disimpan
	return newOrder, nil
}

// FindById to get order by ID
func (s *OrderService) FindById(ctx context.Context, id int) (entity.Order, error) {
	return s.Repo.FindById(ctx, id, s.DB)
}

// FindAll to get all orders
func (s *OrderService) FindAll(ctx context.Context) ([]entity.Order, error) {
	return s.Repo.FindAll(ctx, s.DB)
}

// UpdateOrder to update an existing order
func (s *OrderService) UpdateOrder(ctx context.Context, id int, order entity.OrderReq) (entity.Order, error) {
	return s.Repo.UpdateOrder(ctx, id, order, s.DB)
}

// DeleteOrder to delete order by ID
func (s *OrderService) DeleteOrder(ctx context.Context, id int) error {
	return s.Repo.DeleteOrder(ctx, id, s.DB)
}
