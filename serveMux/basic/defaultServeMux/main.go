package main

import (
	"fmt"
	"net/http"
)

// DefaultServeMux は、Serve によって使用されるデフォルトの ServeMux です。
// パッケージレベルの関数 http.Handle および http.HandleFunc は、
// 内部で http.DefaultServeMux を使用しています。

func main() {
	// http.HandleFunc は内部で http.DefaultServeMux.HandleFunc(pattern, handler) を呼び出します
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "DefaultServeMux を使用したレスポンスです")
	})

	fmt.Println("Server starting on :8080 using DefaultServeMux...")

	// http.ListenAndServe の第2引数に nil を渡すと、
	// デフォルトで http.DefaultServeMux がハンドラとして使用されます。
	http.ListenAndServe(":8080", nil)
}
