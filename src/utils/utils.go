package utils

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func GetAuthToken(c *gin.Context) (*auth.Token, bool) {
	token, exists := c.Get("token")
	if !exists {
		return nil, false
	}
	authToken, ok := token.(*auth.Token)
	return authToken, ok
}
