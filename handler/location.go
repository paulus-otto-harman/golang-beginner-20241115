package handler

import (
	"20241115/lib"
	"20241115/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type LocationHandler struct {
	LocationService service.LocationService
}

func InitLocationHandler(locationService service.LocationService) *LocationHandler {
	return &LocationHandler{LocationService: locationService}
}

func (handler *LocationHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Destination ID")
		return
	}

	location, _ := handler.LocationService.Get(id)
	lib.JsonResponse(w).Success(http.StatusOK, "Location Found", location)
}
