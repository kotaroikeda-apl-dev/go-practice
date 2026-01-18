# ServeMux.Handle

`http.ServeMux` の `Handle` メソッドの使用例です。

## 特徴

- `Handle` メソッドは、指定されたパターンのハンドラーを登録します。
- 第2引数には `http.Handler` インターフェース（`ServeHTTP(ResponseWriter, *Request)` メソッドを持つ型）を渡します。
- もし登録しようとしたパターンが既に登録済みのものと競合する場合、`Handle` は**パニック**を起こします。

## Handle と HandleFunc の違い

- `Handle(pattern string, handler Handler)`: `http.Handler` インターフェースを実装したオブジェクトを登録します。
- `HandleFunc(pattern string, handler func(ResponseWriter, *Request))`: ハンドラー関数を直接登録します。内部的には `Handle(pattern, HandlerFunc(handler))` を呼び出しています。

## 実行方法

```bash
go run main.go
```

## 動作確認

```bash
curl http://localhost:8080/api/v1
curl http://localhost:8080/hello
```

