package entity

import "github.com/google/uuid"

// Category represents a category of products.
type Category struct {
	// ID is the unique identifier of the category.
	ID string
	// Name is the name of the category.
	Name string
}

// NewCategory creates a new Category with a randomly generated ID and the given name.
// Parameters:
// - name: the name of the category.
// Returns:
// - *Category: a pointer to the newly created Category.
func NewCategory(name string) *Category {
	return &Category{
		ID: uuid.New().String(),
		Name: name,
	}
}

// Product represents a product.
type Product struct {
	// ID is the unique identifier of the product.
	ID string
	// Name is the name of the product.
	Name string
	// Description is a brief description of the product.
	Description string
	// Price is the sale price of the product.
	Price float64
	// CategoryID is the ID of the category to which the product belongs.
	CategoryID string
	// ImageURL is the URL of the product's display image.
	ImageURL string
}

// NewProduct creates a new Product with a randomly generated ID and the given attributes.
// Parameters:
// - name: the name of the product.
// - description: a brief description of the product.
// - categoryID: the ID of the category to which the product belongs.
// - imageURL: the URL of the product's display image.
// - price: the sale price of the product.
// Returns:
// - *Product: a pointer to the newly created Product.
func NewProduct(name, description string, categoryID, imageURL string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
	}
}
