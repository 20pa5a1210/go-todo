package handlers

import (
	"net/http"

	"github.com/20pa5a1210/go-todo/repositories"
	"github.com/20pa5a1210/go-todo/utils"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	userRepo repositories.UserRepository
}

func NewProfileHandler(userRepo repositories.UserRepository) *ProfileHandler {
	return &ProfileHandler{
		userRepo: userRepo,
	}
}

func (ph *ProfileHandler) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	user, err := ph.userRepo.GetUserByID(userID.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to Fetch user")
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})

}
