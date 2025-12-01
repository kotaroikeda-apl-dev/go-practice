# channel

channel の基本的な使い方を学習するためのサンプルコード。

## 概要

channel を使用して goroutine 間でデータを送受信する基本的な例を示しています。

## 実行方法

```bash
go run *.go
```

または

```bash
go run .
```

**注意**: `go run main.go` だけでは他のファイルがコンパイルされないため、上記のコマンドを使用してください。

## ファイル構成

- `main.go` - 各例を実行するメイン関数
- `basic.go` - 基本的なチャンネル（バッファなし）の例
- `buffered.go` - バッファ付きチャンネルの例
- `range_close.go` - range と close を使った例

## 学習ポイント

### 基本的なチャンネル（basic.go）

- channel の作成方法（`make(chan type)`）
- goroutine から channel へのデータ送信（`c <- value`）
- channel からのデータ受信（`<-c`）
- goroutine と channel を使った基本的な通信パターン

### バッファ付きチャンネル（buffered.go）

- バッファ付きチャンネルの作成（`make(chan type, size)`）
- バッファが満杯の時のブロック動作
- チャンネルの送受信は同期操作であることの理解

### range と close（range_close.go）

- チャンネルを閉じる方法（`close(c)`）
- `range` でチャンネルから値を受信する方法
- チャンネルが閉じられると range ループが終了すること
