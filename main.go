package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := AutoMigrateModels(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	if err := InitDefaults(db); err != nil {
		log.Fatalf("failed to initialize default data: %v", err)
	}

	log.Println("Database initialized and ready.")
}
