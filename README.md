# GO E-COMMERCE API

## Overview
This repository contains a microservice which manages products and categories for a fictitious application called Full Cycle E-commerce. This service is an API written in Golang responsible for creating and fetching products and categories for the E-commerce platform.

## Table of Contents
- [GO E-COMMERCE API](#go-e-commerce-api)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Technologies](#technologies)
    - [Setup Instructions](#setup-instructions)
  - [API Documentation](#api-documentation)
  - [Future improvements](#future-improvements)
  - [Contributing](#contributing)

## Getting Started

### Technologies
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

## API Documentation
For detailed information about the endpoints provided by this API, please refer to the [API Documentation](API.md)
  
## Future improvements

The functionalities provided by this API are minimal. Here are some known areas for improvement and further development:

- **Error Handling:** Improve error handling across the API, making sure to address all the probable scenarios and provide friendly and clear messages.
- **Testing:** Include unit tests that cover all the functions in the API.
- **Performance:** Check for possible weak spots and improve the overall performance of the queries.

## Contributing

This project is intended for educational purposes, showcasing a simple yet functional API service.
Although this is a study case, we welcome contributions! If you would like to lend a hand on the aforementioned points of improvement, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Commit your changes, making sure to include a descriptive message.
4. Push your changes.
5. Open a pull request to the main repository.
