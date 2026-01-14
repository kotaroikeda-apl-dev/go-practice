# 標準的な http.FileServer による配信

Go の標準ライブラリのみを使用して、ローカルディスク上のディレクトリ（`./static`）を公開するサンプルです。

## 実行方法

```bash
# 実行（カレントディレクトリに static フォルダが必要です）
go run main.go
```

## 動作確認

サーバー起動後、以下の URL にアクセスしてください。

[http://localhost:8080/static/](http://localhost:8080/static/)
