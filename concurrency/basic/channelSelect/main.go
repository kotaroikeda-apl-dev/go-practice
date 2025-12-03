package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n=== select文 ===")
	c1 := make(chan string)
	c2 := make(chan string)

	// 2つのgoroutineが異なる速度でメッセージを送信
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// selectで最初に準備できたチャンネルから受信
	// 2つのチャンネルのうち、どちらかが準備できた方を処理する
	// defaultケースがあると、チャンネルが準備できていない時もブロックせずに処理を続けられる
	// 注意: receivedCountを送信されるメッセージ数より大きくすると、永久にループが終わらない
	// （例: 各チャンネルから1つずつしか送信しないのに、receivedCount < 3にすると永久に待機し続ける）
	receivedCount := 0
	for receivedCount < 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received from c1:", msg1)
			receivedCount++
		case msg2 := <-c2:
			fmt.Println("received from c2:", msg2)
			receivedCount++
		default:
			// チャンネルが準備できていない時はブロックせずに待機
			fmt.Println("waiting for messages...")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println("両方のチャンネルから受信が完了しました")
}
