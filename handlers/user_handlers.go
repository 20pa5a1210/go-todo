package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/20pa5a1210/go-todo/models"
	"github.com/20pa5a1210/go-todo/repositories"
	"github.com/20pa5a1210/go-todo/utils"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if user.Password != user.ConfirmPassword {
		w.WriteHeader(http.StatusForbidden)
		utils.RespondWithError(w, http.StatusForbidden, "Passwords Mismacth")
		return
	}

	userRepo := repositories.NewUserRepository()
	existingUser, _ := userRepo.GetUserByEmail(user.Email)

	if existingUser.Email != "" {
		w.WriteHeader(http.StatusConflict)
		utils.RespondWithError(w, http.StatusConflict, "User Already Exists")
		return
	}

	createdUser, err := userRepo.CreateUser(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, createdUser)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userRepo := repositories.NewUserRepository()
	users, err := userRepo.GetUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, users)
}
