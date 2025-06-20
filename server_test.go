package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gin.Engine {
	// Use in-memory SQLite for testing
	dbTest, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	AutoMigrateModels(dbTest)
	InitDefaults(dbTest)
	// Assign to global db for handlers
	db = dbTest

	return InitServer()
}

func TestCustomerEndpoints(t *testing.T) {
	r := setupTestDB()
	w := performRequest(r, "GET", "/customers", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions and tests for each endpoint
}

// Helper to perform HTTP requests in tests
func performRequest(r http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Repeat similar test functions for ShopItemCategory, ShopItem, and Order endpoints
