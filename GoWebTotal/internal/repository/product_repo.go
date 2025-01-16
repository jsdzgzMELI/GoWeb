package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/domain"
)

type ProductRepo interface {
	GetAllProducts() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	AddProduct(product domain.Product) error
	UpdateProduct(id int, product domain.Product) error
	PatchProduct(id int, product domain.Product) error
	DeleteProduct(id int) error
}

type productRepository struct {
	products []domain.Product
}

func IniProductRepo(filename string) (ProductRepo, error) {
	// PORQUE ES UNA DIRECCION DE MEMORIA?
	repo := &productRepository{}
	err := repo.loadProducts(filename)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (pr *productRepository) loadProducts(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// var products []internal.Product
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pr.products)
	if err != nil {
		return err
	}

	return nil
}

func (pr *productRepository) GetAllProducts() ([]domain.Product, error) {
	if pr.products == nil {
		return nil, errors.New("no products")
	}
	return pr.products, nil
}

func (pr *productRepository) GetById(id int) (domain.Product, error) {
	for _, product := range pr.products {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

func (pr *productRepository) AddProduct(product domain.Product) error {
	product.ID = pr.getNextID()
	pr.products = append(pr.products, product)
	return nil
}

func (pr *productRepository) getNextID() int {
	maxID := 0
	for _, product := range pr.products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}

func (pr *productRepository) UpdateProduct(id int, product domain.Product) error {
	for i, p := range pr.products {
		if p.ID == id {
			product.ID = id
			pr.products[i] = product
			return nil
		}
	}
	return fmt.Errorf("product not found")
}

func (pr *productRepository) PatchProduct(id int, product domain.Product) error {

	idx, _, err := pr.FindById(id)

	if err != nil {
		return err
	}
	if product.Name != "" {
		pr.products[idx].Name = product.Name
	}
	if product.Quantity >= 0 {
		pr.products[idx].Quantity = product.Quantity
	}
	if product.Code_value != "" {
		pr.products[idx].Code_value = product.Code_value
	}
	if product.Is_published {
		pr.products[idx].Is_published = product.Is_published
	}
	if product.Expiration != "" {
		pr.products[idx].Expiration = product.Expiration
	}
	if product.Price >= 0 {
		pr.products[idx].Price = product.Price
	}
	return nil
}

func (pr *productRepository) DeleteProduct(id int) error {
	for i, p := range pr.products {
		if p.ID == id {
			pr.products = append(pr.products[:i], pr.products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Product not found")
}

func (pr *productRepository) FindById(id int) (int, *domain.Product, error) {
	for index, product := range pr.products {
		if product.ID == id {
			return index, &product, nil
		}
	}
	return 0, nil, errors.New("Product not found")
}

func (pr *productRepository) CodeValueUnique(p domain.Product) error {
	for _, product := range pr.products {
		if product.Code_value == p.Code_value {
			return errors.New("Code_value is not unique")
		}
	}
	return nil
}

func (pr *productRepository) ValueCheck(p domain.Product) error {
	if &pr.products == nil {
		return errors.New("Products is nil")
	}
	if (p.ID == 0) || (p.Name == "") || (p.Quantity == 0) || (p.Code_value == "") || (p.Expiration == "") || (p.Price == 0.0) {
		return errors.New("Product is missing values")
	}
	return nil
}

func (pr *productRepository) isValidDateFormat(date string) error {
	currDate, err := time.Parse("02/01/2006", date)
	if err != nil {
		return err
	}
	if currDate.Before(time.Now()) {
		return errors.New("Expiration date should be in the future")
	}
	return nil
}
