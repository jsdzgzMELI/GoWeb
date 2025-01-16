package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/domain"
	"github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/service"
	"github.com/jsdzgzMELI/GoWeb/GoWebTotal/pkg/web/response"
)

type ProductHandler struct {
	service service.ProductsServ
}

func IniProductHandler(serv service.ProductsServ) ProductHandler {
	return ProductHandler{service: serv}
}

func (ph *ProductHandler) UpdateProductHttp(w http.ResponseWriter, r *http.Request) {
	apiToken := r.Header.Get("Authorization")
	if apiToken != os.Getenv("API_TOKEN") {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Token not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
		return
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid id"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}
	var request domain.Product
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &response.Response{Message: "error decoding request", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	// pr := &domain.Product{
	// 	ID:           id,
	// 	Name:         request.Name,
	// 	Quantity:     request.Quantity,
	// 	Code_value:   request.Code_value,
	// 	Is_published: request.Is_published,
	// 	Expiration:   request.Expiration,
	// 	Price:        request.Price,
	// }
	err = ph.service.UpdateProduct(id, request)
	// fmt.Println(service.Products)
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	pr, _ := ph.service.GetById(id)
	code := http.StatusOK
	body := &response.Response{Message: "Product updated", Data: &pr}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (ph *ProductHandler) PatchProductHttp(w http.ResponseWriter, r *http.Request) {
	apiToken := r.Header.Get("Authorization")
	if apiToken != os.Getenv("API_TOKEN") {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Token not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
		return
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "error decoding request"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	var request domain.Product

	// var request pkg.RequestPatch
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "error decoding request"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	err = ph.service.PatchProduct(id, request)
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	product, err := ph.service.GetById(id)
	code := http.StatusOK
	body := &response.Response{Message: "Product patched", Data: &product}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
	return
}

func (ph *ProductHandler) DeleteProductHttp(w http.ResponseWriter, r *http.Request) {
	apiToken := r.Header.Get("Authorization")
	if apiToken != os.Getenv("API_TOKEN") {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Token not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
		return
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid id"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	err = ph.service.DeleteProduct(id)
	// fmt.Println(service.Products)
	if err != nil {
		code := http.StatusNotFound
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Product not found"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &response.Response{Message: "Product deleted", Data: nil}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (ph *ProductHandler) GetProductsHttp(w http.ResponseWriter, r *http.Request) {
	products, err := ph.service.GetAllProducts()
	if err != nil {
		code := http.StatusExpectationFailed
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "No products found"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &response.ResponseGet{Message: "Products", Data: &products}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (ph *ProductHandler) AddProductHttp(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(service.Products)
	apiToken := r.Header.Get("Authorization")
	if apiToken != os.Getenv("API_TOKEN") {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Token not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
		return
	}
	var request domain.Product
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Error decoding request"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	err := ph.service.AddProduct(request)

	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	code := http.StatusCreated
	body := &response.Response{Message: "Product added", Data: &request}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (ph *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("get by id"))
	// id := r.URL.Query().Get("id")
	id := chi.URLParam(r, "id")
	if id == "" || id == "0" {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "id is required and can't be 0"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid id parsing"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	product, err := ph.service.GetById(intID)
	if err != nil {
		code := http.StatusNotFound
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Product not found"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK

	body := &response.Response{Message: "Product found", Data: &product}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}
