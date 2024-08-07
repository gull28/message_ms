package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gull28/message_ms/internal/config"
	"github.com/gull28/message_ms/internal/db"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *gorm.DB
	config   config.Config
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye, World!")
}

func main() {
	port := flag.String("port", ":8080", "HTTP Port to run the server on")
	flag.Parse()

	dotEnv := godotenv.Load("../../.env")

	// simple logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if dotEnv != nil {
		errorLog.Println("Error loading .env file")
		errorLog.Println(dotEnv)
		return
	}

	db, err := openDB()

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		db:       db,
		config:   config.LoadConfig(),
	}

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// http.HandleFunc("/messages", getMessages)

	infoLog.Printf("Starting server on %s", *port)
	err = srv.ListenAndServe()

	errorLog.Fatal(err)
}

func openDB() (*gorm.DB, error) {
	database := db.InitDB()

	return database, nil
}
