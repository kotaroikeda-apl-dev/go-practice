# http.HandlerFunc の使用 (関数スタイル)

このサンプルでは、`http.HandlerFunc` アダプターを使って、普通の関数を `http.Handler` インターフェースとして扱う方法を学びます。

## 仕組み

`http.HandlerFunc` は、特定のシグネチャを持つ関数を `http.Handler` に変換する便利な型です。
画像の説明にあった通り、これを使うことで構造体を作らなくても関数だけでハンドラを定義できます。

## 実行と確認手順

1. **サーバーを起動する**

   ```bash
   go run handler/basic/handlerfunc/main.go
   ```

2. **アクセスする**
   [http://localhost:8080](http://localhost:8080)
