package main

import (
	"fmt"
	"math"
)

// Point は2次元座標を表す構造体
type Point struct {
	X, Y float64
}

// LargeStruct は大きな構造体の例（効率性の観点を示すため）
type LargeStruct struct {
	Data [1000]int
	Name string
}

// Scale はポインタレシーバーを持つメソッド（値を変更する）
func (p *Point) Scale(factor float64) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}

// Length はポインタレシーバーを持つメソッド（値を変更しないが、一貫性のためポインタレシーバーを使用）
func (p *Point) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// UpdateData はポインタレシーバーを持つメソッド（大きな構造体のコピーを避けるため）
func (ls *LargeStruct) UpdateData(index int, value int) {
	if index >= 0 && index < len(ls.Data) {
		ls.Data[index] = value
	}
}

// GetName はポインタレシーバーを持つメソッド（一貫性のため、同じ型の他のメソッドと統一）
func (ls *LargeStruct) GetName() string {
	return ls.Name
}

// SetName はポインタレシーバーを持つメソッド（値を変更する）
func (ls *LargeStruct) SetName(name string) {
	ls.Name = name
}

func main() {
	fmt.Println("=== 値レシーバーとポインタレシーバーの選択 ===")
	fmt.Println()

	// 理由1: 値を変更する必要がある場合
	fmt.Println("【理由1: 値を変更する必要がある場合】")
	point := Point{X: 3, Y: 4}
	fmt.Printf("初期値: %+v\n", point)
	point.Scale(2)
	fmt.Printf("Scale(2)後: %+v\n", point)
	fmt.Println()

	// 理由2: 効率性のため（大きな構造体のコピーを避ける）
	fmt.Println("【理由2: 効率性のため（大きな構造体のコピーを避ける）】")
	large := &LargeStruct{Name: "Test"}
	large.UpdateData(0, 100)
	large.UpdateData(999, 200)
	fmt.Printf("Data[0] = %d, Data[999] = %d\n", large.Data[0], large.Data[999])
	fmt.Println("大きな構造体の場合、値レシーバーだと毎回コピーが発生するため非効率")
	fmt.Println("ポインタレシーバーを使うことで、コピーを避けられる")
	fmt.Println()

	// 一貫性の重要性
	fmt.Println("【一貫性の重要性】")
	fmt.Println("同じ型のメソッドは、全て値レシーバーかポインタレシーバーの")
	fmt.Println("どちらかに統一することが推奨されます。")
	fmt.Println()
	fmt.Println("Point型の例:")
	fmt.Println("- Scale() は値を変更するためポインタレシーバー")
	fmt.Println("- Length() は値を変更しないが、一貫性のためポインタレシーバー")
	fmt.Println()
	fmt.Println("LargeStruct型の例:")
	fmt.Println("- UpdateData() は値を変更するためポインタレシーバー")
	fmt.Println("- GetName() は値を変更しないが、一貫性のためポインタレシーバー")
	fmt.Println("- SetName() は値を変更するためポインタレシーバー")
	fmt.Println()
	fmt.Printf("LargeStructのName: %s\n", large.GetName())
	large.SetName("Updated")
	fmt.Printf("SetName後: %s\n", large.GetName())
}
