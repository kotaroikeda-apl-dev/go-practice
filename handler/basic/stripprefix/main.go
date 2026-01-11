package main

import (
	"fmt"
	"net/http"
)

func main() {
	// メインのハンドラを定義します
	// このハンドラは、リクエストされたパスを表示するだけのシンプルなものです
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "最終的なハンドラが受け取ったパス: %s\n", r.URL.Path)
	})

	// http.StripPrefix を使用して、パスから "/tmp" を取り除きます
	// "/tmp/hello" というリクエストが来た場合、finalHandler には "/hello" として渡されます
	// "/tmp" で始まらないリクエスト（例: "/other"）には 404 を返します
	strippedHandler := http.StripPrefix("/tmp", finalHandler)

	// "/tmp/" で始まるすべてのリクエストを strippedHandler に送ります
	http.Handle("/tmp/", strippedHandler)

	fmt.Println("サーバーを起動しています... http://localhost:8080/tmp/test にアクセスしてください。")
	fmt.Println("'/tmp' で始まらないパス（例: http://localhost:8080/other）は 404 になります。")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("サーバーの起動に失敗しました: %s\n", err)
	}
}


