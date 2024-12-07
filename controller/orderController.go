package controller

import (
	"fmt"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/middleware"
	"project-ta/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type OrderController struct {
	Client coreapi.Client
	S      service.OrderServiceInj
}

type OrderControllerInj interface {
	CreateOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetAllOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetOrderById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	EditOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	DeleteORder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewOrderController(s service.OrderServiceInj, c coreapi.Client) OrderControllerInj {
	return OrderController{S: s, Client: c}
}

func (oc OrderController) CreateOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	userCtx, ok := r.Context().Value(middleware.UserKey).(middleware.UserContext)
	if !ok {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "UNAUTHORIZED",
			Data:    nil,
		}, http.StatusUnauthorized)
		return
	}

	newOrder := entity.OrderRequest{}
	newOrder.User_id = userCtx.ID
	helper.RequestBody(r, &newOrder)

	newOrderResponse, err := oc.S.CreateOrder(r.Context(), newOrder)
	if err != nil {
		fmt.Printf("Layanan found: %+v\n", newOrderResponse.Layanan_id)
		fmt.Printf("OrderRequest after calculation: %+v\n", newOrderResponse.Total)
		fmt.Printf("OrderRequest after calculation: %+v\n", newOrder)
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    "error disini",
		}, http.StatusBadRequest)
		return
	}

	strId := strconv.Itoa(newOrderResponse.Id)

	requestPayment := coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeQris,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strId,
			GrossAmt: int64(newOrder.Total),
		},
	}

	snapResponse, _ := oc.Client.ChargeTransaction(&requestPayment)
	newOrderResponse.Payment_url = snapResponse.Actions[0].URL
	newOrderResponse.Payment_type = snapResponse.PaymentType

	response := entity.WebResponse{
		Code:    http.StatusCreated,
		Message: "CREATED",
		Data:    newOrderResponse,
	}

	fmt.Printf("Layanan found: %+v\n", newOrderResponse.Layanan_id)
	fmt.Printf("OrderRequest after calculation: %+v\n", newOrderResponse.Total)

	helper.ResponseBody(w, response, http.StatusCreated)
}

func (oc OrderController) GetAllOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	allOrder, err := oc.S.GetOrder(r.Context())
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
		Message: "OK",
		Data:    allOrder,
	}

	helper.ResponseBody(w, response, http.StatusOK)

}

func (oc OrderController) EditOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	var orderReq entity.OrderRequest
	helper.RequestBody(r, &orderReq)

	editedOrder, err := oc.S.EditOrderById(r.Context(), idInt, orderReq)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    "invalid input",
		}, http.StatusBadRequest)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "success edit order",
		Data:    editedOrder,
	}

	helper.ResponseBody(w, response, http.StatusOK)

}

func (oc OrderController) GetOrderById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	editedOrder, err := oc.S.GetOrderById(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    "invalid input",
		}, http.StatusBadRequest)
		return
	}

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    editedOrder,
	}

	helper.ResponseBody(w, response, http.StatusOK)

}

func (oc OrderController) DeleteORder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	err := oc.S.DeleteOrder(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "not found",
			Data:    "invalid input",
		}, http.StatusNotFound)
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    fmt.Sprintf("User with ID %s has been deleted", id),
	}, http.StatusOK)

}
