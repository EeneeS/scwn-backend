package models

import (
	"github.com/eenees/scwn-backend/src/config"
	"github.com/google/uuid"
)

type PublishTarget struct {
	ProjectId uuid.UUID `json:"project_id"`
	Platform  string    `json:"platform"`
	Url       string    `json:"url"`
}

func CreatePublishTarget(publishTargets *[]PublishTarget) ([]PublishTarget, error) {
	newPublishTargets := publishTargets
	result := config.DB.Create(&newPublishTargets)
	if result.Error != nil {
		return *newPublishTargets, result.Error
	}
	return *newPublishTargets, nil
}
