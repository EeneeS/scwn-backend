package routes

import (
	"net/http"

	projectcontroller "github.com/eenees/scwn-backend/src/controllers/projectController"
	usercontroller "github.com/eenees/scwn-backend/src/controllers/userController"
	"github.com/eenees/scwn-backend/src/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", getAPIInfo)
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", usercontroller.CreateUser)
	}
}

func AuthRoutes(router *gin.Engine) {
	projectRoutes := router.Group("/projects")
	projectRoutes.Use(middleware.AuthMiddleware())
	{
		projectRoutes.GET("/", projectcontroller.GetAllProjects)
		projectRoutes.GET("/:id", projectcontroller.GetProject)
		projectRoutes.POST("/", projectcontroller.CreateProject)
		projectRoutes.DELETE("/:id", projectcontroller.DeleteProject)
	}
}

func getAPIInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"appname": "scwn",
		"version": "v1.0.0",
	})
}
