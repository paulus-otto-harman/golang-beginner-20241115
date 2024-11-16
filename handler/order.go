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
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid input", nil)
		return
	}
	err := lib.Validate(booking)
	if err != nil {
		lib.JsonResponse(w).ValidationFail(err)
		return
	}

	err = handler.OrderService.Create(&booking)
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Booking Failed", nil)
		return
	}
	lib.JsonResponse(w).Success(http.StatusCreated, "Tour Booked", nil)
}
