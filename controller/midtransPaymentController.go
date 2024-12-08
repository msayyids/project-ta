package controller

import (
	"net/http"
	midtrans "project-ta/Mmidtrans"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go/snap"
)

type WebhookPayload struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
}

type PaymentController struct {
	S      service.PaymentServiceInj
	Os     service.OrderServiceInj
	Client snap.Client
}

type PaymentControllerInj interface {
	CreatePayment(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewPaymentController(s service.PaymentServiceInj, os service.OrderServiceInj, c snap.Client) PaymentControllerInj {
	return PaymentController{
		S:      s,
		Os:     os,
		Client: c,
	}
}

func (P PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	// Define a variable to hold the request data
	var orderReq struct {
		OrderID int `json:"order_id"`
		Amount  int `json:"amount"`
	}

	helper.RequestBody(r, &orderReq)
	url := midtrans.GeneratePaymentUrl(orderReq.OrderID, orderReq.Amount, P.Client)

	// Return the response with the generated payment URL
	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "Payment URL generated successfully",
		Data:    url,
	}, http.StatusOK)

}
