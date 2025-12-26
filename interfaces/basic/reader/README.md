## io.Reader のシンプル例

学べること

- `io.Reader` の `Read([]byte) (int, error)` でバイト列を読み取る
- `Read` は読み取ったバイト数 `n` と `err` を返し、`io.EOF` で終了を通知する
- `strings.NewReader` で文字列を `io.Reader` として扱える
- バッファサイズより小さいデータでも `n` で実際の読み取り量を確認できる

## 実行

```bash
go run main.go
```

## 出力例

```
read 8 bytes: "Hello, R"
read 6 bytes: "eader!"
```
