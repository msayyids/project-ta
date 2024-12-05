package controller

import (
	"fmt"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

type PaymentController struct {
	S  service.PaymentServiceInj
	Os service.OrderServiceInj
}

type PaymentControllerInj interface {
	CreatePayment(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewPaymentController(s service.PaymentServiceInj, os service.OrderServiceInj) PaymentControllerInj {
	return PaymentController{
		S:  s,
		Os: os,
	}
}

func (p PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var reqId entity.ReqId
	helper.RequestBody(r, &reqId)

	// Pastikan OrderID ada di database
	payment, err := p.S.CreatePayment(r.Context(), reqId.OrderID)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   fmt.Errorf("error processing payment: %v", err),
		})
		return
	}

	// Berikan respon sukses
	response := entity.WebResponse{
		Code:   201,
		Status: "SUCCESS CREATE PAYMENT",
		Data:   payment,
	}

	helper.ResponseBody(w, response)
}
