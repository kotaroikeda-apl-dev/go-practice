package main

import (
	"fmt"
	"net/http"
)

// apiHandler は http.Handler インターフェースを実装するカスタム構造体です。
type apiHandler struct {
	message string
}

// ServeHTTP は http.Handler インターフェースのメソッドです。
func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Response: %s\n", h.message)
}

func main() {
	mux := http.NewServeMux()

	// 1. カスタムハンドラーの登録
	// Handle は第2引数に http.Handler インターフェースを期待します。
	mux.Handle("/api/v1", &apiHandler{message: "Version 1.0"})

	// 2. http.HandlerFunc を使用した登録
	// HandleFunc ではなく Handle を使って、関数をハンドラーとして登録することもできます。
	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	}))

	// 3. パニックの例 (コメントアウト)
	// 同じパターンを再度登録しようとすると、Handle はパニックを起こします。
	// mux.Handle("/api/v1", &apiHandler{message: "Conflict!"})

	fmt.Println("Server starting on :8080...")
	fmt.Println("Endpoints:")
	fmt.Println("  - http://localhost:8080/api/v1")
	fmt.Println("  - http://localhost:8080/hello")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
