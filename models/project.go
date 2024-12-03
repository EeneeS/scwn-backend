package models

import "github.com/google/uuid"

type Project struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name"`
}
