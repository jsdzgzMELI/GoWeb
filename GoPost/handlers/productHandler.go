package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jsdzgzMELI/GoWeb/GoPost/structs"
)

func GetProductsHttp(w http.ResponseWriter, r *http.Request) {
	if len(structs.Products) == 0 {
		code := http.StatusExpectationFailed
		body := &structs.ResponseGet{Message: "No products found", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	code := http.StatusOK
	body := &structs.ResponseGet{Message: "Last product", Data: &structs.Products[len(structs.Products)-1]}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func AddProductHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println(structs.Products)
	if structs.Products == nil {
		structs.Products = []structs.Product{}
	}
	// p := structs.Product{ID: 1, Name: "Product 1", Quantity: 100, Code_value: "123", Is_published: true, Expiration: "2022-12-31", Price: 100.0}
	// err := json.NewDecoder(r.Body).Decode(&p)
	// err := AddProduct(p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	var request structs.Product
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &structs.ResponsePost{Message: "error decoding request", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	fmt.Println(request)

	pr := &structs.Product{
		ID:           len(structs.Products) + 1,
		Name:         request.Name,
		Quantity:     request.Quantity,
		Code_value:   request.Code_value,
		Is_published: request.Is_published,
		Expiration:   request.Expiration,
		Price:        request.Price,
	}

	err := AddProduct(*pr)

	if err != nil {
		code := http.StatusBadRequest
		body := &structs.ResponsePost{Message: err.Error(), Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	fmt.Println(structs.Products)

	code := http.StatusCreated
	body := &structs.ResponsePost{Message: "Product added", Data: pr}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func AddProduct(p structs.Product) error {
	err := CodeValueUnique(p)
	if err != nil {
		return err
	}
	err = ValueCheck(p)
	if err != nil {
		return err

	}
	err = isValidDateFormat(p.Expiration)
	if err != nil {
		return errors.New("Invalid date format")
	}
	structs.Products = append(structs.Products, p)
	return nil
}

func CodeValueUnique(p structs.Product) error {
	for _, product := range structs.Products {
		if product.Code_value == p.Code_value {
			return errors.New("Code_value is not unique")
		}
	}
	return nil
}

func ValueCheck(p structs.Product) error {
	if structs.Products == nil {
		return errors.New("Products is nil")
	}
	if (p.ID == 0) || (p.Name == "") || (p.Quantity == 0) || (p.Code_value == "") || (p.Expiration == "") || (p.Price == 0.0) {
		return errors.New("Product is missing values")
	}
	return nil
}

func isValidDateFormat(date string) error {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return err
	}
	return nil
}
