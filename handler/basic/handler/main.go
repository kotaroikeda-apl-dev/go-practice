package main

import (
	"fmt"
	"net/http"
)

// MyHandler は http.Handler インターフェースを実装するための構造体です。
type MyHandler struct {
	Greeting string
}

// ServeHTTP メソッドを定義することで、MyHandler は http.Handler インターフェースを満たします。
// 画像の説明にあったように、ResponseWriter にデータを書き込みます。
func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s! あなたがアクセスしたパスは: %s です。", h.Greeting, r.URL.Path)
}

func main() {
	// 1. ハンドラのインスタンスを作成します
	handler := MyHandler{Greeting: "こんにちは"}

	// 2. サーバーの設定をします
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
