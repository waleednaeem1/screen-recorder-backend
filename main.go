package main

import (
	"screen-recorder-backend/config"
	"screen-recorder-backend/routes"
	"screen-recorder-backend/migrations"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load")
	}

	// Connect to the database
	config.ConnectDB()

	// Auto migrate all models
	err = migrations.AutoMigrateModels() // Call the new migration function
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	// Initialize Gin
	router := gin.Default()

	// Register routes
	routes.AuthRoutes(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
	}
	router.Run(":" + port)
}