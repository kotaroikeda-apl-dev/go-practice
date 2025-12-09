package main

import (
	"fmt"
)

// Temperature は温度を表す型（float64のエイリアス）
type Temperature float64

// Celsius は摂氏温度を取得するメソッド
func (t Temperature) Celsius() float64 {
	return float64(t)
}

// Fahrenheit は華氏温度に変換するメソッド
func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

// Kelvin はケルビン温度に変換するメソッド
func (t Temperature) Kelvin() float64 {
	return float64(t) + 273.15
}

// String はTemperatureを文字列として表現するメソッド
// fmtパッケージ（fmt.Printf、fmt.Printlnなど）で%sや%vを使ってTemperatureを出力する際に自動的に呼び出される
func (t Temperature) String() string {
	return fmt.Sprintf("%.2f°C", t)
}

func main() {
	fmt.Println("=== Go Methods on Non-Struct Types ===")

	// 温度を作成
	temp := Temperature(25.5)
	fmt.Printf("初期温度: %s\n", temp)

	// 摂氏温度を取得
	fmt.Printf("摂氏: %.2f°C\n", temp.Celsius())

	// 華氏温度に変換
	fmt.Printf("華氏: %.2f°F\n", temp.Fahrenheit())

	// ケルビン温度に変換
	fmt.Printf("ケルビン: %.2fK\n", temp.Kelvin())

	// 別の温度で試す
	temp2 := Temperature(0)
	fmt.Printf("\n0度の場合:\n")
	fmt.Printf("摂氏: %.2f°C\n", temp2.Celsius())
	fmt.Printf("華氏: %.2f°F\n", temp2.Fahrenheit())
	fmt.Printf("ケルビン: %.2fK\n", temp2.Kelvin())
}
