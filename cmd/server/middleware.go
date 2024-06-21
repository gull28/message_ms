package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func MethodMiddleware(method string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				w.Header().Set("Allow", method)
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dotEnv := godotenv.Load("../../.env")

		if dotEnv != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		bearerToken := os.Getenv("BEARER_TOKEN")

		if r.Header.Get("Authorization") != bearerToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
