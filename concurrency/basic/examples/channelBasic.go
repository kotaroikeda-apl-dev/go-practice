package examples

import (
	"fmt"
)

// ChannelBasic 基本的なチャンネル（バッファなし）の例
func ChannelBasic() {
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
}
