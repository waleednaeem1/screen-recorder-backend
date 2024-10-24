package main

import (
	"log"
	"screen-recorder-backend/config"
	"screen-recorder-backend/migrations"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.ConnectDB()

	if err := migrations.AutoMigrateModels(); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	log.Println("Migrations completed successfully.")
}