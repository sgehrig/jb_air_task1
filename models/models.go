package models

import (
    "gorm.io/gorm"
)

type Customer struct {
    ID      uint `gorm:"primaryKey"`
    Name    string
    Surname string
    Email   string
}

type ShopItemCategory struct {
    ID          uint `gorm:"primaryKey"`
    Title       string
    Description string
}

type ShopItem struct {
    ID          uint `gorm:"primaryKey"`
    Title       string
    Description string
    Price       float64
    Categories  []ShopItemCategory `gorm:"many2many:shop_items_shop_item_categories;"`
}

type OrderItem struct {
    ID         uint `gorm:"primaryKey"`
    OrderID    uint
    Order      Order `gorm:"foreignKey:OrderID"`
    ShopItemID uint
    ShopItem   ShopItem `gorm:"foreignKey:ShopItemID"`
    Quantity   int
}

type Order struct {
    ID         uint `gorm:"primaryKey"`
    CustomerID uint
    Customer   Customer `gorm:"foreignKey:CustomerID"`
    Items      []OrderItem
}

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &Customer{},
        &ShopItemCategory{},
        &ShopItem{},
        &OrderItem{},
        &Order{},
    )
}
