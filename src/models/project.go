package models

import (
	"github.com/google/uuid"
)

type Project struct {
	Id     uuid.UUID `json:"id,omitempty"`
	Name   string    `json:"name"`
	UserId uuid.UUID `json:"user_id"`
}

func GetAllProjects() []Project {
	var projects []Project
	DB.Find(&projects)
	return projects
}

func GetProject(id uuid.UUID) uuid.UUID {
	return id
}

func CreateProject(project Project) {

}
