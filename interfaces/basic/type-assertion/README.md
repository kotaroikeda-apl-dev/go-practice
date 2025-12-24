## Type Assertion のサンプル

- 単一戻り値の型アサーション（型不一致なら panic）
- 2 値戻りの型アサーション（失敗しても panic しない）

## 実行

```bash
go run main.go
```

## 出力例

```
as string: hello assertion
as int: 0 (ok=false)
as float64: 3.14 (ok=true)
```
