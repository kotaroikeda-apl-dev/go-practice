# channel

channel の基本的な使い方を学習するためのサンプルコード。

## 概要

channel を使用して goroutine 間でデータを送受信する基本的な例を示しています。

## 実行方法

```bash
go run main.go
```

## 学習ポイント

- channel の作成方法（`make(chan type)`）
- goroutine から channel へのデータ送信（`c <- value`）
- channel からのデータ受信（`<-c`）
- goroutine と channel を使った基本的な通信パターン

## 動作

- goroutine 内で channel に 4 つのメッセージを送信
- main 関数で channel からメッセージを受信して出力
