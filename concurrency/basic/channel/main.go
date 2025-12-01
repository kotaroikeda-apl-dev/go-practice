package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 基本的なチャンネル（バッファなし）===")
	c := make(chan string)

	// goroutineでメッセージを送信
	go func() {
		c <- "Hello"
		c <- "World"
		c <- "from"
		c <- "Channel"
	}()

	// メッセージを受信
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("\n=== バッファ付きチャンネル ===")
	ch := make(chan int, 2)

	ch <- 1
	fmt.Println("sent 1")

	ch <- 2
	fmt.Println("sent 2")

	fmt.Println("buffer now full (バッファサイズ2で2つの値が入っている)")

	// 重要：チャンネルの送受信は「同期操作」です
	// バッファが満杯の状態で送信しようとすると、受信側が値を取り出すまでブロックされます
	// goroutineは並列処理ですが、チャンネルの送信操作自体は同期して待ちます

	// goroutineで受信処理を起動（並列で実行される）
	// このgoroutineが <-ch で値を取り出すと、バッファに空きができる
	go func() {
		fmt.Println("goroutine: 受信を開始します...")
		time.Sleep(1 * time.Second) // 1秒待機してから受信（ブロックの様子を明確にするため）
		val := <-ch
		fmt.Println("goroutine: received", val, "(これでバッファに空きができました)")
	}()

	// この時点でバッファは満杯（1, 2が入っている）
	// ch <- 3 を実行すると、バッファに空きができるまでブロックされる
	// 上記のgoroutineが <-ch で値を取り出すまで、ここで待機する
	fmt.Println("3を送信しようとしています...")
	fmt.Println("（バッファが満杯なので、空きができるまで待機中...）")
	ch <- 3 // ここでブロック。goroutineが値を取り出すまで待つ
	fmt.Println("sent 3（ブロックが解除されました）")

	// 残りの値を受信
	fmt.Println("received:", <-ch)
	fmt.Println("received:", <-ch)
}
