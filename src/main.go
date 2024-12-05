package main

import (
	"github.com/eenees/scwn-backend/src/config"
	"github.com/eenees/scwn-backend/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.LoadEnv()
	config.ConnectDatabase()
	config.ConnectFirebase()

	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.Run("localhost:3000")
}
