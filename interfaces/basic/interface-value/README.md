# Interfaces / Interface Value

インターフェース値が「動的型」と「動的な値」の組み合わせであることを確認するサンプルです。

## ポイント

- インターフェースのゼロ値は `<nil>`（動的型も値も空）で、メソッド呼び出し不可
- 値レシーバーの `Greeting` は値のまま代入できる
- ポインタレシーバーの `Alert` は `*Alert` だけが代入可能
- `nil` な `*Alert` を代入すると、インターフェース自体は非 nil（動的型は \*Alert、動的値は nil）になる例を確認

## 実行

```bash
go run main.go
```

出力例:

```
zero: <nil>
greeting: type=main.Greeting value={Gopher} message="Hello, Gopher"
nil alert value: type=*main.Alert value=<nil> message="(nil alert)"
alert: type=*main.Alert value=&{warn 503} message="[warn] code=503"
```



