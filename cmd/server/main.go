package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
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
		// log dotenv variable
		errorLog.Println(dotEnv)
		return
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := openDB(dsn)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
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

func openDB(dsn string) (*sql.DB, error) {
	db = GetDB()
}
