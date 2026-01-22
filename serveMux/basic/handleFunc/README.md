# HandleFunc の基本

`http.ServeMux` の `HandleFunc` メソッドを使用してハンドラー関数を登録するサンプルです。

## 特徴

- `func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))` というシグネチャを持ちます。
- 指定されたパターンに対してハンドラー関数を直接登録できます。
- **注意**: すでに登録されているパターンと衝突する場合、`HandleFunc` はパニックを起こします。

## 実行方法

```bash
go run main.go
```

## 動作確認

```bash
curl http://localhost:8080/hello
```

