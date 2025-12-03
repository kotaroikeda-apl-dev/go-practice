# mutex

sync.Mutex を使った基本的な使い方を学習するためのサンプルコード。

## 実行方法

```bash
go run main.go
```

## 学習ポイント

- `sync.Mutex` を使用して共有リソースへのアクセスを保護する方法
- 複数の goroutine から安全にデータを操作する方法
- `sync.WaitGroup` を使用して複数の goroutine の完了を待つ方法
