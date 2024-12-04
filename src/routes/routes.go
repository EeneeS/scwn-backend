package routes

import (
	"net/http"

	projectcontroller "github.com/eenees/scwn-backend/src/controllers/projectController"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", getAPIInfo)
	router.GET("/projects", projectcontroller.GetAllProjects)
	router.GET("/projects/:id", projectcontroller.GetProject)
	router.POST("/projects", projectcontroller.CreateProject)
	router.DELETE("/projects/:id", projectcontroller.DeleteProject)
}

func AuthRoutes(router *gin.Engine) {
	// todo auth here aka checking if user exists
}

func getAPIInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"appname": "scwn",
		"version": "v1.0.0",
	})
}
