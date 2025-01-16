package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	handler "github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/handler"
	"github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/repository"
	"github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/service"
)

const fileName = "products.json"

func main() {

	iniRep, err := repository.IniProductRepo(fileName)
	if err != nil {
		panic(err)
	}
	iniService := service.IniProductServ(iniRep)
	iniHandler := handler.IniProductHandler(iniService)

	r := chi.NewRouter()

	if err != nil {
		panic(err)
	}

	r.Post("/products", iniHandler.AddProductHttp)
	r.Get("/products", iniHandler.GetProductsHttp)
	r.Get("/products/{id}", iniHandler.GetById)
	r.Put("/products/{id}", iniHandler.UpdateProductHttp)
	r.Patch("/products/{id}", iniHandler.PatchProductHttp)
	r.Delete("/products/{id}", iniHandler.DeleteProductHttp)

	http.ListenAndServe(":8080", r)
}
