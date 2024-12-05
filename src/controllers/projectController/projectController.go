package projectcontroller

import (
	"fmt"
	"net/http"

	"github.com/eenees/scwn-backend/src/models"
	"github.com/eenees/scwn-backend/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllProjects(c *gin.Context) {
	authToken := utils.GetAuthToken(c)
	projects, err := models.GetAllProjects(authToken.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func GetProject(c *gin.Context) {
	projectId, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	project, err := models.GetProject(projectId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, project)
}

func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project.Id = uuid.New()
	createdProject, err := models.CreateProject(project)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusCreated, createdProject)
}

func DeleteProject(c *gin.Context) {
	authToken := utils.GetAuthToken(c)
	projectId, err := uuid.Parse(c.Param("project_id"))
	fmt.Println(projectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid uuid"})
		return
	}
	err = models.DeleteProject(authToken.UID, projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
