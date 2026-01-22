# Goルーター選定ガイド (2026年版)

Go 1.22 で強化された標準の `ServeMux` と、主要な外部ルーター・フレームワーク（chi, Gin, Echo）を比較検討するためのサンプルコード一式です。

## 実装バリエーション
各ディレクトリは独立した Go モジュールとして構成されており、個別に実行・コピーが可能です。

- `standard/`: 標準ライブラリのみの実装
- `chi/`: `chi` ルーターを使用した実装
- `gin/`: `Gin` フレームワークを使用した実装
- `echo/`: `Echo` フレームワークを使用した実装

## 実行方法

各ディレクトリに移動して実行してください。

```bash
# 標準ライブラリ版
cd standard && go run main.go

# chi版
cd chi && go run main.go

# Gin版
cd gin && go run main.go

# Echo版
cd echo && go run main.go
```

サーバーはいずれも `http://localhost:8080` で起動します。
