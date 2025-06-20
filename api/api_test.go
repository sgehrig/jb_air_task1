package api

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "srg.de/jb/air_task1/models"
    "github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, *gorm.DB) {
    gin.SetMode(gin.TestMode)
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    db.AutoMigrate(&models.Customer{}, &models.ShopItemCategory{}, &models.ShopItem{}, &models.OrderItem{}, &models.Order{})
    models.InitDefaults(db)
    r := gin.Default()
    RegisterRoutes(r, db)
    return r, db
}

func TestCustomerCRUD(t *testing.T) {
    r, _ := setupTestRouter()
    testCustomerCreate(r, t)
    testCustomerCreateInvalid(r, t)
    testCustomerList(r, t)
    testCustomerGet(r, t)
    testCustomerGetNotFound(r, t)
    testCustomerUpdate(r, t)
    testCustomerUpdateNotFound(r, t)
    testCustomerUpdateInvalid(r, t)
    testCustomerDelete(r, t)
    testCustomerDeleteNotFound(r, t)
}

func testCustomerCreate(r *gin.Engine, t *testing.T) {
    customer := models.Customer{Name: "Test", Surname: "User", Email: "test@user.com"}
    body, _ := json.Marshal(customer)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/customers", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 201, w.Code)
}

func testCustomerCreateInvalid(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/customers", bytes.NewReader([]byte(`{"Name":123}`)))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 400, w.Code)
}

func testCustomerList(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/customers", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testCustomerGet(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/customers/1", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testCustomerGetNotFound(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/customers/9999", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 404, w.Code)
}

func testCustomerUpdate(r *gin.Engine, t *testing.T) {
    customer := models.Customer{Name: "Updated", Surname: "User", Email: "test@user.com"}
    body, _ := json.Marshal(customer)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/customers/1", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testCustomerUpdateNotFound(r *gin.Engine, t *testing.T) {
    customer := models.Customer{Name: "Updated", Surname: "User", Email: "test@user.com"}
    body, _ := json.Marshal(customer)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/customers/9999", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 404, w.Code)
}

func testCustomerUpdateInvalid(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/customers/1", bytes.NewReader([]byte(`{"Name":123}`)))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 400, w.Code)
}

func testCustomerDelete(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/customers/1", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 204, w.Code)
}

func testCustomerDeleteNotFound(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/customers/9999", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 204, w.Code)
}

// --- Category ---

func TestCategoryCRUD(t *testing.T) {
    r, _ := setupTestRouter()
    testCategoryCreate(r, t)
    testCategoryCreateInvalid(r, t)
    testCategoryList(r, t)
    testCategoryGet(r, t)
    testCategoryGetNotFound(r, t)
    testCategoryUpdate(r, t)
    testCategoryUpdateNotFound(r, t)
    testCategoryUpdateInvalid(r, t)
    testCategoryDelete(r, t)
    testCategoryDeleteNotFound(r, t)
}

func testCategoryCreate(r *gin.Engine, t *testing.T) {
    category := models.ShopItemCategory{Title: "TestCat", Description: "desc"}
    body, _ := json.Marshal(category)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/categories", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 201, w.Code)
}

func testCategoryCreateInvalid(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/categories", bytes.NewReader([]byte(`{"Title":123}`)))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 400, w.Code)
}

func testCategoryList(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/categories", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testCategoryGet(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/categories/1", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testCategoryGetNotFound(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/categories/9999", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 404, w.Code)
}

func testCategoryUpdate(r *gin.Engine, t *testing.T) {
    category := models.ShopItemCategory{Title: "Updated", Description: "desc"}
    body, _ := json.Marshal(category)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/categories/1", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testCategoryUpdateNotFound(r *gin.Engine, t *testing.T) {
    category := models.ShopItemCategory{Title: "Updated", Description: "desc"}
    body, _ := json.Marshal(category)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/categories/9999", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 404, w.Code)
}

func testCategoryUpdateInvalid(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/categories/1", bytes.NewReader([]byte(`{"Title":123}`)))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 400, w.Code)
}

func testCategoryDelete(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/categories/1", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 204, w.Code)
}

func testCategoryDeleteNotFound(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/categories/9999", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 204, w.Code)
}

// --- ShopItem ---

func TestShopItemCRUD(t *testing.T) {
    r, db := setupTestRouter()
    testShopItemCreate(r, db, t)
    testShopItemCreateInvalid(r, t)
    testShopItemList(r, t)
    testShopItemGet(r, t)
    testShopItemGetNotFound(r, t)
    testShopItemUpdate(r, t)
    testShopItemUpdateNotFound(r, t)
    testShopItemUpdateInvalid(r, t)
    testShopItemDelete(r, t)
    testShopItemDeleteNotFound(r, t)
}

func testShopItemCreate(r *gin.Engine, db *gorm.DB, t *testing.T) {
    cat := models.ShopItemCategory{Title: "Cat", Description: "desc"}
    db.Create(&cat)
    item := models.ShopItem{Title: "Item", Description: "desc", Price: 10.0, Categories: []models.ShopItemCategory{cat}}
    body, _ := json.Marshal(item)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/items", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 201, w.Code)
}

func testShopItemCreateInvalid(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/items", bytes.NewReader([]byte(`{"Title":123}`)))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 400, w.Code)
}

func testShopItemList(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/items", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testShopItemGet(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/items/1", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testShopItemGetNotFound(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/items/9999", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 404, w.Code)
}

func testShopItemUpdate(r *gin.Engine, t *testing.T) {
    item := models.ShopItem{Title: "Updated", Description: "desc", Price: 10.0}
    body, _ := json.Marshal(item)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/items/1", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testShopItemUpdateNotFound(r *gin.Engine, t *testing.T) {
    item := models.ShopItem{Title: "Updated", Description: "desc", Price: 10.0}
    body, _ := json.Marshal(item)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/items/9999", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 404, w.Code)
}

func testShopItemUpdateInvalid(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/items/1", bytes.NewReader([]byte(`{"Title":123}`)))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 400, w.Code)
}

func testShopItemDelete(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/items/1", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 204, w.Code)
}

func testShopItemDeleteNotFound(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/items/9999", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 204, w.Code)
}

// --- Order ---

func TestOrderCRUD(t *testing.T) {
    r, db := setupTestRouter()
    testOrderCreate(r, db, t)
    testOrderCreateInvalid(r, t)
    testOrderList(r, t)
    testOrderGet(r, t)
    testOrderGetNotFound(r, t)
    testOrderUpdate(r, t)
    testOrderUpdateNotFound(r, t)
    testOrderUpdateInvalid(r, t)
    testOrderDelete(r, t)
    testOrderDeleteNotFound(r, t)
}

func testOrderCreate(r *gin.Engine, db *gorm.DB, t *testing.T) {
    cust := models.Customer{Name: "Order", Surname: "User", Email: "order@user.com"}
    db.Create(&cust)
    cat := models.ShopItemCategory{Title: "Cat", Description: "desc"}
    db.Create(&cat)
    item := models.ShopItem{Title: "Item", Description: "desc", Price: 10.0, Categories: []models.ShopItemCategory{cat}}
    db.Create(&item)
    order := models.Order{
        CustomerID: cust.ID,
        Items:      []models.OrderItem{{ShopItemID: item.ID, Quantity: 2}},
    }
    body, _ := json.Marshal(order)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/orders", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 201, w.Code)
}

func testOrderCreateInvalid(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/orders", bytes.NewReader([]byte(`{"CustomerID":"abc"}`)))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 400, w.Code)
}

func testOrderList(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/orders", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testOrderGet(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/orders/1", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testOrderGetNotFound(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/orders/9999", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 404, w.Code)
}

func testOrderUpdate(r *gin.Engine, t *testing.T) {
    order := models.Order{
        CustomerID: 1,
        Items:      []models.OrderItem{{ShopItemID: 1, Quantity: 3}},
    }
    body, _ := json.Marshal(order)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

func testOrderUpdateNotFound(r *gin.Engine, t *testing.T) {
    order := models.Order{
        CustomerID: 1,
        Items:      []models.OrderItem{{ShopItemID: 1, Quantity: 3}},
    }
    body, _ := json.Marshal(order)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/orders/9999", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 404, w.Code)
}

func testOrderUpdateInvalid(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewReader([]byte(`{"CustomerID":"abc"}`)))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)
    assert.Equal(t, 400, w.Code)
}

func testOrderDelete(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/orders/1", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 204, w.Code)
}

func testOrderDeleteNotFound(r *gin.Engine, t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/orders/9999", nil)
    r.ServeHTTP(w, req)
    assert.Equal(t, 204, w.Code)
}
