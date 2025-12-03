# channelRangeClose

range と close を使ったチャンネルの基本的な使い方を学習するためのサンプルコード。

## 実行方法

```bash
go run main.go
```

## 学習ポイント

- チャンネルを閉じる方法（`close(c)`）
- `range` でチャンネルから値を受信する方法
- チャンネルが閉じられると range ループが終了すること
