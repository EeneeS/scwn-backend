package usercontroller

import (
	"github.com/eenees/scwn-backend/src/models"
	"github.com/eenees/scwn-backend/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	authToken := utils.GetAuthToken(c)
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = authToken.UID
	createdUser, err := models.CreateUser(authToken, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}
