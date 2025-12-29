# Interfaces / Empty Interface

メソッドを一つも持たない空インターフェース（`interface{}`）のシンプルな例です。異なる型を順に代入して、動的型と値を確認します。

## ポイント

- `interface{}` はメソッドゼロのため全ての型が代入可能
- ゼロ値は `<nil>`（動的型も値も空）
- 異なる型（int/string/bool など）を順番に代入して動的型の変化を確認

## 実行

```bash
go run main.go
```

出力例:

```
(<nil>, <nil>)
(99, int)
(gopher, string)
(true, bool)
```



