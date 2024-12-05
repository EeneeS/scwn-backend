package usercontroller

import (
	"github.com/eenees/scwn-backend/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO: this will probaly change depending on
// what firebase returns to a user after
// login in.
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}
