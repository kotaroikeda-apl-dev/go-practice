# concurrency/basic

Go 言語の並行処理の基本（goroutine と channel）を学習するためのサンプルコード。

## 概要

goroutine の基本的な使い方と、channel を使用して goroutine 間でデータを送受信する基本的な例を示しています。

## 実行方法

```bash
go run main.go
```

または

```bash
go run .
```

## ファイル構成

- `main.go` - 各例を実行するメイン関数
- `README.md` - このファイル
- `examples/` - サンプルコードのディレクトリ
  - `goroutine.go` - goroutine の基本的な例
  - `channelBasic.go` - バッファなしチャンネルの例
  - `channelBuffered.go` - バッファ付きチャンネルの例
  - `channelRangeClose.go` - range と close を使った例
  - `channelSelect.go` - select 文を使った例

## 学習ポイント

### goroutine（goroutine.go）

- `go` キーワードを使用して goroutine を起動する方法
- 複数の goroutine が並行して実行される様子
- goroutine と通常の関数呼び出しの違い

### バッファなしチャンネル（channelBasic.go）

- channel の作成方法（`make(chan type)`）
- goroutine から channel へのデータ送信（`c <- value`）
- channel からのデータ受信（`<-c`）
- goroutine と channel を使った基本的な通信パターン

### バッファ付きチャンネル（channelBuffered.go）

- バッファ付きチャンネルの作成（`make(chan type, size)`）
- バッファが満杯の時のブロック動作
- チャンネルの送受信は同期操作であることの理解

### range と close（channelRangeClose.go）

- チャンネルを閉じる方法（`close(c)`）
- `range` でチャンネルから値を受信する方法
- チャンネルが閉じられると range ループが終了すること

### select 文（channelSelect.go）

- `select` 文で複数のチャンネル操作を待機する方法
- 最初に準備できたチャンネル操作を実行する動作
- 複数のチャンネルから受信する例
