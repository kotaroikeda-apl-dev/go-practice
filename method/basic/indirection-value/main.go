package main

import (
	"fmt"
	"math"
)

// Point は2次元座標を表す構造体
type Point struct {
	X, Y float64
}

// LengthByMethod は値レシーバーを持つメソッド
// メソッドは値でもポインタでも呼び出せる（Goが自動変換）
func (p Point) LengthByMethod() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// LengthByFunc は値引数を持つ関数
// 関数は値を受け取る必要がある
func LengthByFunc(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	fmt.Println("=== メソッドとポインタの間接参照（値レシーバー編） ===")
	fmt.Println()

	// ケース1: 値として定義したPointから呼び出す
	fmt.Println("【ケース1: 値から呼び出す】")
	v := Point{X: 3, Y: 4}
	fmt.Printf("初期値: %+v\n", v)

	// ✅ メソッド（値レシーバー）は値から呼び出せる
	length1 := v.LengthByMethod()
	fmt.Printf("v.LengthByMethod() = %.2f\n", length1)

	// ✅ 関数（値引数）も値を受け取れる
	length2 := LengthByFunc(v)
	fmt.Printf("LengthByFunc(v) = %.2f\n", length2)
	fmt.Println()

	// ケース2: ポインタとして定義したPointから呼び出す
	fmt.Println("【ケース2: ポインタから呼び出す】")
	p := &Point{X: 4, Y: 3}
	fmt.Printf("初期値: %+v\n", p)

	// ✅ メソッド（値レシーバー）はポインタからも呼び出せる（Goが自動変換）
	length3 := p.LengthByMethod()
	fmt.Printf("p.LengthByMethod() = %.2f\n", length3)

	// ❌ 関数（値引数）はポインタから直接呼び出せない（コンパイルエラー）
	// LengthByFunc(p) // エラー: cannot use p (type *Point) as type Point

	// ✅ 関数は値を明示的に渡す必要がある（ポインタをデリファレンス）
	length4 := LengthByFunc(*p)
	fmt.Printf("LengthByFunc(*p) = %.2f\n", length4)
}
