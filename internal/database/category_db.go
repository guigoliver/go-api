package database

import (
	"database/sql"

	"github.com/devfullcycle/imersao-full-cycle/goapi/internal/entity"
)

// CategoryDB represents a database for categories.
type CategoryDB struct {
	db *sql.DB
}

// NewCategoryDB creates a new instance of CategoryDB.
func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db}
}

// GetCategories retrieves all categories from the database.
func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

// GetCategory retrieves a category from the database by its ID.
func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateCategory creates a new category in the database.
func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		return "", err
	}
	return category.ID, nil
}
