package models

import (
	"github.com/eenees/scwn-backend/src/config"
	"github.com/google/uuid"
)

type Change struct {
	ProjectId     uuid.UUID `json:"project_id"`
	Element       string    `json:"element"`
	Type          string    `json:"type"`
	OriginalValue string    `json:"orignal_value"`
	NewValue      string    `json:"new_value"`
}

func CreateChange(changes *[]Change) ([]Change, error) {
	newChanges := changes
	result := config.DB.Create(&newChanges)
	if result.Error != nil {
		return *newChanges, result.Error
	}
	return *newChanges, nil
}
