package routes

import (
	"chat-api/controller"

	"github.com/go-chi/chi/v5"
)

func roomRoutes(r *chi.Mux) {
	r.Route("/rooms", func(r chi.Router) {
		r.Post("/", controller.PostRoom)
		r.Get("/{id}", controller.GetRoom)
	})
}
