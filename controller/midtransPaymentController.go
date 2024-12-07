package controller

import (
	"encoding/json"
	"net/http"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go/coreapi"
)

type WebhookPayload struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
}

type PaymentController struct {
	S      service.PaymentServiceInj
	Os     service.OrderServiceInj
	Client coreapi.Client
}

type PaymentControllerInj interface {
	VerifyPayment(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewPaymentController(s service.PaymentServiceInj, os service.OrderServiceInj, c coreapi.Client) PaymentControllerInj {
	return PaymentController{

		S:      s,
		Os:     os,
		Client: c,
	}
}

func (p PaymentController) VerifyPayment(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	var notificationPayload map[string]interface{}

	// 2. Parse JSON request body and use it to set json to payload
	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		// do something on error when decode
		return
	}
	// 3. Get order-id from payload
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return
	}

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := p.Client.CheckTransaction(orderId)
	if e != nil {
		http.Error(w, e.GetMessage(), http.StatusInternalServerError)
		return
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
				// buat notif dan bukti pembayaran di table payment
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ok"))
}
