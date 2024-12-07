package service

import (
	"context"
	"project-ta/repository"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentServiceInj interface {
	// CreatePayment(ctx context.Context, id int) (entity.MidtransPayment, error)
	// CheckPaymentStatus(ctx context.Context, orderID int) error
	// CheckOrderStatus(ctx context.Context, orderID int) error
}

type PaymentService struct {
	Client            coreapi.Client
	PaymentRepository repository.MidtransPaymentRepositoryInj
	OrderRepository   repository.OrderRepositoryInj
	DB                sqlx.DB
}

func NewPaymentService(client coreapi.Client, db sqlx.DB, pm repository.MidtransPaymentRepositoryInj, op repository.OrderRepositoryInj) PaymentServiceInj {
	return &PaymentService{
		Client:            client,
		PaymentRepository: pm,
		OrderRepository:   op,
		DB:                db,
	}
}

// func (ps PaymentService) CreatePayment(ctx context.Context, id int) (entity.MidtransPayment, error) {
// 	tx, err := ps.DB.Beginx()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)

// 	order, err := ps.OrderRepository.FindOrderById(ctx, id, tx)
// 	if err != nil {

// 		return entity.MidtransPayment{}, fmt.Errorf("error retrieving order: %v", err)
// 	}

// 	// orderId := strconv.Itoa(order.Id)
// 	// request := coreapi.ChargeReq{
// 	// 	PaymentType: coreapi.PaymentTypeQris,
// 	// 	TransactionDetails: midtrans.TransactionDetails{
// 	// 		OrderID:  orderId,
// 	// 		GrossAmt: int64(order.Total),
// 	// 	},
// 	// }

// 	// snapResponse, _ := ps.Client.ChargeTransaction(&request)

// 	// newPayment := entity.MidtransPayment{
// 	// 	OrderID:     order.Id,
// 	// 	RedirectURL: snapResponse.Actions[0].URL,
// 	// 	SubTotal:    order.Total,
// 	// 	PaymentDate: time.Now(),
// 	// 	Status:      snapResponse.TransactionStatus,
// 	// 	CreatedAt:   time.Now(),
// 	// }

// 	savedOrder, err := ps.OrderRepository.AddOrder(ctx,order)

// 	savedPayment, err := ps.PaymentRepository.AddPayment(ctx, newPayment, *tx)
// 	if err != nil {
// 		return entity.MidtransPayment{}, fmt.Errorf("error 3")
// 	}

// 	return savedPayment, nil
// }

func (ps PaymentService) GetPaymentStatus(ctx context.Context, id int) (string, error) {

	strId := strconv.Itoa(id)
	paymentStatus, _ := ps.Client.CheckTransaction(strId)
	status := paymentStatus.TransactionStatus

	return status, nil
}

// func (ps PaymentService) CheckPaymentStatus(ctx context.Context, orderID int) error {
// 	// Mulai transaksi
// 	tx, err := ps.DB.Beginx()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)

// 	// Konversi orderID ke string
// 	strId := strconv.Itoa(orderID)

// 	// Periksa status transaksi dari Midtrans
// 	transactionStatus, _ := ps.Client.CheckTransaction(strId)

// 	intid, _ := strconv.Atoi(transactionStatus.OrderID)

// 	// Tentukan status baru berdasarkan hasil transaksi
// 	var status string
// 	switch transactionStatus.TransactionStatus {
// 	case "settlement":
// 		status = "paid"
// 		_, err = ps.PaymentRepository.EditPaymentbyOrderId(ctx, intid, status, *tx)
// 		if err != nil {
// 			return fmt.Errorf("failed to update order status: %v", err)
// 		}

// 	case "pending":
// 		status = "Pending"
// 	case "cancel", "expire":
// 		status = "Failed"
// 	case "deny":
// 		status = "Denied"
// 	default:
// 		status = "Unknown"
// 	}

// 	return nil
// }

// func (ps PaymentService) CheckOrderStatus(ctx context.Context, orderID int) error {
// 	// Mulai transaksi
// 	tx, err := ps.DB.Beginx()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)

// 	// Konversi orderID ke string
// 	strId := strconv.Itoa(orderID)

// 	// Periksa status transaksi dari Midtrans
// 	transactionStatus, _ := ps.Client.CheckTransaction(strId)

// 	intid, _ := strconv.Atoi(transactionStatus.OrderID)

// 	// Tentukan status baru berdasarkan hasil transaksi
// 	var status string
// 	switch transactionStatus.TransactionStatus {
// 	case "settlement":
// 		status = "paid"
// 		_, err = ps.OrderRepository.EditStatusOrder(ctx, intid, status, tx)
// 		if err != nil {
// 			return fmt.Errorf("failed to update order status: %v", err)
// 		}

// 	case "pending":
// 		status = "Pending"
// 	case "cancel", "expire":
// 		status = "Failed"
// 	case "deny":
// 		status = "Denied"
// 	default:
// 		status = "Unknown"
// 	}

// 	return nil
// }
