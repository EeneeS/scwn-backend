package models

import "github.com/google/uuid"

type Project struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name"`
}

func GetAllProjects() []Project {
	return []Project{}
}

func GetProject(id uuid.UUID) uuid.UUID {
	return id
}

func CreateProject(project Project) {

}
