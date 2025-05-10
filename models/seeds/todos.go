package seeds

import (
	"event-trigger-demo/models"
	"log"
)

func SeedTodos() {
	todos := []models.Todo{
		{
			Text: "Buy groceries",
			Done: false,
			User: &models.User{
				ID: 1,
			},
		},
		{
			Text: "Buy Pen",
			Done: false,
			User: &models.User{
				ID: 1,
			},
		},
		{
			Text: "Buy Food",
			Done: false,
			User: &models.User{
				ID: 1,
			},
		},
		{
			Text: "Buy Shoes",
			Done: false,
			User: &models.User{
				ID: 2,
			},
		},
		{
			Text: "Buy Dog",
			Done: false,
			User: &models.User{
				ID: 2,
			},
		},
	}
	for _, todo := range todos {
		result := models.DB.Create(&todo)
		if result.Error != nil {
			log.Printf("Error seeding user %s: %v", todo.Text, result.Error)
		}
	}
}