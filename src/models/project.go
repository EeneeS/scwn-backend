package models

import (
	"errors"
	"fmt"
	"github.com/eenees/scwn-backend/src/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	Id             uuid.UUID       `json:"id,omitempty"`
	Name           string          `json:"name"`
	UserId         string          `json:"user_id"`
	PublishTargets []PublishTarget `json:"publish_targets"`
}

func GetAllProjects(uid string) ([]Project, error) {
	var projects []Project
	result := config.DB.
		Preload("PublishTargets").
		Where("user_id = ?", uid).
		Find(&projects)
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

func CreateProject(project *Project) (Project, error) {
	if project.PublishTargets == nil {
		project.PublishTargets = []PublishTarget{}
	}
	newProject := project
	result := config.DB.Create(&newProject)
	if result.Error != nil {
		return *newProject, result.Error
	}
	return *newProject, nil
}

func DeleteProject(userId string, projectId uuid.UUID) error {
	if projectId == uuid.Nil {
		return fmt.Errorf("Invalid project ID.")
	}
	result := config.DB.Where("user_id = ? and id = ?", userId, projectId).Delete(&Project{}, projectId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Project not found.")
	}
	return nil
}
