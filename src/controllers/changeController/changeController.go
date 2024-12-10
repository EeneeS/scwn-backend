package changecontroller

import (
	"github.com/eenees/scwn-backend/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateChange(c *gin.Context) {
	projectId, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	var req struct {
		Changes []models.Change `json:"changes"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	batchId := uuid.New()
	var changes []models.Change
	for _, change := range req.Changes {
		changes = append(changes, models.Change{
			ProjectId:     projectId,
			Element:       change.Element,
			Type:          change.Type,
			OriginalValue: change.OriginalValue,
			NewValue:      change.NewValue,
			Route:         change.Route,
			BatchId:       batchId,
		})
	}
	createdChanges, err := models.CreateChange(&changes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if available publish to publishTargets
	c.JSON(http.StatusCreated, gin.H{
		"batch_id": batchId,
		"changes":  createdChanges,
	})
}

func GetAllChanges(c *gin.Context) {
	projectId, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	changes, err := models.GetAllChanges(projectId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, changes)
}
