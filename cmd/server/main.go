package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	http.HandleFunc("/hello", handler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
