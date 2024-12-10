package controller

import (
	"encoding/json"
	"net/http"
	"project-ta/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentController struct {
	Service service.PaymentServiceInj
	Os      service.OrderServiceInj
	C       coreapi.Client
}

func NewPaymentController(
	service service.PaymentServiceInj,
	os service.OrderServiceInj,
	c coreapi.Client,

) *PaymentController {
	return &PaymentController{
		Service: service,
		Os:      os,
		C:       c,
	}
}

func (pc *PaymentController) VerifyPayment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		// do something on error when decode
		return
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return
	}

	intId, _ := strconv.Atoi(orderId)
	_ = intId

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := pc.C.CheckTransaction(orderId)
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
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
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

// // FindById to get payment by ID
// func (pc *PaymentController) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	id, err := strconv.Atoi(ps.ByName("id"))
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	payment, err := pc.Service.FindById(r.Context(), id)
// 	if err != nil {
// 		http.Error(w, "Payment Not Found", http.StatusNotFound)
// 		return
// 	}

// 	response := entity.WebResponse{
// 		Code:    http.StatusOK,
// 		Message: "Payment Found",
// 		Data:    payment,
// 	}

// 	helper.ResponseBody(w, response, http.StatusOK)
// }

// // FindAll to get all payments
// func (pc *PaymentController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	payments, err := pc.Service.FindAll(r.Context())
// 	if err != nil {
// 		http.Error(w, "Unable to retrieve payments", http.StatusInternalServerError)
// 		return
// 	}

// 	response := entity.WebResponse{
// 		Code:    http.StatusOK,
// 		Message: "Payments Found",
// 		Data:    payments,
// 	}

// 	helper.ResponseBody(w, response, http.StatusOK)
// }
