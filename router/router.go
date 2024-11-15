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
	handleLocation := handler.InitLocationHandler(*service.InitLocationService(*repository.InitLocationRepo(db)))
	handleItinerary := handler.InitItineraryHandler(*service.InitItineraryService(*repository.InitItineraryRepo(db)))

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.JsonResponse())
		r.Route("/events", func(r chi.Router) {
			r.Get("/", handleEvent.All)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handleEvent.Get)

				r.Get("/plan", handleItinerary.Get)
				r.Get("/location", handleLocation.Get)
			})

		})
		r.Post("/bookings", handleOrder.Create)

	})

	return r
}
