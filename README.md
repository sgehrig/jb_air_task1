# Air Research Task 1 API

This project is a Go REST API using GORM (with SQLite) and Gin, providing full CRUD for Customers, ShopItemCategories, ShopItems, and Orders.

## Prerequisites
- Go 1.24 or newer

## Setup
1. Clone the repository or copy the project files to your machine.
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Running the API
1. Start the API server:
   ```sh
   go run main.go
   ```
   The server will start on `http://localhost:8080`.

2. The SQLite database file (`db.sqlite`) will be created automatically in the project directory.

## API Endpoints
- `GET    /customers`, `POST   /customers`, `GET    /customers/:id`, `PUT    /customers/:id`, `DELETE /customers/:id`
- `GET    /categories`, `POST   /categories`, `GET    /categories/:id`, `PUT    /categories/:id`, `DELETE /categories/:id`
- `GET    /items`, `POST   /items`, `GET    /items/:id`, `PUT    /items/:id`, `DELETE /items/:id`
- `GET    /orders`, `POST   /orders`, `GET    /orders/:id`, `PUT    /orders/:id`, `DELETE /orders/:id`

## Running the Tests
1. Run all tests with:
   ```sh
   go test -v ./...
   ```
   This will run endpoint and error scenario tests using an in-memory SQLite database.

## Notes
- The API auto-creates tables and seeds some default data on first run.
- All error scenarios (not found, invalid data, etc.) are handled with appropriate HTTP status codes.
- You can use tools like curl, Postman, or httpie to interact with the API.

## Troubleshooting
- If you see missing dependency errors, run `go mod tidy` again.
- Make sure you are running Go 1.18 or newer.

---

For any issues, please check the code comments or open an issue.
