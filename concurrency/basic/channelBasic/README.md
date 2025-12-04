# channelBasic

バッファなしチャンネルの基本的な使い方を学習するためのサンプルコード。

## 実行方法

```bash
go run main.go
```

## 学習ポイント

- channel の作成方法（`make(chan type)`）
- goroutine から channel へのデータ送信（`c <- value`）
- channel からのデータ受信（`<-c`）
- goroutine と channel を使った基本的な通信パターン
