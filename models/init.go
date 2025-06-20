package models

import (
    "gorm.io/gorm"
)

func InitDefaults(db *gorm.DB) error {
    db.Transaction(func(tx *gorm.DB) error {

        var count int64
        tx.Model(&Customer{}).Count(&count)
        if count > 0 {
            return nil
        }

        customers := []Customer{
            {Name: "John", Surname: "Doe", Email: "john.doe@example.com"},
            {Name: "Jane", Surname: "Smith", Email: "jane.smith@example.com"},
        }
        if err := tx.Create(customers).Error; err != nil {
            return err
        }

        categories := []ShopItemCategory{
            {Title: "Electronics", Description: "Electronic items"},
            {Title: "Books", Description: "Books and literature"},
        }
        if err := tx.Create(categories).Error; err != nil {
            return err
        }

        var electronics, books ShopItemCategory
        if err := tx.First(&electronics, "title = ?", "Electronics").Error; err != nil {
            return err
        }
        if err := tx.First(&books, "title = ?", "Books").Error; err != nil {
            return err
        }

        items := []*ShopItem{
            {
                Title:       "Laptop",
                Description: "A powerful laptop",
                Price:       1200.0,
                Categories:  []ShopItemCategory{electronics},
            },
            {
                Title:       "Novel",
                Description: "A best-selling novel",
                Price:       20.0,
                Categories:  []ShopItemCategory{books},
            },
        }
        if err := tx.Create(items).Error; err != nil {
            return err
        }

        var customer Customer
        var laptop, novel ShopItem
        if err := tx.First(&customer).Error; err != nil {
            return err
        }
        if err := tx.First(&laptop, "title = ?", "Laptop").Error; err != nil {
            return err
        }
        if err := tx.First(&novel, "title = ?", "Novel").Error; err != nil {
            return err
        }

        order := Order{
            Customer: customer,
            Items: []OrderItem{
                {ShopItem: laptop, Quantity: 2},
                {ShopItem: novel, Quantity: 4},
            },
        }
        tx.Create(&order)

        // return nil will commit the whole transaction
        return nil
    })
    return nil
}
