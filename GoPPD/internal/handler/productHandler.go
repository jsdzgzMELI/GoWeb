package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jsdzgzMELI/GoWeb/GoPPD/internal"
	"github.com/jsdzgzMELI/GoWeb/GoPPD/pkg"
)

func UpdateProductHttp(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusBadRequest
		body := &pkg.Response{Message: "Invalid id", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}
	var request pkg.RequestUpdate
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &pkg.Response{Message: "error decoding request", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	pr := &internal.Product{
		ID:           id,
		Name:         request.Name,
		Quantity:     request.Quantity,
		Code_value:   request.Code_value,
		Is_published: request.Is_published,
		Expiration:   request.Expiration,
		Price:        request.Price,
	}
	err = UpdateProduct(id, pr)
	fmt.Println(internal.Products)
	if err != nil {
		code := http.StatusBadRequest
		body := &pkg.Response{Message: err.Error(), Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &pkg.Response{Message: "Product updated", Data: pr}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func UpdateProduct(id int, p *internal.Product) error {
	pr := &internal.Product{
		ID:           p.ID,
		Name:         p.Name,
		Quantity:     p.Quantity,
		Code_value:   p.Code_value,
		Is_published: p.Is_published,
		Expiration:   p.Expiration,
		Price:        p.Price,
	}

	err := ValueCheck(*p)
	if err != nil {
		return err

	}
	err = isValidDateFormat((*p).Expiration)
	if err != nil {
		return errors.New("Invalid date format")
	}
	index, _, err := FindById(id)
	if err != nil {
		internal.Products = append(internal.Products, *p)
		return err
	}
	internal.Products[index] = *pr
	return nil
}

func PatchProductHttp(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusBadRequest
		body := &pkg.Response{Message: "Invalid id", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	_, pr, err := FindById(id)
	if err != nil {
		code := http.StatusNotFound
		body := &pkg.Response{Message: "Product not found", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	request := &pkg.RequestPatch{
		Name:         pr.Name,
		Quantity:     pr.Quantity,
		Code_value:   pr.Code_value,
		Is_published: pr.Is_published,
		Expiration:   pr.Expiration,
		Price:        pr.Price,
	}

	// var request pkg.RequestPatch
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &pkg.Response{Message: "error decoding request", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	pr = &internal.Product{
		ID:           id,
		Name:         request.Name,
		Quantity:     request.Quantity,
		Code_value:   request.Code_value,
		Is_published: request.Is_published,
		Expiration:   request.Expiration,
		Price:        request.Price,
	}

	err = PatchProduct(id, pr)
	if err != nil {
		code := http.StatusBadRequest
		body := &pkg.Response{Message: err.Error(), Data: nil}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &pkg.Response{Message: "Product patched", Data: pr}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
	return
}

func PatchProduct(id int, p *internal.Product) error {
	index, _, err := FindById(id)
	if err != nil {
		return err
	}
	pr := &internal.Product{
		ID:           id,
		Name:         p.Name,
		Quantity:     p.Quantity,
		Code_value:   p.Code_value,
		Is_published: p.Is_published,
		Expiration:   p.Expiration,
		Price:        p.Price,
	}
	internal.Products[index] = *pr
	// err = ValueCheck(*p)
	// if err != nil {
	// 	return err
	// }
	// err = isValidDateFormat((*p).Expiration)
	// if err != nil {
	// 	return errors.New("Invalid date format")
	// }
	return nil
}

func DeleteProductHttp(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusBadRequest
		body := &pkg.Response{Message: "Invalid id", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	err = DeleteProduct(id)
	fmt.Println(internal.Products)
	if err != nil {
		code := http.StatusNotFound
		body := &pkg.Response{Message: "Product not found", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &pkg.Response{Message: "Product deleted", Data: nil}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
	return
}
func DeleteProduct(id int) error {
	index, _, err := FindById(id)
	if err != nil {
		return err
	}
	internal.Products = append(internal.Products[:index], internal.Products[index+1:]...)
	return nil
}

func FindById(id int) (int, *internal.Product, error) {
	for index, product := range internal.Products {
		if product.ID == id {
			return index, &product, nil
		}
	}
	return 0, nil, errors.New("Product not found")
}

func ExistingProduct(p internal.Product) error {
	for _, product := range internal.Products {
		if product.ID == p.ID {
			return nil
		}
	}
	return errors.New("Product doesn't exist")
}

func AddNotExistingProduct(p internal.Product) error {
	pr := &internal.Product{p.ID, p.Name, p.Quantity, p.Code_value, p.Is_published, p.Expiration, p.Price}

	err := ExistingProduct(*pr)

	if err != nil {
		internal.Products = append(internal.Products, *pr)
	}
	internal.Products[pr.ID] = *pr
	return nil
}

func CodeValueUnique(p internal.Product) error {
	for _, product := range internal.Products {
		if product.Code_value == p.Code_value {
			return errors.New("Code_value is not unique")
		}
	}
	return nil
}

func ValueCheck(p internal.Product) error {
	if &internal.Products == nil {
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
