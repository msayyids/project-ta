package controller

import (
	"fmt"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type OrderController struct {
	S service.OrderServiceInj
}

type OrderControllerInj interface {
	CreateOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetAllOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetOrderById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	EditOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	DeleteORder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewOrderController(s service.OrderServiceInj) OrderControllerInj {
	return OrderController{S: s}
}

func (oc OrderController) CreateOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	newOrder := entity.OrderRequest{}
	helper.RequestBody(r, &newOrder)

	newOrderResponse, err := oc.S.CreateOrder(r.Context(), newOrder)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   400,
			Status: "bad request",
			Data:   fmt.Errorf("layanan not found"),
		})
		return
	}

	response := entity.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data:   newOrderResponse,
	}

	helper.ResponseBody(w, response)
}

func (oc OrderController) GetAllOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	allOrder, err := oc.S.GetOrder(r.Context())
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   500,
			Status: "internal server error",
			Data:   nil,
		})
		return
	}

	response := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   allOrder,
	}

	helper.ResponseBody(w, response)

}

func (oc OrderController) EditOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	var orderReq entity.OrderRequest
	helper.RequestBody(r, &orderReq)

	editedOrder, err := oc.S.EditOrderById(r.Context(), idInt, orderReq)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   400,
			Status: "bad request",
			Data:   "invalid input",
		})
		return
	}

	response := entity.WebResponse{
		Code:   200,
		Status: "success edit order",
		Data:   editedOrder,
	}

	helper.ResponseBody(w, response)

}

func (oc OrderController) GetOrderById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	editedOrder, err := oc.S.GetOrderById(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   400,
			Status: "bad request",
			Data:   "invalid input",
		})
		return
	}

	response := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   editedOrder,
	}

	helper.ResponseBody(w, response)

}

func (oc OrderController) DeleteORder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	err := oc.S.DeleteOrder(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   401,
			Status: "not found",
			Data:   "invalid input",
		})
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   fmt.Sprintf("User with ID %s has been deleted", id),
	})

}
