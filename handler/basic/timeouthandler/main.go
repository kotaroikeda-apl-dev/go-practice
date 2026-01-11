package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 非常に時間のかかる処理をシミュレートするハンドラー
	slowHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Fprintln(w, "処理が完了しました")
	})

	// TimeoutHandler を使用して、指定した時間（1秒）以内にレスポンスが返らない場合に
	// 503 Service Unavailable を返すようにラップする
	timeoutHandler := http.TimeoutHandler(slowHandler, 1*time.Second, "タイムアウトしました（サーバーが混雑しています）")

	fmt.Println("Starting server on :8080...")
	fmt.Println("Visit http://localhost:8080 to see the timeout in action.")

	if err := http.ListenAndServe(":8080", timeoutHandler); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
