package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jsdzgzMELI/GoWeb/GoPost/handlers"
)

func main() {
	r := chi.NewRouter()

	r.Post("/products", handlers.AddProductHttp)
	r.Get("/products", handlers.GetProductsHttp)

	http.ListenAndServe(":8080", r)
}

/*
{
    "name": "Product 1",
    "quantity": 1,
    "code_value": "1",
    "is_published": false,
    "expiration": "01/01/2020",
    "price": 200
}
*/
