package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// 1. パスのみのマッチング
	// "/index.html" は、任意のホストとメソッドに対してパス "/index.html" にマッチします。
	mux.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Path: /index.html")
	})

	// 2. メソッドとパスのマッチング
	// "GET /static/" は、パスが "/static/" で始まる GET リクエストにマッチします。
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Method: GET, Path Prefix: /static/")
	})

	// 3. ホストのマッチング
	// "example.com/" は、ホスト "example.com" への任意のリクエストにマッチします。
	mux.HandleFunc("example.com/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Host: example.com")
	})

	// 4. ホストと完全一致パス
	// "example.com/{$}" は、ホスト "example.com" かつパスが "/" のリクエストにマッチします。
	mux.HandleFunc("example.com/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Host: example.com, Exact Path: /")
	})

	// 5. ワイルドカードのマッチング
	// "/b/{bucket}/o/{objectname...}"
	// {bucket} は1つのセグメント、{objectname...} は残りのパス全体にマッチします。
	mux.HandleFunc("/b/{bucket}/o/{objectname...}", func(w http.ResponseWriter, r *http.Request) {
		bucket := r.PathValue("bucket")
		objectname := r.PathValue("objectname")
		fmt.Fprintf(w, "Bucket: %s, ObjectName: %s\n", bucket, objectname)
	})

	// 6. 完全一致のマッチング ({$})
	// "/{$}" はパス "/" のみにマッチします。 "/" はすべてのパスにマッチします。
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Exact Path: /")
	})

	// 7. 優先順位 (Precedence)
	// より具体的なパターンが優先されます。
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Path Prefix: /images/")
	})
	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Path Prefix: /images/thumbnails/ (More specific)")
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", mux)
}
