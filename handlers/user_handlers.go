package handlers

import (
	"net/http"

	"github.com/20pa5a1210/go-todo/models"
	"github.com/20pa5a1210/go-todo/repositories"
	"github.com/20pa5a1210/go-todo/utils"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if user.Password != user.ConfirmPassword {
		utils.RespondWithError(c, http.StatusForbidden, "Passwords Mismacth")
		return
	}

	userRepo := repositories.NewUserRepository()
	existingUser, _ := userRepo.GetUserByEmail(user.Email)

	if existingUser.Email != "" {
		utils.RespondWithError(c, http.StatusConflict, "User Already Exists")
		return
	}

	createdUser, err := userRepo.CreateUser(user)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	utils.RespondWithJSON(c, http.StatusCreated, createdUser)
}

func GetUserHandler(c *gin.Context) {
	userRepo := repositories.NewUserRepository()
	users, err := userRepo.GetUsers()
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to retrieve Users")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, users)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	userRepo := repositories.NewUserRepository()
	user, err := userRepo.GetUserByEmail(email)
	if err != nil {
		utils.RespondWithError(c, http.StatusForbidden, "Failed To retrieve User")
		return
	}
	if user.Email == "" {
		utils.RespondWithError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, user)
}
