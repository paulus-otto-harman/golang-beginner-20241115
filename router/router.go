package router

import (
	"20241115/handler"
	"20241115/middleware"
	"20241115/repository"
	"20241115/service"
	"database/sql"
	"github.com/go-chi/chi/v5"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	handleEvent := handler.InitEventHandler(*service.InitEventService(*repository.InitEventRepo(db)))
	handleOrder := handler.InitOrderHandler(*service.InitOrderService(*repository.InitOrderRepo(db)))

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.JsonResponse())
		r.Route("/events", func(r chi.Router) {
			r.Get("/", handleEvent.All)
			r.Get("/{id}", handleEvent.Get)
		})
		r.Post("/bookings", handleOrder.Create)

	})

	return r
}
