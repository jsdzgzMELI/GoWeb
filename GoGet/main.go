package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jsdzgzMELI/GoWebGet/handlers"
)

func main() {
	// http.HandleFunc("/products", handlers.ProductHandler)
	// http.HandleFunc("/products/search", handlers.PriceHigherThan)
	// http.HandleFunc("/id", handlers.GetById)
	r := chi.NewRouter()
	r.Get("/ping", handlers.PingHandler)
	// r.Get("/products/{id}", handlers.GetById)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handlers.ProductHandler)
		r.Get("/id/{id}", handlers.GetById)
		r.Get("/search", handlers.PriceHigherThan)
	})
	http.ListenAndServe(":8080", r)
}
