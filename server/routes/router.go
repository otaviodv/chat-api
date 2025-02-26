package routes

import (
	"chat-api/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer)
	roomRoutes(r)
	messageRoutes(r)
	return r
}

func SetupWS() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/subscribe/room/{id}", controller.HandleSubscribe)
	return r
}
