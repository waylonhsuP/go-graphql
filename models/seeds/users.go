package seeds

import (
	"event-trigger-demo/models"
	"log"
)

func SeedUsers() {
	users := []models.User{
		{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123", // In production, this should be hashed
		},
		{
			Name:     "Jane Smith",
			Email:    "jane@example.com",
			Password: "password456", // In production, this should be hashed
		},
	}

	for _, user := range users {
		result := models.DB.Create(&user)
		if result.Error != nil {
			log.Printf("Error seeding user %s: %v", user.Email, result.Error)
		}
	}
}