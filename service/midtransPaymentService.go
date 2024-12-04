package service

import (
	"context"
	"log"
	"project-ta/entity"
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentServiceInj interface {
	CreatePayment(ctx context.Context, payment entity.MidtransPaymentRequest) (entity.MidtransPayment, error)
}

type paymentService struct {
	Client snap.Client
	// PaymentRepository repository.MidtransPaymentRepositoryInj
	// OrderRepository   repository.OrderRepositoryInj
	// DB                sqlx.DB
}

func NewPaymentService(client snap.Client) PaymentServiceInj {
	return &paymentService{
		Client: client,
		// PaymentRepository: paymentRepo,
		// OrderRepository:   orderRepo,
		// DB:                db,
	}
}

func (ps paymentService) CreatePayment(ctx context.Context, payment entity.MidtransPaymentRequest) (entity.MidtransPayment, error) {

	// tx, err := ps.DB.Beginx()
	// if err != nil {
	// 	return entity.MidtransPayment{}, err
	// }
	// defer helper.CommitOrRollback(tx)

	// payments, err := ps.PaymentRepository.AddPayment(ctx, payment, *tx)
	// if err != nil {
	// 	return entity.MidtransPayment{}, err
	// }

	orderId := strconv.Itoa(payment.Order_id)
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(payment.Amount),
		},
		// CustomerDetail: &midtrans.CustomerDetails{
		// 	FName: "john",
		// 	Phone: "081213131212",
		// },
	}

	snapResponse, err := ps.Client.CreateTransaction(request)
	if err != nil {
		log.Printf("Failed to create Midtrans transaction: %v", err)
		return entity.MidtransPayment{}, err.GetRawError()
	}
	response := entity.MidtransPayment{
		Id:          1,
		Order_id:    payment.Order_id,
		Amount:      payment.Amount,
		RedirectUrl: snapResponse.RedirectURL,
		PaymentTime: time.Now(),
		Status:      "done",
	}

	return response, nil

}
