package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ServeMux の作成
	mux := http.NewServeMux()

	// 1. 静的なパスの登録
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// 2. パスパラメータ（ワイルドカード）を含むパスの登録 (Go 1.22+)
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "User ID: %s\n", id)
	})

	// 3. プレフィックス一致の登録
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API Root")
	})

	// サーバーをラップして、リクエストごとに mux.Handler を呼び出して
	// どのハンドラがマッチするかをログ出力するデモ
	loggingHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// このメソッドは、リクエストにマッチするハンドラとその登録パターンを返します。
		// 実際にハンドラを実行する前に、どの定義にマッチしたかを知ることができます。
		h, pattern := mux.Handler(r)

		fmt.Printf("\n[Request Log]\n")
		fmt.Printf("  Method: %s, Path: %s\n", r.Method, r.URL.Path)
		fmt.Printf("  Matched Pattern: %q\n", pattern)
		fmt.Printf("  Handler Type   : %T\n", h)

		// 注意: 画像の説明にある通り、mux.Handler はマッチしたハンドラを特定するだけで、
		// r.PathValue などのパラメータをセットすることはありません。
		// そのため、この時点では id は空です。
		fmt.Printf("  PathValue 'id' (before ServeHTTP): %q\n", r.PathValue("id"))

		// 実際に mux にリクエストを処理させます。
		// ここで内部的にマッチングが再度行われ、r.PathValue がセットされた上で
		// 登録されたハンドラが実行されます。
		mux.ServeHTTP(w, r)
	})

	fmt.Println("Server starting at http://localhost:8080...")
	fmt.Println("Try these URLs and check the terminal output:")
	fmt.Println("  1. http://localhost:8080/hello")
	fmt.Println("  2. http://localhost:8080/users/42")
	fmt.Println("  3. http://localhost:8080/api/test")
	fmt.Println("  4. http://localhost:8080/unknown")

	err := http.ListenAndServe(":8080", loggingHandler)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
