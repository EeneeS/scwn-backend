package projectcontroller

import (
	"github.com/eenees/scwn-backend/src/models"
	"github.com/eenees/scwn-backend/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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

// TODO: check if its the users project else return project not found
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
	authToken := utils.GetAuthToken(c)
	var req struct {
		Name           string                 `json:"name"`
		PublishTargets []models.PublishTarget `json:"publish_targets"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project := models.Project{
		Id:     uuid.New(),
		Name:   req.Name,
		UserId: authToken.UID,
	}
	var publishTargets []models.PublishTarget
	for _, target := range req.PublishTargets {
		publishTargets = append(publishTargets, models.PublishTarget{
			ProjectId: project.Id,
			Platform:  target.Platform,
			Url:       target.Url,
		})
	}
	project.PublishTargets = publishTargets
	createdProject, err := models.CreateProject(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusCreated, createdProject)
}

func DeleteProject(c *gin.Context) {
	authToken := utils.GetAuthToken(c)
	projectId, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	err = models.DeleteProject(authToken.UID, projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
