package models

import (
	"github.com/eenees/scwn-backend/src/config"
	"github.com/google/uuid"
)

type Change struct {
	ProjectId     uuid.UUID `json:"project_id"`
	Element       string    `json:"element"`
	Type          string    `json:"type"`
	OriginalValue string    `json:"original_value"`
	NewValue      string    `json:"new_value"`
	Route         string    `json:"route"`
	BatchId       uuid.UUID `json:"batch_id"`
}

func GetAllChanges(id uuid.UUID) ([]Change, error) {
	var changes []Change
	result := config.DB.Where("project_id = ?", id).Find(&changes)
	if result.Error != nil {
		return changes, result.Error
	}
	return changes, nil
}

func CreateChange(changes *[]Change) ([]Change, error) {
	newChanges := changes
	result := config.DB.Create(&newChanges)
	if result.Error != nil {
		return *newChanges, result.Error
	}
	return *newChanges, nil
}
