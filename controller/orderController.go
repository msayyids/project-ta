package controller

import (
	"github.com/midtrans/midtrans-go/snap"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go"
)

type OrderController struct {
	Service service.OrderServiceInj
	S       snap.Client
	V       validator.Validate
}

func NewOrderController(service service.OrderServiceInj, s snap.Client, v validator.Validate) *OrderController {
	return &OrderController{
		Service: service,
		S:       s,
		V:       v,
	}
}

func (oc *OrderController) CreateOrderCash(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var order entity.OrderReq

	helper.RequestBody(r, &order)

	newOrder, err := oc.Service.CreateOrder(r.Context(), order)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		}, http.StatusInternalServerError)
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
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		}, http.StatusInternalServerError)
		return
	}

	strId := strconv.Itoa(newOrder.ID)

	chargeReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strId,
			GrossAmt: int64(newOrder.Total),
		}, CustomerDetail: &midtrans.CustomerDetails{
			FName: newOrder.NamaPelanggan,
			Phone: newOrder.NoTeleponPelanggan,
		},
	}

	snapRes, _ := oc.S.CreateTransaction(chargeReq)

	err = oc.Service.UpdatePaymentURL(r.Context(), newOrder.ID, snapRes.RedirectURL)
	if err != nil {
		http.Error(w, "Failed to update payment URL and QR data", http.StatusInternalServerError)
		return
	}

	newOrder.Payment_url = snapRes.RedirectURL
	response := entity.WebResponse{
		Code:    http.StatusCreated,
		Message: "Order Created Successfully",
		Data:    newOrder,
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
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		}, http.StatusInternalServerError)
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
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		}, http.StatusInternalServerError)
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
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		}, http.StatusInternalServerError)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "Order Updated Successfully",
		Data:    updatedOrder,
	}

	helper.ResponseBody(w, response, http.StatusOK)

}

func (oc *OrderController) FindByStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	status := ps.ByName("status")

	order, err := oc.Service.FindByStatus(r.Context(), status)
	if err != nil {
		helper.ResponseBody(w,
			entity.WebResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			}, http.StatusInternalServerError)
	}

	respose := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    order,
	}

	helper.ResponseBody(w, respose, http.StatusOK)

}
