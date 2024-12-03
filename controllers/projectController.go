package controllers

import (
	"github.com/eenees/scwn-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var projects = []models.Project{
	{Id: uuid.New(), Name: "project 1"},
	{Id: uuid.New(), Name: "project 2"},
}

func GetAllProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func GetProject(c *gin.Context) {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
	}
	for _, project := range projects {
		if project.Id == id {
			c.JSON(http.StatusOK, project)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

// FIX: watch out for race conditions when implementing db.
func CreateProject(c *gin.Context) {
	var newProject models.Project
	if err := c.BindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newProject.Id = uuid.New()
	projects = append(projects, newProject)
	c.JSON(http.StatusCreated, newProject)
}
