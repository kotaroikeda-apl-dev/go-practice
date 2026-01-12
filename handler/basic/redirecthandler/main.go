package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.RedirectHandler は、すべてのリクエストを
	// 指定した URL に指定したステータスコードでリダイレクトするハンドラを返します。
	// ここでは、"/new-path" へ 301 (StatusMovedPermanently) でリダイレクトします。
	redirectHandler := http.RedirectHandler("/new-path", http.StatusMovedPermanently)

	// "/old-path" へのアクセスを redirectHandler で処理します。
	http.Handle("/old-path", redirectHandler)

	// リダイレクト先となる "/new-path" のハンドラ
	http.HandleFunc("/new-path", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the new path!")
	})

	// ルートパス "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Redirect handler demo.")
		fmt.Fprintln(w, "Access http://localhost:8080/old-path to see it in action.")
	})

	fmt.Println("Server starting at http://localhost:8080...")
	fmt.Println("Try http://localhost:8080/old-path")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

