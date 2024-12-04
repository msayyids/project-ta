package controller

import (
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

type PaymentController struct {
	S service.PaymentServiceInj
}

type PaymentControllerInj interface {
	CreatePayment(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewPaymentController(s service.PaymentServiceInj) PaymentControllerInj {
	return PaymentController{
		S: s,
	}
}

func (p PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var paymentRequest entity.MidtransPaymentRequest

	helper.RequestBody(r, &paymentRequest)

	payment, err := p.S.CreatePayment(r.Context(), paymentRequest)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err,
		})
		return
	}

	response := entity.WebResponse{
		Code:   201,
		Status: "SUCCESS CREATE PAYMENT",
		Data:   payment,
	}

	helper.ResponseBody(w, response)
}
