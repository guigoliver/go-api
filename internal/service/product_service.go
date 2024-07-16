package service

import (
	"github.com/devfullcycle/imersao-full-cycle/goapi/internal/database"
	"github.com/devfullcycle/imersao-full-cycle/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

// Creates a new instance of the ProductService struct
func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: productDB,
	}
}

// Retrieves all products from the ProductDB
func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Retrieves a product from the ProductDB by its ID
func (ps *ProductService) GetProductByID(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Retrieves all products from the ProductDB that belong to a specific category
func (ps *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Creates a new product in the ProductDB
func (ps *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, category_id, image_url, price)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
