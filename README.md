# Air Task 1 API

This project provides a REST API for managing Customers, ShopItemCategories, ShopItems, and Orders using Go, Gin, GORM, and SQLite.

## Prerequisites

- Go 1.24 or newer
- [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io/) (installed via `go mod tidy`)
- No external database setup required (uses SQLite)

## Setup

1. Clone the repository:
   ```sh
   git clone <your-repo-url>
   cd air-research-task1
   ```

2. Download dependencies:
   ```sh
   go mod tidy
   ```

## Running the API

Start the API server with:
```sh
go run main.go
```
The server will start (by default) on `localhost:8080`.

## API Endpoints

- `GET    /customers`
- `POST   /customers`
- `GET    /customers/:id`
- `PUT    /customers/:id`
- `DELETE /customers/:id`
- `GET    /categories`
- `POST   /categories`
- `GET    /categories/:id`
- `PUT    /categories/:id`
- `DELETE /categories/:id`
- `GET    /items`
- `POST   /items`
- `GET    /items/:id`
- `PUT    /items/:id`
- `DELETE /items/:id`
- `GET    /orders`
- `POST   /orders`
- `GET    /orders/:id`
- `PUT    /orders/:id`
- `DELETE /orders/:id`

## Running Tests

To run all API endpoint tests:
```sh
go test ./api
```

## Notes

- The database is automatically initialized with some default data if empty.
- All data is stored in a local SQLite file (`db.sqlite` by default).
- No additional configuration is required.
