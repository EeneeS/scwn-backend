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

type Changes struct {
	Changes []Change `json:"changes"`
}

// TODO: left here
func CreateChange(change *Change) (Change, error) {
	if err := config.DB.Create(&change).Error; err != nil {
		return *change, err
	}
	return *change, nil
}
