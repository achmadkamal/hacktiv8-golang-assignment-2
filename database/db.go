package database

import (
	"fmt"
	"hacktiv8_golang_assignment_2/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=hacktiv8_golang_assignment_2 port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	err = DB.AutoMigrate(&entity.Order{}, &entity.Item{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	fmt.Println("Database connected")
}
