# goToolingBasics（Go module / package / go tool を小さな CLI で学ぶ）

この `goToolingBasics/` は、Go の基本要素である **module / package / import / 依存管理 / テスト** を扱うミニサンプルです。

題材は「入力テキストを整形して、扱いやすい形（例: slug）にする小さな CLI」です。

## ここで触れること

- **モジュール**: `go.mod` があるディレクトリが module の境界になる
- **パッケージ**: `cmd/`（コマンド）と `textfx/`（ライブラリ）で分割する
- **import path**: `module path + サブディレクトリ`（例: `example.com/go-practice/goToolingBasics/textfx`）
- **外部依存**: `go mod tidy` で `go.mod` / `go.sum` に確定する（この例では `go-cmp`）
- **テスト**: `go test ./...` で package 単位に検証する

## 構成

- `go.mod`: module 定義
- `cmd/goToolingBasics/main.go`: 実行可能コマンド（`package main`）
- `textfx/`: 自作パッケージ
  - `textfx.go`:
    - 変換ステップ: `Step`（`func(string) string`）
    - 合成: `ApplyAll`（ステップを左から順に適用）
    - 機能: slug 化 / 桁伏せ / 切り詰め
      - slug（スラッグ）化: タイトル等の文字列を「URL やファイル名に使いやすい識別子」に変換（例: `"Go Practice 2026!"` → `"go-practice-2026"`）
  - `textfx_test.go`: テーブル駆動テスト（比較に `go-cmp/cmp` を使用）

## 使い方

```bash
# go.mod ファイルが存在するディレクトリに移動
cd goToolingBasics

# 依存関係の確定（外部モジュールの取得 + go.mod/go.sum 更新）
go mod tidy

# テスト
go test ./...

# 実行（デフォルト値で動かす）
go run ./cmd/goToolingBasics

# 例: 任意の文字列を slug にする
go run ./cmd/goToolingBasics -text 'How to Write Go Code 2026!'

# 例: 数字マスクを無効化
go run ./cmd/goToolingBasics -text 'order-2026-01-01' -redact-digits=false

# 例: rune 単位で短くする
go run ./cmd/goToolingBasics -text 'こんにちは世界' -max-runes 5 -slug=false
```

## go install（任意）

```bash
cd goToolingBasics
go install ./cmd/goToolingBasics

# インストール先（GOBIN / GOPATH の影響を受ける）
go list -f '{{.Target}}' ./cmd/goToolingBasics
```

必要なら `go env -w GOBIN=...` でインストール先を固定できます（解除は `go env -u GOBIN`）。

## Makefile（任意）

```bash
make tidy
make test
make run
make install
```

それぞれの意味:

- **`make tidy`**: `go mod tidy`（依存関係を整理して `go.mod` / `go.sum` を整合させる）
- **`make test`**: `go test ./...`（モジュール配下の全パッケージをテスト）
- **`make run`**: `go run ./cmd/goToolingBasics`（CLI をビルドせずに実行）
- **`make install`**: `go install ./cmd/goToolingBasics`（CLI をビルドしてインストール）
