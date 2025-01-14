package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jsdzgzMELI/GoWebGet/structs"

	"github.com/jsdzgzMELI/GoWebGet/helper"
)

const (
	// ProductPath is the path to the products file
	ProductPath = "./products.json"
)

type Controller struct {
	st map[string]string
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	products, err := helper.LoadProducts("./products.json")
	if err != nil {
		log.Fatalf("Error loading products: %v", err)
	}
	json.NewEncoder(w).Encode(products)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("get by id"))
	// id := r.URL.Query().Get("id")
	id := chi.URLParam(r, "id")
	if id == "" || id == "0" {
		code := http.StatusBadRequest
		body := &structs.ResponseId{Message: "id is required and can't be 0", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	idInt, intErr := strconv.ParseInt(id, 10, 64)
	fmt.Println(idInt)
	if intErr != nil {
		log.Fatalf("Error converting id: %v to int: %v", id, intErr)
	}
	products, err := helper.LoadProducts("./products.json")
	if err != nil {
		log.Fatalf("Error loading products: %v", err)
	}
	idx := slices.IndexFunc(products, func(p structs.Product) bool {
		return p.ID == int(idInt)
	})
	product := &structs.Products[idx]
	// json.NewEncoder(w).Encode(*product)
	// fmt.Println(*product)
	code := http.StatusOK
	// AQUI TOCA ITERAR LOS PRODUCTOS Y BUSCAR EL PRODUCTO POR ID
	body := &structs.ResponseId{Message: "product found", Data: &structs.Product{
		ID:           product.ID,
		Name:         product.Name,
		Quantity:     product.Quantity,
		Code_value:   product.Code_value,
		Is_published: product.Is_published,
		Expiration:   product.Expiration,
		Price:        product.Price,
	}}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func PriceHigherThan(w http.ResponseWriter, r *http.Request) {
	price := r.URL.Query().Get("price")
	if price == "" {
		code := http.StatusBadRequest
		body := &structs.ResponseId{Message: "price is required", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	priceInt, intErr := strconv.ParseInt(price, 10, 64)
	if intErr != nil {
		log.Fatalf("Error converting price: %v to int: %v", price, intErr)
	}
	products, err := helper.LoadProducts("./products.json")
	if err != nil {
		log.Fatalf("Error loading products: %v", err)
	}
	var productsHigherThan []structs.Product
	for _, product := range products {
		if product.Price > float64(priceInt) {
			productsHigherThan = append(productsHigherThan, product)
		}
	}
	code := http.StatusOK
	body := &structs.ResponsePrice{Message: "products found", Data: &productsHigherThan, Total: len(productsHigherThan)}
	// body := map[string]string{"message": "products found", "data": fmt.Sprintf("%v", productsHigherThan)}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}
