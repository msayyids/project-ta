package controller

import (
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

type OrderController struct {
	S service.OrderServiceInj
}

type OrderControllerInj interface {
	CreateOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params)
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
			Data:   "invalid input",
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
