package examples

import (
	"fmt"
	"sync"
)

// Counter は複数のgoroutineから安全にアクセスできるカウンター
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment はカウンターを1増やす（Mutexで保護）
func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Get はカウンターの現在の値を取得する（Mutexで保護）
func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// Mutex sync.Mutexの例
func Mutex() {
	fmt.Println("\n=== sync.Mutex ===")
	counter := &Counter{count: 0}

	// 10個のgoroutineが同時にカウンターを増やす
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // WaitGroupのカウンターを1増やす（新しいgoroutineを登録）
		go func() {
			defer wg.Done() // goroutineが終了する時にWaitGroupのカウンターを1減らす
			// deferを使うことで、エラーが発生しても確実にDone()が呼ばれる
			// 各goroutineが10回カウンターを増やす
			for j := 0; j < 10; j++ {
				counter.Increment()
			}
		}()
	}

	// 全てのgoroutineが完了するまで待機
	// WaitGroupのカウンターが0になるまでブロックする
	// これがないと、main関数がgoroutineの完了を待たずに終了してしまう可能性がある
	wg.Wait()

	// Mutexで保護されているため、正しい値（100）が取得できる
	fmt.Printf("最終的なカウンターの値: %d (期待値: 100)\n", counter.Get())
	fmt.Println("Mutexにより、複数のgoroutineからの同時アクセスが安全に保護されました")
}
