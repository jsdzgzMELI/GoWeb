package service

import (
	"errors"

	"github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/domain"
	"github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/repository"
)

type ProductsServ interface {
	GetAllProducts() ([]domain.Product, error)
	GetById(int) (domain.Product, error)
	PatchProduct(int, domain.Product) error
	AddProduct(domain.Product) error
	UpdateProduct(int, domain.Product) error
	DeleteProduct(int) error
}

type productService struct {
	repo repository.ProductRepo
}

func IniProductServ(repo repository.ProductRepo) ProductsServ {
	return &productService{
		repo: repo,
	}
}
func (ps *productService) GetById(id int) (domain.Product, error) {
	product, err := ps.repo.GetById(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (ps *productService) GetAllProducts() ([]domain.Product, error) {
	products, err := ps.repo.GetAllProducts()
	if err != nil {
		return nil, errors.New("no products")
	}
	return products, nil
}

func (ps *productService) UpdateProduct(id int, p domain.Product) error {
	return ps.repo.UpdateProduct(id, p)
}

func (ps *productService) PatchProduct(id int, p domain.Product) error {
	return ps.repo.PatchProduct(id, p)
}

func (ps *productService) DeleteProduct(id int) error {
	return ps.repo.DeleteProduct(id)
}

func (ps *productService) AddProduct(p domain.Product) error {
	return ps.repo.AddProduct(p)
}
