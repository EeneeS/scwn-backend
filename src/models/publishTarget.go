package models

import "github.com/google/uuid"

type PublishTarget struct {
	Id        string    `json:"id"`
	ProjectId uuid.UUID `json:"project_id"`
	Platform  string    `json:"platform"`
	Url       string    `json:"url"`
}
