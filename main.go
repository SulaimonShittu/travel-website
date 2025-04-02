package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-website/api/handlers"
	"travel-website/api/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/contacts", handlers.ContactRoutes)
		r.Route("/destinations", handlers.DestinationRoutes)
	})

	storage.SeedDestinations()

	fmt.Println("Starting server on :3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
