# echo

Web フレームワーク `Echo` を活用しつつ、標準ライブラリとの共存を図る構成案です。

## 特徴
- 充実した内蔵ミドルウェアや、`echo.Context` による強力なヘルパー機能が利用可能です。
- 本サンプルでは、標準の `http.HandlerFunc` と Echo を組み合わせるブリッジ手法を提示しています。
- フレームワーク独自の機能（Context）への依存を最小限に抑える設計を確認できます。

## 実行方法
```bash
go run main.go
```

## 動作確認
```bash
# ヘルスチェック
curl http://localhost:8080/health

# ユーザー詳細 (Echoから標準ハンドラへのブリッジ確認)
curl http://localhost:8080/users/123
```

