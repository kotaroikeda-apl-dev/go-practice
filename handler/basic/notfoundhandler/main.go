package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.NotFoundHandler() は、すべてのリクエストに対して
	// "404 page not found" というレスポンスを返すハンドラを返します。
	handler := http.NotFoundHandler()

	// 特定のパス（例: /missing）に対して NotFoundHandler を登録します。
	// これにより、そのパスへのアクセスは常に 404 エラーになります。
	http.Handle("/missing", handler)

	// ルートパス "/" には通常のメッセージを返すハンドラを登録します。
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome! Try accessing /missing to see the NotFoundHandler in action.")
	})

	fmt.Println("Server starting at http://localhost:8080...")
	fmt.Println("Check http://localhost:8080/missing for 404 response.")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
