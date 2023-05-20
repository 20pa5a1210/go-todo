package middleware

import (
	"net/http"
	"time"

	"github.com/20pa5a1210/go-todo/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

var (
	secretKey = "your-secret-key"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header is missing")
		c.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}
	
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		utils.RespondWithError(c, http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}

	c.Set("user_id", claims.UserID)
	c.Next()
}

func GenerateJwt(UserID string) (string, error) {
	claims := &Claims{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
