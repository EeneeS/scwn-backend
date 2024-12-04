package routes

import (
	"net/http"

	projectcontroller "github.com/eenees/scwn-backend/src/controllers/projectController"
	usercontroller "github.com/eenees/scwn-backend/src/controllers/userController"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", getAPIInfo)

	projectRoutes := router.Group("/projects")
	{
		projectRoutes.GET("/", projectcontroller.GetAllProjects)
		projectRoutes.GET("/:id", projectcontroller.GetProject)
		projectRoutes.POST("/", projectcontroller.CreateProject)
		projectRoutes.DELETE("/:id", projectcontroller.DeleteProject)
	}

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", usercontroller.CreateUser)
	}
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
