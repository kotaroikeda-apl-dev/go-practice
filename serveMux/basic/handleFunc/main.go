package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// HandleFunc は、指定されたパターンのハンドラー関数を登録します。
	// func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, HandleFunc!")
	})

	// 同じパターンを再度登録しようとすると、HandleFunc はパニックを起こします。
	// 画像のドキュメントに記載されている通り、パターンが衝突するとパニックになります。
	// 次の行のコメントを外すと、起動時にパニックが発生します。
	// mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "This will cause a panic!")
	// })

	fmt.Println("Server starting on :8080...")
	fmt.Println("Check: http://localhost:8080/hello")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
