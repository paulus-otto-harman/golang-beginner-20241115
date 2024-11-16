package handler

import (
	"20241115/lib"
	"20241115/model"
	"20241115/service"
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	OrderService service.OrderService
}

func InitOrderHandler(orderService service.OrderService) OrderHandler {
	return OrderHandler{OrderService: orderService}
}

func (handler OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var booking model.Order
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid input")
		return
	}

	if err := lib.Validate(booking); err != nil {
		lib.JsonResponse(w).ValidationFail(err)
		return
	}

	if err := handler.OrderService.Create(&booking); err != nil {
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Booking Failed")
		return
	}
	lib.JsonResponse(w).Success(http.StatusCreated, "Tour Booked", nil)
}
