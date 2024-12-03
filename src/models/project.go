package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func GetProject(id uuid.UUID) (Project, error) {
	var project Project
	result := DB.First(&project, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return project, errors.New("Record not found")
		}
		return project, result.Error
	}
	return project, nil
}

func CreateProject(project Project) {

}
