package service

import (
	"context"
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
	CreatePayment(ctx context.Context, payment entity.MidtransPaymentRequest) (entity.MidtransPayment, error)
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

func (ps paymentService) CreatePayment(ctx context.Context, payment entity.MidtransPaymentRequest) (entity.MidtransPayment, error) {

	tx, err := ps.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := ps.OrderRepository.FindOrderById(ctx, payment.OrderID, tx)
	if err != nil {
		return entity.MidtransPayment{}, err
	}

	payments, err := ps.PaymentRepository.AddPayment(ctx, payment, *tx)
	if err != nil {
		return entity.MidtransPayment{}, err
	}

	payments.SubTotal = order.Total

	orderId := strconv.Itoa(payment.OrderID)
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(payments.SubTotal),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: order.Nama_pelanggan,
			Phone: order.No_Telepon_Pelanggan,
		},
	}

	snapResponse, _ := ps.Client.CreateTransaction(request)

	response := entity.MidtransPayment{
		OrderID:     payment.OrderID,
		RedirectURL: snapResponse.RedirectURL,
		SubTotal:    payments.SubTotal,
		PaymentDate: time.Now(),
		Status:      "pending",
		CreatedAt:   time.Now(),
	}

	return response, nil

}
