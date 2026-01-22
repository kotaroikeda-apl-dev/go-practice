# standard

Go 1.22 で強化された標準の `net/http.ServeMux` のみを使用した実装構成です。

## 特徴
- 外部ライブラリへの依存が一切ありません。
- Go 1.22 から導入された HTTP メソッド指定（`GET`, `POST` 等）とパスパラメータ（`{id}`）を使用しています。
- ミドルウェアは関数をラップする標準的な手法で実装しています。

## 実行方法
```bash
go run main.go
```

## 動作確認
```bash
# ヘルスチェック
curl http://localhost:8080/health

# ユーザー詳細 (パラメータの抽出確認)
curl http://localhost:8080/users/123
```

