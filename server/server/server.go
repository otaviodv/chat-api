package server

import (
	"chat-api/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func setupCors(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func buildHandler() *chi.Mux {
	handler := chi.NewMux()
	chi.NewRouter()
	setupCors(handler)
	routes.SetupRoutes(handler)

	return handler
}

func Serve() {
	handler := buildHandler()
	port := 8080
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
		if err != nil {
			panic(err)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	fmt.Printf("Listening on port %d\n", port)
	<-quit
}
