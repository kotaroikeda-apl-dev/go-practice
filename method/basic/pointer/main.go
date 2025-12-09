package main

import (
	"fmt"
)

// Counter はカウンターを表す構造体
type Counter struct {
	Value int
}

// SetValue は値を設定しようとするメソッド（値レシーバーでは元の値は変更されない）
func (c Counter) SetValue(value int) {
	c.Value = value
}

// Multiply は値を倍にするメソッド（ポインタレシーバーで状態を更新する）
func (c *Counter) Multiply(factor int) {
	c.Value = c.Value * factor
}

// Increment は値を増やすメソッド（ポインタレシーバーで状態を更新する）
func (c *Counter) Increment() {
	c.Value++
}

// String はCounterを文字列として表現するメソッド
// fmtパッケージ（fmt.Printf、fmt.Printlnなど）で%sや%vを使ってCounterを出力する際に自動的に呼び出される
func (c Counter) String() string {
	return fmt.Sprintf("Counter{Value: %d}", c.Value)
}

func main() {
	fmt.Println("=== Go Methods with Pointer Receivers ===")

	// カウンターを作成
	counter := Counter{Value: 5}
	fmt.Printf("初期状態: %s\n", counter)

	// 値レシーバーで値を設定しようとしても元の値は変更されない
	counter.SetValue(100)
	fmt.Printf("SetValue(100)呼び出し後（変更なし）: %s\n", counter)

	// ポインタレシーバーで値を倍にする（元の値が変更される）
	counter.Multiply(3)
	fmt.Printf("3倍後: %s\n", counter)

	// ポインタレシーバーで値を増やす（元の値が変更される）
	counter.Increment()
	fmt.Printf("1増やした後: %s\n", counter)
}
