# Interfaces / Distance

値レシーバーとポインタレシーバーがインターフェース実装に与える違いを、自前の `Distance` インターフェースで確認するサンプルです。

## ポイント

- `Distance` は `Distance() float64` を要求するインターフェース
- `Meter` は値レシーバーでメソッドを持つため、そのまま代入できる
- `Walker` はポインタレシーバーでメソッドを持つため、`*Walker` でなければ代入できない
- 値の `Walker` を代入しようとするとコンパイルエラーになる（コード内でコメントアウト）

## 実行

```bash
go run main.go
```

出力例:

```
trail distance: 120.5m
hiker distance: 960.0m
sqrt(2) as a bonus: 1.4142
```
