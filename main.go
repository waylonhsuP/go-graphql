package main

import (
	"event-trigger-demo/controllers"
	"event-trigger-demo/models"
	"event-trigger-demo/models/seeds"
	"event-trigger-demo/graph"
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
	
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	models.InitDB()

	if os.Getenv("APP_ENV") == "development" {
		seeds.RunSeeding()
	}

	r := gin.Default()

	api := r.Group("/api")

	// User routes
	userController := controllers.NewUserController(models.DB)
	api.GET("/users", userController.GetUsers)
	api.GET("/users/:id", userController.GetUser)
	api.POST("/users", userController.CreateUser)
	api.PUT("/users/:id", userController.UpdateUser)
	api.DELETE("/users/:id", userController.DeleteUser)

	// Todo routes
	todoController := controllers.NewTodoController(models.DB)
	api.GET("/todos", todoController.GetTodos)
	api.GET("/todos/:id", todoController.GetTodo)

	// GraphQL routes
	graphql := r.Group("/graphql")
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	graphql.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		playground.Handler("GraphQL playground", "/graphql/query").ServeHTTP(c.Writer, c.Request)
	})
	graphql.Any("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("APP_ENV:", os.Getenv("APP_ENV"))

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