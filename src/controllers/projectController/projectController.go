package projectcontroller

import (
	"github.com/eenees/scwn-backend/src/models"
	"github.com/eenees/scwn-backend/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetAllProjects(c *gin.Context) {
	authToken, ok := utils.GetAuthToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "could not retrieve token"})
	}
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
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	DBerr := models.DeleteProject(id)
	if DBerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": DBerr.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
