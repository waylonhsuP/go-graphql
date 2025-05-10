package main

import (
	"event-trigger-demo/controllers"
	"event-trigger-demo/database"
	"event-trigger-demo/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&models.User{})

	// Initialize router
	r := gin.Default()

	// Create API group
	api := r.Group("/api")

	// User routes
	userController := controllers.NewUserController(db)
	api.GET("/users", userController.GetUsers)
	api.GET("/users/:id", userController.GetUser)
	api.POST("/users", userController.CreateUser)
	api.PUT("/users/:id", userController.UpdateUser)
	api.DELETE("/users/:id", userController.DeleteUser)

	// Get port from environment variable or use default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 