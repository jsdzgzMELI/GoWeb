package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	handler "github.com/jsdzgzMELI/GoWeb/GoPPD/internal/handler"
)

func main() {
	r := chi.NewRouter()

	r.Put("/products/{id}", handler.UpdateProductHttp)
	r.Patch("/products/{id}", handler.PatchProductHttp)
	r.Delete("/products/{id}", handler.DeleteProductHttp)

	http.ListenAndServe(":8080", r)
}
