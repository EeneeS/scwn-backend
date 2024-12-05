package models

import (
	"context"
	"fmt"
	"time"

	"github.com/eenees/scwn-backend/src/config"
)

type User struct {
	Id         string    `json:"id"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}

func CreateUser(user *User) (User, error) {
	validUser, err := isValidUser(user)
	if err != nil {
		return *user, err
	}
	if err := config.DB.Create(&validUser).Error; err != nil {
		return validUser, err
	}
	return validUser, nil
}

func isValidUser(user *User) (User, error) {
	firebaseUser, err := config.AuthClient.GetUser(context.Background(), user.Id)
	if err != nil {
		return *user, fmt.Errorf("failed to fetch user by UID: %w", err)
	}
	user.Email = firebaseUser.Email
	return *user, nil
}
