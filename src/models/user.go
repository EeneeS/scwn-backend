package models

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/auth"
	"github.com/eenees/scwn-backend/src/config"
)

type User struct {
	Id         string    `json:"id"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}

func CreateUser(authToken *auth.Token, user *User) (User, error) {
	validUser, err := isValidUser(authToken, user)
	if err != nil {
		return *user, err
	}
	if err := config.DB.Create(&validUser).Error; err != nil {
		return validUser, err
	}
	return validUser, nil
}

// TEST: this function is no longer neccesary because we are getting an auth token that is already being checked
// but i dont think its bad to have this function still in place (double check)
func isValidUser(authToken *auth.Token, user *User) (User, error) {
	firebaseUser, err := config.AuthClient.GetUser(context.Background(), authToken.UID)
	if err != nil {
		return *user, fmt.Errorf("failed to fetch user by UID: %w", err)
	}
	user.Email = firebaseUser.Email
	return *user, nil
}
