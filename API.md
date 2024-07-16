# API Documentation
This page lists the available endpoints in this API with all the information relevant to their use.

## Table of Contents
- [Category Endpoints](#category-endpoints)
- [Product Endpoints](#product-endpoints)

## Category Endpoints

### List Categories

Retrieves data related to all the existing categories.

#### URL

`GET /categories`

#### Parameters
 
None.

#### Example request

```js
fetch(`http://localhost:8080/categories`, {
    method: 'get',
}).then(response => response.json()) 
.then(...);
```

#### Response

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| id | string | ID of the category, in the uuid format.
| name | string | Name of the category.

#### ***Example responses***

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

---

### Show Category

Retrieves data related to a specific category.

#### URL

`GET /categories/:id`

#### Path Parameter

| Parameter | Type | Description | Requirement |
| --------- | ---- | ----------- | ----------- |
| id | string | ID of the requested category | Required |


#### Example request

```js
fetch(`http://localhost:8080/categories/6b291e52-9aa1-43d4-a80f-aa005718e5c6`, {
    method: 'get',
}).then(response => response.json())
.then(...); 
```

#### Response

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| id | string | ID of the category, in the uuid format. |
| name | string | Name of the category. |


#### ***Response examples:***

  - `200 OK`: Returns the requested category
    ```json
        {
          "id": "6b291e52-9aa1-43d4-a80f-aa005718e5c6",
          "name": "Electronics"
        }
    ```

  - `400 Bad Request`: If the `id` parameter is missing.

  - `500 Internal Server Error`: If an error occurs.

---

### Create Category

Creates a new category.

#### URL

`POST /categories`

#### Request

| Parameter | Type | Description | Requirement |
| --------- | ---- | ----------- | ----------- |
| name | string | Name of the new category | Required |

#### ***Example Request***
```js
    let formData = new FormData();
    formData.append("name", "Furniture");

    fetch(`http://localhost:8080/categories`, {
        method: 'post',
        body: formData,
    }).then(response => response.json())
    .then(...);
```

#### Response

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| id | string | ID of the new category, in the uuid format. |
| name | string | Name of the category. |

#### ***Response examples:***

  - `201 Created` Returns the created category object
    ```json
        {
          "id": "1702b3ab-fcfd-4e58-9d64-fab69daf86ee",
          "name": "Furniture"
        }
    ```

  - `400 Bad Request`: If the request body is invalid.

  - `500 Internal Server Error`: If an error occurs.

---

## Product Endpoints

### List All Products

Retrieves data related to all products.

#### URL

`GET /products`

#### Parameters

None.

#### Example Request

```js
fetch(`http://localhost:8080/products`, {
    method: 'get',
}).then(response => response.json()) 
.then(...);
```
#### Response

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| id | string | ID of the product, in the uuid format.
| name | string | Name of the product.
| description | string | A brief description of the product. |
| categoryID | string | ID of the product's category. |
| imageURL | string | URL of the product's display image. |
| price | number | Sale price of the product. | 

#### ***Example responses***  

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

---

### Show Product

Retrieves data related to a specific product.

### URL

`GET /products/:id`

#### Path Parameter

| Parameter | Type | Description | Requirement |
| --------- | ---- | ----------- | ----------- |
| id | string | ID of the requested product | Required |

#### Example request

```js
fetch(`http://localhost:8080/products/55e69544-00a4-49ca-a38f-62fd9b764492`, {
    method: 'get',
}).then(response => response.json())
.then(...); 
```

#### Response

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| id | string | ID of the product, in the uuid format.
| name | string | Name of the product.
| description | string | A brief description of the product. |
| categoryID | string | ID of the product's category. |
| imageURL | string | URL of the product's display image. |
| price | number | Sale price of the product. | 

#### ***Response examples:***
  
  -`200 OK`: Returns the requested product

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

---

### Show Products by Category

Retrieves all products under a specific category.

#### URL

`GET /products/category/:categoryID`

#### Path Parameter

| Parameter | Type | Description | Requirement |
| --------- | ---- | ----------- | ----------- |
| categoryID | string | ID of the selected category | Required |

#### Example request

```js
fetch(`http://localhost:8080/products/category/0a07b2f2-384f-4684-9749-a91ecde66c2c`, {
    method: 'get',
}).then(response => response.json())
.then(...); 
```

#### Response

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| id | string | ID of the product, in the uuid format.
| name | string | Name of the product.
| description | string | A brief description of the product. |
| categoryID | string | ID of the product's category. |
| imageURL | string | URL of the product's display image. |
| price | number | Sale price of the product. | 

#### ***Response examples:***

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

---

### Create Product

Creates a new product.

#### URL

`POST /products`

#### Request 

| Parameter | Type | Description | Requirement |
| --------- | ---- | ----------- | ----------- |
| name | string | Name of the product. | Required |
| description | string | A brief description of the product. | Required |
| categoryID | string | ID of the product's category. | Required |
| imageURL | string | URL of the product's display image. | Required |
| price | number | Sale price of the product. | Required |

#### ***Example Request:***

```js
    let formData = new FormData();
    formData.append("name", "Tablet 128GB");
    formData.append("description", "A new generation tablet");
    formData.append("categoryID", "38687117-ecca-440c-8f9b-0b73489d5a2f");
    formData.append("imageURL", "http://example.com/tablet2.jpg");
    formData.append("price", 5000);

    fetch(`http://localhost:8080/categories`, {
        method: 'post',
        body: formData,
    }).then(response => response.json())
    .then(...);
```
#### ***Response examples:***

  - `201 Created`: Returns the created product object
    ```json
        {
          "id": "4fbed1c4-5374-49f6-97e3-1b61c67733df",
          "name": "Tablet 128GB",
          "description": "A new generation tablet",
          "categoryID": "38687117-ecca-440c-8f9b-0b73489d5a2f",
          "imageURL": "http://example.com/tablet2.jpg",
          "price": 5000
        }
    ```

  - `400 Bad Request`: If the request body is invalid.

  - `500 Internal Server Error`: If an error occurs.