package utils

import (
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, gin.H{"error": errorMessage})
}
func RespondWithJSON(c *gin.Context, statusCode int, key string, payload interface{}) {
	c.JSON(statusCode, gin.H{key: payload})
}
