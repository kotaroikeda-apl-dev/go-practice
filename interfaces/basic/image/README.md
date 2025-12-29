## image.Image インターフェースのサンプル

学べること

- `image.Image` は `ColorModel()`, `Bounds()`, `At(x, y)` の 3 つのメソッドを持つインターフェース
- `image.NewRGBA` で RGBA 画像を作成し、`image.Image` を満たす
- `At(x, y)` でピクセル値を取得し、`RGBA()` で R/G/B/A を取得できる（16bit なので `>>8` で 8bit に変換）
- `Set(x, y, color)` でピクセル値を設定できる

## 実行

```bash
go run main.go
```

## 出力例

```
bounds: (0,0)-(100,100)
color model: RGBA
at (0, 0): R=0 G=0 B=0 A=0
at (50, 50): R=0 G=0 B=0 A=0
at (99, 99): R=0 G=0 B=0 A=0
after Set(10,10): R=255 G=128 B=64 A=255

--- Gray image ---
bounds: (0,0)-(50,50)
color model: Gray
```


