package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// --- Handlers ---

func userDetailHandler(id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Detail for ID: %s (chi)\n", id)
	}
}

func main() {
	r := chi.NewRouter()

	// chiの内蔵ミドルウェアを使用
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			userDetailHandler(id)(w, r)
		})
	})

	fmt.Println("Server (chi) starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

