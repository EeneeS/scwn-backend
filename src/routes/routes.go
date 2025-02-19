package routes

import (
	changecontroller "github.com/eenees/scwn-backend/src/controllers/changeController"
	projectcontroller "github.com/eenees/scwn-backend/src/controllers/projectController"
	publishtargetcontroller "github.com/eenees/scwn-backend/src/controllers/publishTargetController"
	usercontroller "github.com/eenees/scwn-backend/src/controllers/userController"
	"github.com/eenees/scwn-backend/src/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", getAPIInfo)
}

func AuthRoutes(router *gin.Engine) {
	projectRoutes := router.Group("/projects")
	projectRoutes.Use(middleware.AuthMiddleware())
	{
		projectRoutes.GET("/", projectcontroller.GetAllProjects)
		projectRoutes.GET("/:project_id", projectcontroller.GetProject)
		projectRoutes.POST("/", projectcontroller.CreateProject)
		projectRoutes.DELETE("/:project_id", projectcontroller.DeleteProject)
		projectRoutes.POST("/:project_id/publish_targets", publishtargetcontroller.CreatePublishTarget)
	}

	router.POST("projects/:project_id/changes", changecontroller.CreateChange)
	// router.GET("projects/:project_id/changes", changecontroller.GetAllChanges)

	userRoutes := router.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.POST("/", usercontroller.CreateUser)
	}
}

func getAPIInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"appname": "scwn",
		"version": "1.0.0",
	})
}
