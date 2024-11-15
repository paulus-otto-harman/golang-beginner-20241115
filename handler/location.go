package handler

import (
	"20241115/service"
	"net/http"
)

type LocationHandler struct {
	LocationService service.LocationService
}

func InitLocationHandler(locationService service.LocationService) *LocationHandler {
	return &LocationHandler{LocationService: locationService}
}

func (handler *LocationHandler) Get(w http.ResponseWriter, r *http.Request) {

}
