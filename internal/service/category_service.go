package service

import (
	"github.com/devfullcycle/imersao-full-cycle/goapi/internal/database"
	"github.com/devfullcycle/imersao-full-cycle/goapi/internal/entity"
)

// CategoryService represents a service for handling categories.
type CategoryService struct {
	CategoryDB database.CategoryDB
}

// NewCategoryService creates a new instance of CategoryService.
func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

// GetCategories retrieves all categories from the database.
func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// CreateCategory creates a new category in the database.
func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := cs.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// GetCategory retrieves a category from the database by its ID.
func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
