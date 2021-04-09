package database

import (
	"github.com/iamrafaelsoares/go-admin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dns := "host=localhost user=postgres password=postgres dbname=go_admin port=5432"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
