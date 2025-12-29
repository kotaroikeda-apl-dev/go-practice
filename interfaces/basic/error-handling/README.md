## error インターフェースの学習用サンプル

学べること

- `error` は `Error() string` を持つインターフェースで、`fmt` が自動で呼び出す
- `PaymentError` で「どの操作で起きたか」を文字列整形しつつ、値比較で分岐できる（`switch err` で扱いやすい）
- sentinel エラー（`ErrInvalidAmount`, `ErrGateway`）をカスタム型として定義し、用途別に使い分ける

## 実行

```bash
go run main.go
```

## 出力例

```
pay(-1) failed: validation: amount must be positive
 -> ask user to enter a positive amount
pay(500) failed: gateway: gateway unreachable
 -> retry later or use another gateway
pay(42) ok
```
