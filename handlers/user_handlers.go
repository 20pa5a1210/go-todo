package handlers

import (
	"net/http"

	"github.com/20pa5a1210/go-todo/middleware"
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
	todoRepo := repositories.NewTodoRepository()
	todoInstance := models.Todo{
		Email: createdUser.Email,
		Todos: []string{},
	}
	_, err = todoRepo.CreateTodoInstance(todoInstance)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed To Create Todo Instance")
	}
	utils.RespondWithJSON(c, http.StatusCreated, createdUser)
}

func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid Payload")
		return
	}
	userRepo := repositories.NewUserRepository()
	user, err := userRepo.GetUserByEmail(loginData.Email)

	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid Credintials")
		return
	}
	if user.Password != loginData.Password {
		utils.RespondWithError(c, http.StatusUnauthorized, "wrong Password(mismatch)")
		return
	}
	token, err := middleware.GenerateJwt(user.Id.Hex())
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed To generate Token")
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
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
