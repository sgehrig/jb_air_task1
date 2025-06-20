package main

import (
	"gorm.io/gorm"
)

// Customer represents a customer in the system
type Customer struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Surname string
	Email   string `gorm:"uniqueIndex"`
	Orders  []Order
}

// ShopItemCategory represents a category for shop items
type ShopItemCategory struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	ShopItems   []*ShopItem `gorm:"many2many:shop_item_categories_items;"`
}

// ShopItem represents an item in the shop
type ShopItem struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Price       float64
	Categories  []*ShopItemCategory `gorm:"many2many:shop_item_categories_items;"`
	OrderItems  []OrderItem
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID         uint `gorm:"primaryKey"`
	ShopItemID uint
	ShopItem   ShopItem
	Quantity   int
	OrderID    uint
}

// Order represents a customer order
type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Customer   Customer
	Items      []OrderItem
}

// AutoMigrateModels migrates all database models
func AutoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&Customer{},
		&ShopItemCategory{},
		&ShopItem{},
		&OrderItem{},
		&Order{},
	)
}

// InitDefaults initializes the database with some default data if empty
func InitDefaults(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var count int64
		// Check if there are any customers
		tx.Model(&Customer{}).Count(&count)
		if count > 0 {
			return nil // Data already exists
		}

		// Create default categories
		categories := []ShopItemCategory{
			{Title: "Electronics", Description: "Electronic items"},
			{Title: "Books", Description: "Books and literature"},
		}
		if err := tx.Create(&categories).Error; err != nil {
			return err
		}

		// Create default shop items
		items := []ShopItem{
			{Title: "Laptop", Description: "A powerful laptop", Price: 1200.0, Categories: []*ShopItemCategory{&categories[0]}},
			{Title: "Novel", Description: "A best-selling novel", Price: 20.0, Categories: []*ShopItemCategory{&categories[1]}},
		}
		if err := tx.Create(&items).Error; err != nil {
			return err
		}

		// Create a default customer
		customer := Customer{Name: "John", Surname: "Doe", Email: "john.doe@example.com"}
		if err := tx.Create(&customer).Error; err != nil {
			return err
		}

		// Create order items
		orderItems := []OrderItem{
			{ShopItemID: items[0].ID, Quantity: 1},
			{ShopItemID: items[1].ID, Quantity: 2},
		}

		// Create an order for the customer
		order := Order{
			CustomerID: customer.ID,
			Items:      orderItems,
		}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})
}
