package publishtargetcontroller

import (
	"github.com/eenees/scwn-backend/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// TODO: check if the platform has already been added else dont create it
func CreatePublishTarget(c *gin.Context) {
	projectId, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	var req struct {
		PublishTargets []models.PublishTarget `json:"publish_targets"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var publishTargets []models.PublishTarget
	for _, target := range req.PublishTargets {
		publishTargets = append(publishTargets, models.PublishTarget{
			ProjectId: projectId,
			Platform:  target.Platform,
			Url:       target.Url,
		})
	}
	createdPublishTargets, err := models.CreatePublishTarget(&publishTargets)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdPublishTargets)
}
