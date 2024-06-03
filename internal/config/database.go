package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func (app *application) initDB() *gorm.DB {

	dotEnv := godotenv.Load("")

	if dotEnv != nil {
		fmt.Println("Error loading .env file")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Open a connection to the database
	dbUrl := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName)
	db, err := gorm.Open("mysql", dbUrl)

	if err != nil {
		fmt.Println("Failed to connect to the database!")
		panic(err)
	}

	DB = db

	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	DB.Close()
}
