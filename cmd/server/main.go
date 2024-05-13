package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye, World!")
}

func main() {
	http.HandleFunc("/hello", handler)

	// http.HandleFunc("/messages", getMessages)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
