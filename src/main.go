package main

import (
	"net/http"

	projectcontroller "github.com/eenees/scwn-backend/src/controllers/projectController"
	"github.com/eenees/scwn-backend/src/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// TODO: dont put this in models...
	models.LoadEnv()
	models.ConnectDatabase()
	models.ConnectFirebase()

	router.GET("/", getAPIInfo)
	router.GET("/projects", projectcontroller.GetAllProjects)
	router.GET("/projects/:id", projectcontroller.GetProject)
	router.POST("/projects", projectcontroller.CreateProject)
	router.DELETE("/projects/:id", projectcontroller.DeleteProject)
	router.Run("localhost:3000")
}

func getAPIInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"appname": "scwn",
		"version": "v1.0.0",
	})
}
