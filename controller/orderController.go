package controller

import (
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type OrderController struct {
	Service service.OrderServiceInj
	C       coreapi.Client
	V       validator.Validate
}

func NewOrderController(service service.OrderServiceInj, c coreapi.Client, v validator.Validate) *OrderController {
	return &OrderController{
		Service: service,
		C:       c,
		V:       v,
	}
}

func (oc *OrderController) CreateOrderCash(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var order entity.OrderReq

	// if err := oc.V.Struct(order); err != nil {
	// 	helper.ResponseBody(w, entity.WebResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "BAD REQUEST",
	// 		Data:    "INVALID INPUT",
	// 	}, http.StatusBadRequest)
	// 	return
	// }

	helper.RequestBody(r, &order)

	newOrder, err := oc.Service.CreateOrder(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusCreated,
		Message: "Order Created Successfully",
		Data:    newOrder,
	}

	helper.ResponseBody(w, response, http.StatusCreated)
}

func (oc *OrderController) CreateOrderCashless(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var order entity.OrderReq

	helper.RequestBody(r, &order)

	newOrder, err := oc.Service.CreateOrder(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	strId := strconv.Itoa(newOrder.ID)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeQris,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strId,
			GrossAmt: int64(newOrder.Total),
		},
	}

	coreApiRes, _ := oc.C.ChargeTransaction(chargeReq)

	err = oc.Service.UpdatePaymentURL(r.Context(), newOrder.ID, coreApiRes.Actions[0].URL)
	if err != nil {
		http.Error(w, "Failed to update payment URL and QR data", http.StatusInternalServerError)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusCreated,
		Message: "Order Created Successfully",
		Data: map[string]interface{}{
			"order":          newOrder,
			"payment_link":   coreApiRes.Actions[0].URL,
			"transaction_id": coreApiRes.TransactionID,
			"qr_string":      coreApiRes.QRString,
		},
	}

	helper.ResponseBody(w, response, http.StatusCreated)
}

func (oc *OrderController) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	order, err := oc.Service.FindById(r.Context(), id)
	if err != nil {
		http.Error(w, "Order Not Found", http.StatusNotFound)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "Order Found",
		Data:    order,
	}

	helper.ResponseBody(w, response, http.StatusOK)
}

func (oc *OrderController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	orders, err := oc.Service.FindAll(r.Context())
	if err != nil {
		http.Error(w, "Unable to retrieve orders", http.StatusInternalServerError)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "Orders Found",
		Data:    orders,
	}

	helper.ResponseBody(w, response, http.StatusOK)
}

func (oc *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var order entity.OrderReq
	helper.RequestBody(r, &order)

	updatedOrder, err := oc.Service.UpdateOrder(r.Context(), id, order)
	if err != nil {
		http.Error(w, "Unable to update order", http.StatusInternalServerError)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "Order Updated Successfully",
		Data:    updatedOrder,
	}

	helper.ResponseBody(w, response, http.StatusOK)

}
