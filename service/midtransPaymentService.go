package service

import (
	"context"
	"fmt"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentServiceInj interface {
	CreatePayment(ctx context.Context, id int) (entity.MidtransPayment, error)
}

type paymentService struct {
	Client            snap.Client
	PaymentRepository repository.MidtransPaymentRepositoryInj
	OrderRepository   repository.OrderRepositoryInj
	DB                sqlx.DB
}

func NewPaymentService(client snap.Client, db sqlx.DB, pm repository.MidtransPaymentRepositoryInj, op repository.OrderRepositoryInj) PaymentServiceInj {
	return &paymentService{
		Client:            client,
		PaymentRepository: pm,
		OrderRepository:   op,
		DB:                db,
	}
}

func (ps paymentService) CreatePayment(ctx context.Context, id int) (entity.MidtransPayment, error) {
	tx, err := ps.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := ps.OrderRepository.FindOrderById(ctx, id, tx)
	if err != nil {
		// Jika order tidak ditemukan
		return entity.MidtransPayment{}, fmt.Errorf("error retrieving order: %v", err)
	}

	// Generate transaksi dengan Midtrans Snap
	orderId := strconv.Itoa(order.Id)
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(order.Total),
		},
	}

	snapResponse, _ := ps.Client.CreateTransaction(request)

	newPayment := entity.MidtransPayment{
		OrderID:     order.Id,
		RedirectURL: snapResponse.RedirectURL,
		SubTotal:    order.Total,
		PaymentDate: time.Now(),
		Status:      "pending",
		CreatedAt:   time.Now(),
	}

	// Simpan ke database
	savedPayment, err := ps.PaymentRepository.AddPayment(ctx, newPayment, *tx)
	if err != nil {
		return entity.MidtransPayment{}, fmt.Errorf("error 3")
	}

	return savedPayment, nil
}
