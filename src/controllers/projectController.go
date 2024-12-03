package controllers

import (
	"github.com/eenees/scwn-backend/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var projects = []models.Project{
	{Id: uuid.New(), Name: "project 1"},
	{Id: uuid.New(), Name: "project 2"},
}

func GetAllProjects(c *gin.Context) {
	projectsDB, err := models.GetAllProjects()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, projectsDB)
}

func GetProject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	projectDB, err := models.GetProject(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projectDB)
}

// FIX: watch out for race conditions when implementing db.
func CreateProject(c *gin.Context) {
	var newProject models.Project

	if err := c.BindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProject.Id = uuid.New()

	newProjectDB, err := models.CreateProject(newProject)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusCreated, newProjectDB)
}
