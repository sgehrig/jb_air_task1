package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "srg.de/jb/air_task1/models"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    // Customer CRUD
    r.GET("/customers", func(c *gin.Context) {
        var customers []models.Customer
        if err := db.Find(&customers).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, customers)
    })
    r.GET("/customers/:id", func(c *gin.Context) {
        var customer models.Customer
        if err := db.First(&customer, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, customer)
    })
    r.POST("/customers", func(c *gin.Context) {
        var customer models.Customer
        if err := c.ShouldBindJSON(&customer); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Create(&customer).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, customer)
    })
    r.PUT("/customers/:id", func(c *gin.Context) {
        var customer models.Customer
        if err := db.First(&customer, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        if err := c.ShouldBindJSON(&customer); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Save(&customer).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, customer)
    })
    r.DELETE("/customers/:id", func(c *gin.Context) {
        if err := db.Delete(&models.Customer{}, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    })

    // ShopItemCategory CRUD
    r.GET("/categories", func(c *gin.Context) {
        var categories []models.ShopItemCategory
        if err := db.Find(&categories).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, categories)
    })
    r.GET("/categories/:id", func(c *gin.Context) {
        var category models.ShopItemCategory
        if err := db.First(&category, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, category)
    })
    r.POST("/categories", func(c *gin.Context) {
        var category models.ShopItemCategory
        if err := c.ShouldBindJSON(&category); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Create(&category).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, category)
    })
    r.PUT("/categories/:id", func(c *gin.Context) {
        var category models.ShopItemCategory
        if err := db.First(&category, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        if err := c.ShouldBindJSON(&category); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Save(&category).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, category)
    })
    r.DELETE("/categories/:id", func(c *gin.Context) {
        if err := db.Delete(&models.ShopItemCategory{}, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    })

    // ShopItem CRUD
    r.GET("/items", func(c *gin.Context) {
        var items []models.ShopItem
        if err := db.Preload("Categories").Find(&items).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, items)
    })
    r.GET("/items/:id", func(c *gin.Context) {
        var item models.ShopItem
        if err := db.Preload("Categories").First(&item, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, item)
    })
    r.POST("/items", func(c *gin.Context) {
        var item models.ShopItem
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Create(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, item)
    })
    r.PUT("/items/:id", func(c *gin.Context) {
        var item models.ShopItem
        if err := db.First(&item, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Save(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, item)
    })
    r.DELETE("/items/:id", func(c *gin.Context) {
        if err := db.Delete(&models.ShopItem{}, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    })

    // Order CRUD
    r.GET("/orders", func(c *gin.Context) {
        var orders []models.Order
        if err := db.Preload("Customer").Preload("Items.ShopItem").Find(&orders).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, orders)
    })
    r.GET("/orders/:id", func(c *gin.Context) {
        var order models.Order
        if err := db.Preload("Customer").Preload("Items.ShopItem").First(&order, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, order)
    })
    r.POST("/orders", func(c *gin.Context) {
        var order models.Order
        if err := c.ShouldBindJSON(&order); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Create(&order).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, order)
    })
    r.PUT("/orders/:id", func(c *gin.Context) {
        var order models.Order
        if err := db.First(&order, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        if err := c.ShouldBindJSON(&order); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Save(&order).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, order)
    })
    r.DELETE("/orders/:id", func(c *gin.Context) {
        if err := db.Delete(&models.Order{}, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    })
}
