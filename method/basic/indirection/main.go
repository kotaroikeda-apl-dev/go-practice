package main

import (
	"fmt"
)

// Point は2次元座標を表す構造体
type Point struct {
	X, Y float64
}

// MoveByMethod はポインタレシーバーを持つメソッド
// メソッドは値でもポインタでも呼び出せる（Goが自動変換）
func (p *Point) MoveByMethod(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// MoveByFunc はポインタ引数を持つ関数
// 関数はポインタを明示的に渡す必要がある
func MoveByFunc(p *Point, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
	fmt.Println("=== メソッドとポインタの間接参照 ===")
	fmt.Println()

	// ケース1: 値として定義したPointから呼び出す
	fmt.Println("【ケース1: 値から呼び出す】")
	v := Point{X: 3, Y: 4}
	fmt.Printf("初期値: %+v\n", v)

	// ✅ メソッド（ポインタレシーバー）は値から呼び出せる（Goが自動変換）
	v.MoveByMethod(2, 1)
	fmt.Printf("v.MoveByMethod(2, 1)後: %+v\n", v)

	// ❌ 関数（ポインタ引数）は値から直接呼び出せない（コンパイルエラー）
	// MoveByFunc(v, 5, 3) // エラー: cannot use v (type Point) as type *Point

	// ✅ 関数はポインタを明示的に渡す必要がある
	MoveByFunc(&v, 5, 3)
	fmt.Printf("MoveByFunc(&v, 5, 3)後: %+v\n", v)
	fmt.Println()

	// ケース2: ポインタとして定義したPointから呼び出す
	fmt.Println("【ケース2: ポインタから呼び出す】")
	p := &Point{X: 4, Y: 3}
	fmt.Printf("初期値: %+v\n", p)

	// ✅ メソッド（ポインタレシーバー）はポインタからも呼び出せる
	p.MoveByMethod(-1, 2)
	fmt.Printf("p.MoveByMethod(-1, 2)後: %+v\n", p)

	// ✅ 関数（ポインタ引数）もポインタを渡せる
	MoveByFunc(p, 3, -1)
	fmt.Printf("MoveByFunc(p, 3, -1)後: %+v\n", p)
}
