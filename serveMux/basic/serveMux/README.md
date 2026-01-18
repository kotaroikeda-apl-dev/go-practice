# ServeMux (Go 1.22+)

Go 1.22 で導入された `http.ServeMux` の新しいパターンマッチング機能のサンプルです。

## 特徴

### Patterns (パターン)
パターンは、リクエストのメソッド、ホスト、パスにマッチさせることができます。

- `"/index.html"`: 任意のホスト・メソッドで `/index.html` にマッチ。
- `"GET /static/"`: `/static/` で始まる GET リクエストにマッチ。
- `"example.com/"`: ホスト `example.com` への任意のリクエストにマッチ。
- `"example.com/{$}"`: ホスト `example.com` かつパス `/` にマッチ。
- `"/b/{bucket}/o/{objectname...}"`: ワイルドカードセグメント。`{bucket}` は1セグメント、`{objectname...}` は残りすべてにマッチ。

### Precedence (優先順位)
複数のパターンがマッチする場合、最も具体的なパターンが優先されます。
例: `/images/thumbnails/` は `/images/` より具体的です。

### Trailing-slash redirection (末尾スラッシュのリダイレクト)
末尾スラッシュを持つパターン（または `...` ワイルドカード）を登録すると、末尾スラッシュなしのリクエストは自動的にリダイレクトされます。

### Request sanitizing (リクエストのサニタイズ)
`.` や `..` セグメント、繰り返しのスラッシュをクリーンな URL にリダイレクトします。

### {$}: 完全一致
`{$}` は URL の末尾にのみマッチします。
例: `"/{$}"` は `/` のみにマッチし、`"/"` はすべてのパスにマッチします。

## 実行方法

```bash
go run serveMux/basic/serveMux/main.go
```

## テスト

```bash
# パスのみ
curl http://localhost:8080/index.html

# メソッド指定
curl -X GET http://localhost:8080/static/foo

# ワイルドカード
curl http://localhost:8080/b/my-bucket/o/path/to/object.txt

# 完全一致
curl http://localhost:8080/
```
