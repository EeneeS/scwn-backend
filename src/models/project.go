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

func GetAllProjects() ([]Project, error) {
	var projects []Project
	result := DB.Find(&projects)
	if result.Error != nil {
		return projects, result.Error
	}
	return projects, nil
}

func GetProject(id uuid.UUID) (Project, error) {
	var project Project
	result := DB.First(&project, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return project, errors.New("Project not found")
		}
		return project, result.Error
	}
	return project, nil
}

func CreateProject(project Project) (Project, error) {
	newProject := project
	result := DB.Create(&newProject)
	if result.Error != nil {
		return newProject, result.Error
	}
	return newProject, nil
}

func DeleteProject(id uuid.UUID) error {
	result := DB.Delete(&Project{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
