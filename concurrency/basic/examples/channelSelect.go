package examples

import (
	"fmt"
	"time"
)

// ChannelSelect select文の例
func ChannelSelect() {
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
	// 2回ループするので、両方のチャンネルから受信したら終了する
	// 注意: ループ回数を送信されるメッセージ数より多くすると、デッドロックが発生する
	// （例: 各チャンネルから1つずつしか送信しないのに、3回ループすると3回目で待機し続ける）
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received from c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("received from c2:", msg2)
		}
	}
	fmt.Println("両方のチャンネルから受信が完了しました")
}
