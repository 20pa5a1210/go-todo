package main

import (
	"github.com/20pa5a1210/go-todo/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	users := router.Group("/user")
	users.Use()
	{
		users.POST("/create", handlers.CreateUserHandler)
		users.GET("/getall", handlers.GetUserHandler)
		users.GET("/:email", handlers.GetUserByEmail)
	}
	router.Run(":3033") // list
}
