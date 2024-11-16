package handler

import (
	"20241115/lib"
	"20241115/model"
	"20241115/service"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type EventHandler struct {
	EventService service.EventService
}

func InitEventHandler(eventService service.EventService) EventHandler {
	return EventHandler{EventService: eventService}
}

func (handler EventHandler) All(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	page := 1
	var err error
	if p := r.URL.Query().Get("page"); p != "" {
		page, err = strconv.Atoi(p)
	}

	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Page")
		return
	}

	sort := r.URL.Query().Get("sort")

	totalItems, totalPages, items, err := handler.EventService.All(date, page, sort)
	if err != nil {
		lib.JsonResponse(w).Fail(0, err.Error())
		return
	}

	json.NewEncoder(w).Encode(model.DataPage{
		Success:    true,
		Page:       page,
		Limit:      6,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       items,
	})
}

func (handler EventHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Destination ID")
		return
	}

	event := model.Event{Id: id}
	err = handler.EventService.Get(&event)
	if err != nil {
		log.Println(err)
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Failed to retrieve Event")
		return
	}
	lib.JsonResponse(w).Success(http.StatusOK, "Event Found", event)
}
