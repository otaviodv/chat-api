package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(r *chi.Mux) *chi.Mux {
	r.Use(middleware.RequestID, middleware.Recoverer)
	roomRoutes(r)
	return r
}
