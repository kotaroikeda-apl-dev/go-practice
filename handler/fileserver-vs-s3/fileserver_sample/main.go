package main

import (
	"log"
	"net/http"
)

// http.FileServer を使った標準的な実装例
// ./static ディレクトリ配下のファイルを配信します。
func main() {
	// ファイルサーバーのハンドラーを作成
	fs := http.FileServer(http.Dir("./static"))

	// "/static/" へのリクエストを処理し、パスのプレフィックスを削ってファイルを探す
	// 例: /static/index.html -> ./static/index.html
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Serving files on http://localhost:8080/static/...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
