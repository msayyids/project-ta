package service

import (
	"project-ta/repository"

	"github.com/jmoiron/sqlx"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentServiceInj interface {
	// CreatePaymenUrl(ctx context.Context, id int, amount int) (entity.CreatePaymentResponse, error)
	// CheckPaymentStatus(ctx context.Context, orderID int) error
	// CheckOrderStatus(ctx context.Context, orderID int) error
}

type PaymentService struct {
	Client            snap.Client
	PaymentRepository repository.MidtransPaymentRepositoryInj
	OrderRepository   repository.OrderRepositoryInj
	DB                sqlx.DB
}

func NewPaymentService(client snap.Client, db sqlx.DB, pm repository.MidtransPaymentRepositoryInj, op repository.OrderRepositoryInj) PaymentServiceInj {
	return &PaymentService{
		Client:            client,
		PaymentRepository: pm,
		OrderRepository:   op,
		DB:                db,
	}
}

// func (ps PaymentService) CreatePaymenUrl(ctx context.Context, id int, amount int) (entity.CreatePaymentResponse, error) {

// 	intId := strconv.Itoa(id)
// 	req := &snap.Request{
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID:  intId,
// 			GrossAmt: int64(amount),
// 		},
// 		CreditCard: &snap.CreditCardDetails{
// 			Secure: true,
// 		},
// 	}

// 	// 3. Request create Snap transaction to Midtrans
// 	snapResp, err := ps.Client.CreateTransaction(req)
// 	if err != nil {
// 		return entity.CreatePaymentResponse{}, err
// 	}

// 	response := entity.CreatePaymentResponse{
// 		Snap_url: snapResp.RedirectURL,
// 		Orderid:  id,
// 	}

// 	return response, nil
// }
