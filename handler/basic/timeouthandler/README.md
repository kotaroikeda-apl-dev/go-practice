# TimeoutHandler Example

`http.TimeoutHandler` を使用して、ハンドラーの実行時間に制限を設けるサンプルです。

## 概要

`http.TimeoutHandler` は、指定された `http.Handler` をラップし、実行時間が制限時間を超えた場合に `503 Service Unavailable` エラーを返します。

```go
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```

- `h`: ラップする元のハンドラー
- `dt`: タイムアウトまでの時間
- `msg`: タイムアウト時にボディに表示するメッセージ（空の場合はデフォルトメッセージ）

## 実行方法

1. サーバーを起動します：
   ```bash
   go run main.go
   ```

2. ブラウザまたは `curl` でアクセスします：
   ```bash
   curl -i http://localhost:8080
   ```

## 期待される結果

このサンプルでは、ハンドラーが 2 秒間スリープしますが、タイムアウトは 1 秒に設定されています。そのため、アクセスすると 1 秒後に以下のレスポンスが返ります：

```text
HTTP/1.1 503 Service Unavailable
...
タイムアウトしました（サーバーが混雑しています）
```

