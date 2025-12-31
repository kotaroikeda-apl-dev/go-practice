package main

import "fmt"

// Speaker はメソッドだけが宣言されたインターフェース
// 実装は紐づかず、ゼロ値は nil（型も値も空）
type Speaker interface {
	Say()
}

// describe はインターフェース値の中身を表示する
func describe(s Speaker) {
	fmt.Printf("(%v, %T)\n", s, s)
}

func main() {
	var s Speaker // ゼロ値は nil（型も値も空）
	describe(s)

	// nil インターフェースにメソッド呼び出しすると実行時エラーになる
	// 実際に確認したいときは下をコメントアウト解除
	// s.Say()
}
