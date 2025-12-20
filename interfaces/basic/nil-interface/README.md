# Interfaces / Nil Interface

nil なインターフェース値（動的型も動的値も空）にメソッドを呼ぶと実行時エラーになることを確認するサンプルです。

## ポイント

- インターフェースのゼロ値は nil（動的型も値も空）
- nil インターフェースにメソッド呼び出しすると panic（型情報がないため）

## 実行

```bash
go run main.go
```

期待される挙動:

```
(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
```

※ panic を確認したいときは `s.Say()` のコメントアウトを外してください（外したままにすると毎回 panic します）。

