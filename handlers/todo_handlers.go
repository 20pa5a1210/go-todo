package handlers

import (
	"net/http"

	"github.com/20pa5a1210/go-todo/models"
	"github.com/20pa5a1210/go-todo/repositories"
	"github.com/20pa5a1210/go-todo/utils"
	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	userId := c.Param("id")
	var todo models.Todos
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid payload")
		return
	}
	todoRepo := repositories.NewTodoRepository()
	updatedUser, err := todoRepo.AddTodo(userId, todo)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to add todo")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, updatedUser)
}
