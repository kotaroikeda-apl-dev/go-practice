# http.NotFoundHandler

このサンプルでは、標準ライブラリの `http.NotFoundHandler` を使用する方法を学びます。

## ディレクトリ構造

`handler/basic/notfoundhandler/main.go`

## 実行と確認手順

1. **サーバーを起動する**

   ```bash
   go run handler/basic/notfoundhandler/main.go
   ```

2. **動作確認**

   - **通常のアクセス**: [http://localhost:8080/](http://localhost:8080/)
     - ウェルカムメッセージが表示されます。
   - **404 ハンドラのテスト**: [http://localhost:8080/missing](http://localhost:8080/missing)
     - `404 page not found` と表示されます。

3. **ターミナルでの確認 (curl)**

   ```bash
   curl -I http://localhost:8080/missing
   ```

   実行結果に `HTTP/1.1 404 Not Found` が含まれていることが確認できます。

## 解説

`http.NotFoundHandler()` は、呼び出されるたびに `404 page not found` というテキストと `404` ステータスコードを返すシンプルな `http.Handler` を作成します。

特定のパスを無効にしたい場合や、カスタムマルチプレクサでデフォルトの挙動として 404 を返したい場合などに便利です。
