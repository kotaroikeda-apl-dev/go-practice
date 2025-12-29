## Type Switch のサンプル例

- interface{} に入った複数の型を type switch で振り分ける
- string/int/float64/[]byte/nil/未知型 をまとめて処理

## 実行

```bash
go run main.go
```

## 出力例

```
nil value
string len=6 val="gopher"
int doubled=42
float ceil-ish=4
bytes len=3
```


