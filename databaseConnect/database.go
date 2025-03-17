package databaseConnect

import (
	"freshers-bootcamp/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:Rsk@4436@tcp(127.0.0.1:3306)/retailer_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = db
	DB.AutoMigrate(&models.Product{}, &models.Order{})
}
