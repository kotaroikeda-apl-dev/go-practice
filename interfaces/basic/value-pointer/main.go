package main

import (
	"fmt"
)

// Distance は移動距離を計算するインターフェース
type Distance interface {
	Distance() float64
}

type Meter float64

// 値レシーバー: そのままインターフェースを満たす
func (m Meter) Distance() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

type Walker struct {
	Steps   int
	StepLen float64
}

// ポインタレシーバー: *Walker でなければ Distance を実装しない
func (w *Walker) Distance() float64 {
	return float64(w.Steps) * w.StepLen
}

func main() {
	var d Distance

	trail := Meter(-120.5)
	hiker := Walker{Steps: 1200, StepLen: 0.8}

	d = trail // Meter は値レシーバーなので OK
	fmt.Printf("trail distance: %.1fm\n", d.Distance())

	d = &hiker // *Walker だけが Distance を実装
	fmt.Printf("hiker distance: %.1fm\n", d.Distance())

	// 値の Walker は Distance を実装していないためコンパイルエラー
	// d = hiker
}
