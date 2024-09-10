package main

import (
	"hacktiv8_golang_assignment_2/database"
	"hacktiv8_golang_assignment_2/router"
	"log"
)

func main() {
	database.StartDB()
	r := router.SetupRouter()

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
