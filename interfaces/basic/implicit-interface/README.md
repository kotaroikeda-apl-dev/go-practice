# Interfaces / Implicit Interface

暗黙的にインターフェースを実装する例。`Greeter` を満たすために「implements」などは不要で、メソッド集合が一致すれば代入できる。

## ポイント

- `Person` は値レシーバーで `Greet` を持つため、`Person` 値をそのまま `Greeter` に代入可能
- `Robot` はポインタレシーバーで `Greet` を持つため、`*Robot` でなければ `Greeter` を満たさない
- コメントアウトしている `g = Robot{...}` はコンパイルエラー例（ポインタレシーバーゆえ）

## 実行

```bash
go run main.go
```

出力例:

```
Hello, Taro
Beep-042
```

