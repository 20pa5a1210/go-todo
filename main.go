package main

import (
	"github.com/20pa5a1210/go-todo/handlers"
	"github.com/20pa5a1210/go-todo/middleware"
	"github.com/20pa5a1210/go-todo/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userRepo := repositories.NewUserRepository()
	users := router.Group("/user")
	users.Use()
	{
		profileHandler := handlers.NewProfileHandler(*userRepo)
		users.POST("/create", handlers.CreateUserHandler)
		users.POST("/login", handlers.LoginUser)
		users.GET("/getall", handlers.GetUserHandler)
		users.GET("/:email", handlers.GetUserByEmail)
		// Nested Route
		authGroup := users.Group("/")
		authGroup.Use(middleware.AuthMiddleware)
		{
			authGroup.GET("/profile", profileHandler.GetProfile)
		}
	}

	router.POST("/addtodo/:id", handlers.AddTodo)
	router.Run(":3033") // list
}
