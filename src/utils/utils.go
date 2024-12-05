package utils

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func GetAuthToken(c *gin.Context) *auth.Token {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "could not retrieve token"})
		c.Abort()
		return nil
	}
	authToken, ok := token.(*auth.Token)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "could not retrieve token"})
		c.Abort()
		return nil
	}
	return authToken
}
