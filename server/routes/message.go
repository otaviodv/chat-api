package routes

import (
	"chat-api/controller"

	"github.com/go-chi/chi/v5"
)

func messageRoutes(r *chi.Mux) {
	r.Route("/rooms/{id}/messages", func(r chi.Router) {
		r.Post("/", controller.PostMessage)
	})
}
