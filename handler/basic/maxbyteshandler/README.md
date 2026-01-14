# MaxBytesHandler

`http.MaxBytesHandler` は、HTTP リクエストのボディのサイズを制限するためのミドルウェア（ハンドラー）です。

## 概要

`MaxBytesHandler` は、指定したバイト数を超えてリクエストボディが読み取られないように `http.MaxBytesReader` でリクエストをラップします。これにより、悪意のある巨大なリクエストや予期しない大量のデータ送信によるサーバーリソース（メモリなど）の枯渇を防ぐことができます。

## 使い方

```go
func MaxBytesHandler(h Handler, n int64) Handler
```

- `h`: ラップする対象の `http.Handler`
- `n`: 許容する最大バイト数

## 実行方法

1. サーバーを起動します：
   ```bash
   go run main.go
   ```

2. 制限内のリクエストを送信する：
   ```bash
   curl -X POST -d "hello world" http://localhost:8080
   ```

3. 制限を超えるリクエストを送信する：
   ```bash
   # 1KB (1024バイト) を超えるデータを送信
   head -c 2000 /dev/zero | curl -X POST --data-binary @- http://localhost:8080
   ```

