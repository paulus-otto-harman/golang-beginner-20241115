package handler

import (
	"20241115/lib"
	"20241115/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ItineraryHandler struct {
	ItineraryService service.ItineraryService
}

func InitItineraryHandler(itineraryService service.ItineraryService) *ItineraryHandler {
	return &ItineraryHandler{ItineraryService: itineraryService}
}

func (handler *ItineraryHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Destination ID")
		return
	}

	itinerary, _ := handler.ItineraryService.Get(id)
	lib.JsonResponse(w).Success(http.StatusOK, "Itinerary Found", itinerary)
}
