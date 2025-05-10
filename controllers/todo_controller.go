package controllers

import (
	"event-trigger-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoController struct {
	db *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{db: db}
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	var todos []models.Todo
	if err := c.db.Preload("User").Find(&todos).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo models.Todo
	if err := c.db.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}