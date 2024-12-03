package main

import (
	"github.com/eenees/scwn-backend/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", getAPIInfo)
	router.GET("/projects", controllers.GetAllProjects)
	router.GET("/projects/:id", controllers.GetProject)
	router.Run("localhost:3000")
}

func getAPIInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"appname": "scwn",
		"version": "v1.0.0",
	})
}
