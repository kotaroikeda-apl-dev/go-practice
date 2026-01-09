# http.Handler の基本

このサンプルでは、`http.Handler` インターフェースを実装する方法を学びます。

## ディレクトリ構造

`handler/basic/main.go`

## 実行と確認手順

1. **サーバーを起動する**

   ```bash
   go run handler/basic/main.go
   ```

2. **リクエストを送る（動作確認）**
   別のターミナルを開くか、ブラウザで以下の URL にアクセスしてください。

   - **ブラウザの場合**: [http://localhost:8080/world](http://localhost:8080/world) にアクセス
   - **ターミナルの場合**:
     ```bash
     curl http://localhost:8080/world
     ```

3. **結果**
   ブラウザやターミナルに以下のように表示されれば成功です！
   `こんにちは! あなたがアクセスしたパスは: /world です。`

## 解説

`http.Handler` インターフェースは、以下のメソッドを持つ型であれば何でも実装できます。

```go
ServeHTTP(w http.ResponseWriter, r *http.Request)
```

このサンプルでは `MyHandler` 構造体にこのメソッドを実装しています。
