package main

import (
	"github.com/eenees/scwn-backend/src/models"
	"github.com/eenees/scwn-backend/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// TODO: dont put this in models...
	models.LoadEnv()
	models.ConnectDatabase()
	models.ConnectFirebase()

	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.Run("localhost:3000")
}
