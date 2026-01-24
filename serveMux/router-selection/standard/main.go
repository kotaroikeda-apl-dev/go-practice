package main

import (
	"fmt"
	"log"
	"net/http"
	"std-example/middleware"
)

// --- Handlers ---

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

// UserHandler は、ルーターに依存しない形式で実装されたハンドラの例です
func UserHandler(id string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Detail for ID: %s (Standard Library)\n", id)
	})
}

// --- Main ---

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	// {id} パラメータの例
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		UserHandler(id).ServeHTTP(w, r)
	})

	fmt.Println("Server (std) starting on :8080...")
	// ミドルウェアを適用して起動（外部パッケージから呼び出し）
	h := middleware.Logging(mux)
	h = middleware.Recover(h)
	log.Fatal(http.ListenAndServe(":8080", h))
}
