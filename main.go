package main

import (
	"log"
	"net/http"

	"github.com/20pa5a1210/go-todo/handlers"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/user/create", handlers.CreateUserHandler)
	log.Fatal(http.ListenAndServe(":3033", router))
	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// router.POST("/user/create", func(ctx *gin.Context) {
	// 	handlers.CreateUserHandler(ctx.Writer, ctx.Request)
	// })
	// router.Run(":3033") // list
}
