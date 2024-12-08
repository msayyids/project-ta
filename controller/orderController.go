package controller

import (
	"fmt"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type OrderController struct {
	Service service.OrderServiceInj
}

func NewOrderController(service service.OrderServiceInj) *OrderController {
	return &OrderController{
		Service: service,
	}
}

func (oc *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var order entity.OrderReq

	helper.RequestBody(r, &order)
	validate := validator.New()
	err := validate.Struct(order)
	if err != nil {
		// Tampilkan detail error untuk debugging
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    fmt.Sprintf("Invalid input: %v", err),
		}, http.StatusBadRequest)
		return
	}

	newOrder, err := oc.Service.CreateOrder(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := entity.WebResponse{
		Code:    http.StatusCreated,
		Message: "success login",
		Data:    newOrder,
	}

	helper.ResponseBody(w, response, http.StatusCreated)
}
