package main

import (
	"go-restapi/database"
	"go-restapi/models"
	"go-restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.Connect()

	// Auto-migrate the models
	database.DB.AutoMigrate(&models.Book{})

	// Initialize the Gin router
	r := gin.Default()

	// Register routes
	routes.BookRoutes(r)

	// Start the server
	r.Run(":8080")
}
