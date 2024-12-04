package models

import (
	"context"
	"fmt"
	"time"
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
	if err := DB.Create(&validUser).Error; err != nil {
		return validUser, err
	}
	return validUser, nil
}

func isValidUser(user *User) (User, error) {
	client, err := FB.Auth(context.Background())
	if err != nil {
		return *user, fmt.Errorf("failed to get Auth client: %w", err)
	}
	firebaseUser, err := client.GetUser(context.Background(), user.Id)
	if err != nil {
		return *user, fmt.Errorf("failed to fetch user by UID: %w", err)
	}
	user.Email = firebaseUser.Email
	return *user, nil
}
