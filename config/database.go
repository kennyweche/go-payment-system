package config

import (
	"fmt"
	"log"
	"payment-system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "dev:Portmore254@@tcp(localhost:3306)/payments_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// Auto-create missing tables
	err = DB.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database connected & migrated successfully")
}
