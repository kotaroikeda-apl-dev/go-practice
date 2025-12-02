package examples

import (
	"fmt"
	"time"
)

func sendNumbers(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c) // チャンネルを閉じる（これによりrangeループが終了する）
}

// ChannelRangeClose range と close の例
func ChannelRangeClose() {
	fmt.Println("\n=== range と close ===")
	numChan := make(chan int, 5)
	go sendNumbers(10, numChan)
	// rangeでチャンネルから値を受信（チャンネルが閉じられるまで続く）
	// close()が呼ばれると、rangeループが自動的に終了する
	for num := range numChan {
		fmt.Println("received:", num+1)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("チャンネルが閉じられました")
}
