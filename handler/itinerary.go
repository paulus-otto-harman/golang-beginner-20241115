package handler

import (
	"20241115/service"
	"net/http"
)

type ItineraryHandler struct {
	ItineraryService service.ItineraryService
}

func InitItineraryHandler(itineraryService service.ItineraryService) *ItineraryHandler {
	return &ItineraryHandler{ItineraryService: itineraryService}
}

func (handler *ItineraryHandler) Get(w http.ResponseWriter, r *http.Request) {

}
