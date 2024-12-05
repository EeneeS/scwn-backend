package models

import (
	"errors"
	"fmt"
	"github.com/eenees/scwn-backend/src/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	Id     uuid.UUID `json:"id,omitempty"`
	Name   string    `json:"name"`
	UserId string    `json:"user_id"`
}

func GetAllProjects() ([]Project, error) {
	var projects []Project
	result := config.DB.Find(&projects)
	if result.Error != nil {
		return projects, result.Error
	}
	return projects, nil
}

func GetProject(id uuid.UUID) (Project, error) {
	var project Project
	result := config.DB.First(&project, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return project, errors.New("Project not found.")
		}
		return project, result.Error
	}
	return project, nil
}

func GetProjectByUser(user_id string) ([]Project, error) {
	var projects []Project
	result := config.DB.Where("user_id = ?", user_id).Find(&projects)
	if result.Error != nil {
		return projects, result.Error
	}
	return projects, nil
}

func CreateProject(project Project) (Project, error) {
	newProject := project
	result := config.DB.Create(&newProject)
	if result.Error != nil {
		return newProject, result.Error
	}
	return newProject, nil
}

func DeleteProject(id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("Invalid project ID.")
	}
	result := config.DB.Delete(&Project{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Project not found.")
	}
	return nil
}
