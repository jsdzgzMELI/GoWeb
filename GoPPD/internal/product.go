package internal

import (
	"errors"
	"time"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

type ProductIdless struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

var Products []Product

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
