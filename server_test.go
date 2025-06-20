package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestAPI() *gin.Engine {
	// Use in-memory SQLite for testing
	dbTest, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	AutoMigrateModels(dbTest)
	InitDefaults(dbTest)
	// Assign to global db for handlers
	db = dbTest

	gin.SetMode(gin.TestMode)
	r := InitServer()
	return r
}

func TestCustomerEndpoints(t *testing.T) {
	r := setupTestAPI()
	t.Run("GET /customers", func(t *testing.T) {
		w := performRequest(r, "GET", "/customers", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GET /customers/:id (existing)", func(t *testing.T) {
		w := performRequest(r, "GET", "/customers/1", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GET /customers/:id (not found)", func(t *testing.T) {
		w := performRequest(r, "GET", "/customers/9999", nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("POST /customers (invalid)", func(t *testing.T) {
		w := performRequest(r, "POST", "/customers", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("POST /customers (valid)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"Name":"Jane","Surname":"Smith","Email":"jane.smith@example.com"}`)
		w := performRequest(r, "POST", "/customers", jsonBody)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("PUT /customers/:id (update existing)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"ID":2,"Name":"JaneUpdated","Surname":"Smith","Email":"jane.smith@example.com"}`)
		w := performRequest(r, "PUT", "/customers/2", jsonBody)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("PUT /customers/:id (not found)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"ID":9999,"Name":"Ghost","Surname":"User","Email":"ghost@example.com"}`)
		w := performRequest(r, "PUT", "/customers/9999", jsonBody)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("DELETE /customers/:id (existing)", func(t *testing.T) {
		w := performRequest(r, "DELETE", "/customers/2", nil)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
	t.Run("DELETE /customers/:id (not found)", func(t *testing.T) {
		w := performRequest(r, "DELETE", "/customers/9999", nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestCategoryEndpoints(t *testing.T) {
	r := setupTestAPI()
	t.Run("GET /categories", func(t *testing.T) {
		w := performRequest(r, "GET", "/categories", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GET /categories/:id (existing)", func(t *testing.T) {
		w := performRequest(r, "GET", "/categories/1", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GET /categories/:id (not found)", func(t *testing.T) {
		w := performRequest(r, "GET", "/categories/9999", nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("POST /categories (invalid)", func(t *testing.T) {
		w := performRequest(r, "POST", "/categories", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("POST /categories (valid)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"Title":"NewCat","Description":"Desc"}`)
		w := performRequest(r, "POST", "/categories", jsonBody)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("PUT /categories/:id (update existing)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"ID":2,"Title":"UpdatedCat","Description":"Desc"}`)
		w := performRequest(r, "PUT", "/categories/2", jsonBody)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("PUT /categories/:id (not found)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"ID":9999,"Title":"Ghost","Description":"GhostDesc"}`)
		w := performRequest(r, "PUT", "/categories/9999", jsonBody)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("DELETE /categories/:id (existing)", func(t *testing.T) {
		w := performRequest(r, "DELETE", "/categories/2", nil)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
	t.Run("DELETE /categories/:id (not found)", func(t *testing.T) {
		w := performRequest(r, "DELETE", "/categories/9999", nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestItemEndpoints(t *testing.T) {
	r := setupTestAPI()
	t.Run("GET /items", func(t *testing.T) {
		w := performRequest(r, "GET", "/items", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GET /items/:id (existing)", func(t *testing.T) {
		w := performRequest(r, "GET", "/items/1", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GET /items/:id (not found)", func(t *testing.T) {
		w := performRequest(r, "GET", "/items/9999", nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("POST /items (invalid)", func(t *testing.T) {
		w := performRequest(r, "POST", "/items", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("POST /items (valid)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"Title":"NewItem","Description":"Desc","Price":10.5}`)
		w := performRequest(r, "POST", "/items", jsonBody)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("PUT /items/:id (update existing)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"ID":2,"Title":"UpdatedItem","Description":"Desc","Price":20.0}`)
		w := performRequest(r, "PUT", "/items/2", jsonBody)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("PUT /items/:id (not found)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"ID":9999,"Title":"Ghost","Description":"GhostDesc","Price":1.0}`)
		w := performRequest(r, "PUT", "/items/9999", jsonBody)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("DELETE /items/:id (existing)", func(t *testing.T) {
		w := performRequest(r, "DELETE", "/items/2", nil)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
	t.Run("DELETE /items/:id (not found)", func(t *testing.T) {
		w := performRequest(r, "DELETE", "/items/9999", nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestOrderEndpoints(t *testing.T) {
	r := setupTestAPI()
	t.Run("GET /orders", func(t *testing.T) {
		w := performRequest(r, "GET", "/orders", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GET /orders/:id (existing)", func(t *testing.T) {
		w := performRequest(r, "GET", "/orders/1", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GET /orders/:id (not found)", func(t *testing.T) {
		w := performRequest(r, "GET", "/orders/9999", nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("POST /orders (invalid)", func(t *testing.T) {
		w := performRequest(r, "POST", "/orders", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("POST /orders (valid)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"CustomerID":1,"Items":[{"ShopItemID":1,"Quantity":1}]}`)
		w := performRequest(r, "POST", "/orders", jsonBody)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("PUT /orders/:id (update existing)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"ID":2,"CustomerID":1,"Items":[{"ShopItemID":1,"Quantity":2}]}`)
		w := performRequest(r, "PUT", "/orders/2", jsonBody)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("PUT /orders/:id (not found)", func(t *testing.T) {
		jsonBody := bytes.NewBufferString(`{"ID":9999,"CustomerID":1,"Items":[{"ShopItemID":1,"Quantity":1}]}`)
		w := performRequest(r, "PUT", "/orders/9999", jsonBody)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("DELETE /orders/:id (existing)", func(t *testing.T) {
		w := performRequest(r, "DELETE", "/orders/2", nil)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
	t.Run("DELETE /orders/:id (not found)", func(t *testing.T) {
		w := performRequest(r, "DELETE", "/orders/9999", nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

// Helper to perform HTTP requests in tests
func performRequest(r http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	var req *http.Request
	if body != nil {
		if buf, ok := body.(*bytes.Buffer); ok {
			req = httptest.NewRequest(method, path, buf)
			req.Header.Set("Content-Type", "application/json")
		} else {
			req = httptest.NewRequest(method, path, nil)
		}
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Repeat similar test functions for ShopItemCategory, ShopItem, and Order endpoints
