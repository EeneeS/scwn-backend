package models

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var FB *firebase.App

func LoadEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDatabase() {
	sqluser := os.Getenv("MYSQL_USER")
	sqlPassword := os.Getenv("MYSQL_PASSWORD")
	sqlDB := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", sqluser, sqlPassword, sqlDB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database.")
	}

	DB = db
}

func ConnectFirebase() {

	credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentialsFile == "" {
		log.Fatal("GOOGLE_APPLICATION_CREDENTIALS is not set")
	}

	fb, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(credentialsFile))

	if err != nil {
		log.Fatalf("error initializing firebase: %v\n", err)
	}

	// client, err := app.Auth(context.Background())
	// if err != nil {
	// 	log.Fatalf("error getting Auth client: %v\n", err)
	// }
	//
	// user, err := client.GetUser(context.Background(), "yebcLJvxT3akhZeT98w3Gfu3qCp1")
	// if err != nil {
	// 	log.Fatalf("error fetching user by UID: %v\n", err)
	// }
	//
	// fmt.Println(user)

	FB = fb
}
