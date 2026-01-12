# http.RedirectHandler

このサンプルでは、標準ライブラリの `http.RedirectHandler` を使用する方法を学びます。

## ディレクトリ構造

`handler/basic/redirecthandler/main.go`

## 実行と確認手順

1. **サーバーを起動する**

   ```bash
   go run handler/basic/redirecthandler/main.go
   ```

2. **動作確認**

   - **通常のリクエスト**: [http://localhost:8080/](http://localhost:8080/)
     - デモの説明が表示されます。
   - **リダイレクトのテスト**: [http://localhost:8080/old-path](http://localhost:8080/old-path)
     - 自動的に [http://localhost:8080/new-path](http://localhost:8080/new-path) へリダイレクトされ、"Welcome to the new path!" と表示されます。

3. **ターミナルでの確認 (curl)**

   ```bash
   curl -I http://localhost:8080/old-path
   ```

   実行結果に `HTTP/1.1 301 Moved Permanently` と `Location: /new-path` が含まれていることが確認できます。

## 解説

`http.RedirectHandler(url string, code int)` は、受信したすべてのリクエストを指定された `url` に指定されたステータスコード `code` でリダイレクトする `http.Handler` を作成します。

主なステータスコード:
- `http.StatusMovedPermanently` (301): 恒久的な移転
- `http.StatusFound` (302): 一時的な移転
- `http.StatusSeeOther` (303): 別の場所を参照（POST後のリダイレクトなど）

特定のパスへのアクセスを別のURLへ転送したい場合に非常に便利です。
