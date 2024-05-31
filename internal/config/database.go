package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func (app *application) main() {

	dotEnv := godotenv.Load("")

	if dotEnv != nil {
		fmt.Println("Error loading .env file")
		return
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Open a connection to the database
	dbUrl := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", username, password, dbName)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the database!")
}
