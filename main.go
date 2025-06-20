package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "srg.de/jb/air_task1/models"
    "srg.de/jb/air_task1/api"
)

func main() {
    db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    err = models.AutoMigrate(db)
    if err != nil {
        panic("failed to migrate schema")
    }

    // Initialize defaults
    err = models.InitDefaults(db)
    if err != nil {
        panic("failed to initialize defaults")
    }

    r := gin.Default()
    r.SetTrustedProxies(nil)
    api.RegisterRoutes(r, db)
    r.Run()

}
