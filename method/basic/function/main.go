package main

import (
	"fmt"
)

// Counter はカウンターを表す構造体
type Counter struct {
	Value int
}

// SetValue は値を設定しようとする関数（値で受け取るため、元の値は変更されない）
func SetValue(c Counter, value int) {
	c.Value = value
}

// Scale は値を倍にする関数（ポインタで受け取るため、元の値が変更される）
func Scale(c *Counter, factor int) {
	c.Value = c.Value * factor
}

// Increment は値を増やす関数（ポインタで受け取るため、元の値が変更される）
func Increment(c *Counter) {
	c.Value++
}

// String はCounterを文字列として表現するメソッド
// fmtパッケージ（fmt.Printf、fmt.Printlnなど）で%sや%vを使ってCounterを出力する際に自動的に呼び出される
func (c Counter) String() string {
	return fmt.Sprintf("Counter{Value: %d}", c.Value)
}

func main() {
	fmt.Println("=== Pointers and Functions ===")

	// カウンターを作成
	counter := Counter{Value: 5}
	fmt.Printf("初期状態: %s\n", counter)

	// 値で受け取る関数（元の値は変更されない）
	SetValue(counter, 100)
	fmt.Printf("SetValue(counter, 100)呼び出し後（変更なし）: %s\n", counter)

	// ポインタで受け取る関数（元の値が変更される）
	Scale(&counter, 3)
	fmt.Printf("Scale(&counter, 3)呼び出し後（変更あり）: %s\n", counter)

	// ポインタで受け取る関数（元の値が変更される）
	Increment(&counter)
	fmt.Printf("Increment(&counter)呼び出し後（変更あり）: %s\n", counter)
}
