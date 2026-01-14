package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// リクエストボディを読み込んでサイズを表示するハンドラー
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			// MaxBytesHandler によって制限を超えた場合、ここでエラーが発生する
			http.Error(w, fmt.Sprintf("Body read error: %v", err), http.StatusRequestEntityTooLarge)
			return
		}
		fmt.Fprintf(w, "Received body of size: %d bytes\n", len(body))
		fmt.Fprintf(w, "Body content: %s\n", string(body))
	})

	// MaxBytesHandler を使用して、リクエストボディのサイズを 1024 バイト（1KB）に制限する
	// これにより、大きなリクエストによるリソースの枯渇を防ぐことができる
	limit := int64(1024)
	restrictedHandler := http.MaxBytesHandler(handler, limit)

	fmt.Println("Starting server on :8080...")
	fmt.Println("Try sending a small request:")
	fmt.Printf("  curl -X POST -d \"hello\" http://localhost:8080\n")
	fmt.Println("Try sending a large request (> 1KB):")
	fmt.Printf("  head -c 2000 /dev/zero | curl -X POST --data-binary @- http://localhost:8080\n")

	if err := http.ListenAndServe(":8080", restrictedHandler); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
