package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.Dir は、指定したディレクトリを http.FileSystem として扱います。
	// ここでは "./static" ディレクトリを指定しています。
	fs := http.Dir("./static")

	// http.FileServer は、指定したファイルシステムの内容を配信する http.Handler を返します。
	fileServer := http.FileServer(fs)

	// ルートパス "/" へのリクエストを FileServer に渡します。
	// ブラウザで http://localhost:8080 にアクセスすると、./static/index.html が表示されます。
	http.Handle("/", fileServer)

	fmt.Println("サーバーを起動しています... http://localhost:8080 にアクセスしてください。")
	fmt.Println("静的ファイル (./static/index.html) が配信されます。")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("サーバーの起動に失敗しました: %s\n", err)
	}
}
