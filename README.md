# GO E-COMMERCE API

## Overview
This repository contains a microservice which manages products and categories for a ficticious application called Full Cycle Ecommerce. This service is an API written in Golang responsible for creating and fetching products and categories for the Ecommerce platform.

## Table of Contents
- [Overview](#overview)
- [Getting Started](#getting-started)
- [Endpoints](#endpoints)
  - [Category Endpoints](#category-endpoints)
  - [Product Endpoints](#product-endpoints)

## Getting Started

### Prerequisites
- Docker
- Golang
- MySQL

### Setup Instructions
1. Clone the repository:
    ```sh
    git clone https://github.com/guigoliver/go-api.git
    cd go-api
    ```

2. Start the database using Docker Compose:
    ```sh
    docker-compose up
    ```

3. Open another terminal and start the application:

    ```sh
    go run cmd/catalog/main.go
    ```
The application will be available on port 8080.

## Endpoints

### Category Endpoints

#### GET /categories
- **Description:** Retrieves data related to all the existing categories.
- **HTTP Method:** GET
- **Parameters:** None
- **Response examples:**

  - `200 OK`: Returns a list of categories
    ```json
      [
        {
          "id": "6b291e52-9aa1-43d4-a80f-aa005718e5c6",
          "name": "Electronics"
        },
        {
          "id": "/b7bc8664-c962-4bf2-9ba7-510cf1790dfe",
          "name": "Books"
        }
      ]
    ```

  - `500 Internal Server Error`: If an error occurs.


#### GET /categories/{id}
- **Description:** Retrieves data related to a specific category.
- **HTTP Method:** GET
- **Path Parameter**:
  - `id` (string, required) - The ID of the category.
- **Response examples:**

  - `200 OK`: Returns the requested category
    ```json
        {
          "id": "6b291e52-9aa1-43d4-a80f-aa005718e5c6",
          "name": "Electronics"
        }
    ```

  - `400 Bad Request`: If the `id` parameter is missing.

  - `500 Internal Server Error`: If an error occurs.


#### POST /categories
- **Description:** Create a new category.
- **HTTP Method**: POST
- **Request Body:** JSON object representing the category.
  - Example: 
    ```json
    {
      "name": "Furniture"
    }
- **Response examples:**

  - `201 Created` Returns the created category object
    ```json
        {
          "id": "1702b3ab-fcfd-4e58-9d64-fab69daf86ee",
          "name": "Furniture"
        }
    ```

  - `400 Bad Request`: If the request body is invalid.

  - `500 Internal Server Error`: If an error occurs.


### Product Endpoints

#### GET /products  
- **Description:** Retrieves data related to all product.
- **HTTP Method:** GET
- **Parameters:** None  
- **Response examples:**  

  - `200 OK`: Returns a list of products.
    ```json
      [
        {
          "id": "55e69544-00a4-49ca-a38f-62fd9b764492",
          "name": "Smart TV 52''",
          "description": "A large screen HD TV",
          "categoryID": "0a07b2f2-384f-4684-9749-a91ecde66c2c",
          "imageURL": "http://example.com/smarttv.jpg",
          "price": 5400
        },
        {
          "id": "d0717ec2-1b1b-4ba7-b5e0-8c82aaf99b18",
          "name": "Tablet 64GB",
          "description": "A high-end tablet",
          "categoryID": "38687117-ecca-440c-8f9b-0b73489d5a2f",
          "imageURL": "http://example.com/tablet.jpg",
          "price": 3200
        }
      ]
    ```

  - `500 Internal Server Error`: If an error occurs.

#### GET /products/{id}
- **Description:** Retrieves data related to a specific product.
- **HTTP Method:** GET
- **Path Parameter:**
  -  `id` (string, required) - The ID of the product.  
- **Response examples:**
  
  -`200 OK`: Returns the requested product.
    ```json
        {
          "id": "55e69544-00a4-49ca-a38f-62fd9b764492",
          "name": "Smart TV 52''",
          "description": "A large screen HD TV",
          "categoryID": "0a07b2f2-384f-4684-9749-a91ecde66c2c",
          "imageURL": "http://example.com/smarttv.jpg",
          "price": 5400
        }
    ```

  -`400 Bad Request`: If the `id` parameter is missing.

  -`500 Internal Server Error`: If an error occurs.

#### GET /products/category/{categoryID}
- **Description:** Retrieve products by category ID.
- **HTTP Method:** GET
- **Path Parameter:**
  - `categoryID` (string, required) - The ID of the category.  
- **Response examples:**

  - `200 OK`: Returns a list of products under the same category.
    ```json
    [
        {
          "id": "55e69544-00a4-49ca-a38f-62fd9b764492",
          "name": "Smart TV 52''",
          "description": "A large screen HD TV",
          "categoryID": "0a07b2f2-384f-4684-9749-a91ecde66c2c",
          "imageURL": "http://example.com/smarttv.jpg",
          "price": 5400
        },
         {
          "id": "091fbc1f-b56b-4d86-842e-e3e4e40a111a",
          "name": "HD TV 70''",
          "description": "A huge screen HD TV",
          "categoryID": "0a07b2f2-384f-4684-9749-a91ecde66c2c",
          "imageURL": "http://example.com/smarttv.jpg",
          "price": 5400
        }
    ]
    ```

  -`400 Bad Request`: If the `categoryID` parameter is missing.

  -`500 Internal Server Error`: If an error occurs.

#### POST /products
- **Description:** Create a new product.
- **HTTP Method:** POST
- **Request Body:** JSON object representing the product.
  - Example:
    ```json
    {
      "name": "Tablet 128GB",
      "description": "A new generation tablet",
      "categoryID": "38687117-ecca-440c-8f9b-0b73489d5a2f",
      "imageURL": "http://example.com/tablet2.jpg",
      "price": 500
    }
    ```  
- **Response examples:**

  - `201 Created`: Returns the created product object
    ```json
        {
          "id": "4fbed1c4-5374-49f6-97e3-1b61c67733df",
          "name": "Tablet 128GB",
          "description": "A new generation tablet",
          "categoryID": "38687117-ecca-440c-8f9b-0b73489d5a2f",
          "imageURL": "http://example.com/tablet2.jpg",
          "price": 500
        }
    ```

  - `400 Bad Request`: If the request body is invalid.

  - `500 Internal Server Error`: If an error occurs.
  
    
