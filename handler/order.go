package handler

import (
	"20241115/lib"
	"20241115/model"
	"20241115/service"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
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
	err := validator.New().Struct(booking)

	if err != nil {
		errorBags := map[string]string{}
		if errors, ok := err.(validator.ValidationErrors); ok {
			lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "Validation Failed", errors)
			return
		}
		if _, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range err.(validator.ValidationErrors) {
				errorBags[validationErr.Field()] = fmt.Sprintf("Error: Field '%s' failed validation with tag '%s'.\n", validationErr.Field(), validationErr.Tag())
				fmt.Printf("Error: Field '%s' failed validation with tag '%s'.\n", validationErr.Field(), validationErr.Tag())
				fmt.Printf("  Value: '%v'\n", validationErr.Value())
				fmt.Printf("  Condition: '%s'\n", validationErr.Param())
			}
		} else {
			fmt.Println("Validation failed:", err)
		}
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "Validation Failed", errorBags)
		return
	}

	err = handler.OrderService.Create(&booking)
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Booking Failed", nil)
		return
	}
	lib.JsonResponse(w).Success(http.StatusCreated, "Tour Booked", nil)
}
