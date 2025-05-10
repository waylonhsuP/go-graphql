package seeds

import (
	"event-trigger-demo/models"
)

func RunSeeding() {
	models.DB.Migrator().DropTable(&models.User{}, &models.Todo{})
	models.DB.AutoMigrate(&models.User{}, &models.Todo{})
	SeedUsers()
	SeedTodos()
}
