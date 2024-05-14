package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	http.HandleFunc("/hello", handler)
	flag.Parse()

	// simple logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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
	err := srv.ListenAndServe()

	errorLog.Fatal(err)

}
