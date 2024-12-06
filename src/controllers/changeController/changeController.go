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
	var changes []models.Change
	for _, change := range req.Changes {
		changes = append(changes, models.Change{
			ProjectId:     projectId,
			Element:       change.Element,
			Type:          change.Type,
			OriginalValue: change.OriginalValue,
			NewValue:      change.NewValue,
		})
	}
	createdChanges, err := models.CreateChange(&changes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdChanges)
}
