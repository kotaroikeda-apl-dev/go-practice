package main

import (
	"fmt"
	"net/http"
)

// helloHandler は普通の関数です。
// 引数に (http.ResponseWriter, *http.Request) を取ります。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは! あなたがアクセスしたパスは: %s です。（HandlerFuncを使用）", r.URL.Path)
}

func main() {
	// http.HandlerFunc はアダプターです。
	// これを使うことで、特定のシグネチャを持つ関数を http.Handler インターフェースに適合させることができます。
	handler := http.HandlerFunc(helloHandler)

	// "/"（すべてのパス）に対して、作成したハンドラを登録します
	http.Handle("/", handler)

	fmt.Println("サーバーを起動しています... http://localhost:8080 にアクセスしてください。")

	// http.ListenAndServe はサーバーを起動し、リクエストを待ち受けます
	// 第2引数が nil の場合、デフォルトのマルチプレクサ（DefaultServeMux）が使用されます
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("サーバーの起動に失敗しました: %s\n", err)
	}
}
