package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// --- Middleware ---

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// --- Handlers ---

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func userDetailHandler(id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Detail for ID: %s (Standard Library)\n", id)
	}
}

// --- Main ---

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	// {id} パラメータの例
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		userDetailHandler(id)(w, r)
	})

	fmt.Println("Server (std) starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", logger(mux)))
}

