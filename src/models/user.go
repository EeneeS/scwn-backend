package models

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}

func CreateUser(user User) (User, error) {
	newUser, err := isValidUser(user)

	if err != nil {
		return newUser, err
	}

	result := DB.Create(&newUser)

	if result.Error != nil {
		return newUser, result.Error
	}

	return newUser, nil
}

func isValidUser(user User) (User, error) {
	client, err := FB.Auth(context.Background())

	if err != nil {
		return user, fmt.Errorf("error getting Auth client: %v\n", err)
	}

	_, userError := client.GetUser(context.Background(), user.Id)

	if userError != nil {
		return user, fmt.Errorf("error fetch user by UID: %v\n", userError)
	}

	return user, nil
}
