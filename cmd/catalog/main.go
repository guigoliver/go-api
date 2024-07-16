package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/devfullcycle/imersao-full-cycle/goapi/internal/database"
	"github.com/devfullcycle/imersao-full-cycle/goapi/internal/service"
	"github.com/devfullcycle/imersao-full-cycle/goapi/internal/webserver"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	// Open a connection to the MySQL database.
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Create instances of the database and service layers.
	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	// Create instances of the web server.
	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	// Configure the server.
	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	// Define routes for categories and products.
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/categories", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/products", webProductHandler.GetProducts)
	c.Get("/product/category/{categoryId}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	// Start the server.
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
