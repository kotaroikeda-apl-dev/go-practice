package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	// RGBA 画像を作成（100x100、初期値は透明な黒）
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// 画像の境界（左上と右下の座標）を確認
	fmt.Printf("bounds: %v\n", m.Bounds())

	// カラーモデル（色空間）を確認
	cm := m.ColorModel()
	if cm == color.RGBAModel {
		fmt.Println("color model: RGBA")
	}
	if cm == color.GrayModel {
		fmt.Println("color model: Gray")
	}
	if cm == color.NRGBAModel {
		fmt.Println("color model: NRGBA")
	}
	if cm != color.RGBAModel && cm != color.GrayModel && cm != color.NRGBAModel {
		fmt.Printf("color model: %T\n", cm)
	}

	// いくつかの座標のピクセル値を確認（初期値は透明）
	coords := []struct{ x, y int }{
		{0, 0},
		{50, 50},
		{99, 99},
	}

	for _, c := range coords {
		clr := m.At(c.x, c.y)
		r, g, b, a := clr.RGBA()
		fmt.Printf("at (%d, %d): R=%d G=%d B=%d A=%d\n", c.x, c.y, r>>8, g>>8, b>>8, a>>8)
	}

	// ピクセルを設定して確認
	m.Set(10, 10, color.RGBA{R: 255, G: 128, B: 64, A: 255})
	clr := m.At(10, 10)
	r, g, b, a := clr.RGBA()
	fmt.Printf("after Set(10,10): R=%d G=%d B=%d A=%d\n", r>>8, g>>8, b>>8, a>>8)

	// グレースケール画像も作成してカラーモデルを確認
	fmt.Println("\n--- Gray image ---")
	gray := image.NewGray(image.Rect(0, 0, 50, 50))
	fmt.Printf("bounds: %v\n", gray.Bounds())
	cmGray := gray.ColorModel()
	if cmGray == color.RGBAModel {
		fmt.Println("color model: RGBA")
	}
	if cmGray == color.GrayModel {
		fmt.Println("color model: Gray")
	}
	if cmGray == color.NRGBAModel {
		fmt.Println("color model: NRGBA")
	}
	if cmGray != color.RGBAModel && cmGray != color.GrayModel && cmGray != color.NRGBAModel {
		fmt.Printf("color model: %T\n", cmGray)
	}
}
