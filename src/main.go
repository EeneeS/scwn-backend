package main

import (
	"database/sql"
	"fmt"
	"github.com/eenees/scwn-backend/src/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqluser := os.Getenv("MYSQL_USER")
	sqlPassword := os.Getenv("MYSQL_PASSWORD")
	sqlDB := os.Getenv("MYSQL_DATABASE")

	sqlConnString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", sqluser, sqlPassword, sqlDB)

	db, err := sql.Open("mysql", sqlConnString)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	router := gin.Default()

	router.GET("/", getAPIInfo)
	router.GET("/projects", controllers.GetAllProjects)
	router.GET("/projects/:id", controllers.GetProject)
	router.POST("/projects", controllers.CreateProject)

	router.Run("localhost:3000")
}

func getAPIInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"appname": "scwn",
		"version": "v1.0.0",
	})
}

