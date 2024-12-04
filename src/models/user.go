package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}

func CreateUser(user User) {
	// check if user exist in fire base if so make it on the database otherwise dont do it
}
