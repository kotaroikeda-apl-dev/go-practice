## Stringer のサンプル（レシート表示）

このサンプルで学べること

- `fmt.Stringer` を実装すると、`fmt.Print*` 系が `String()` を自動で呼び出す仕組み
- 表示用のロジックを `String()` に閉じ込めると、呼び出し側は `fmt.Println(rcpt)` だけで整形済み出力が得られる
- `strings.Builder` を使うと複数行連結を効率的に書ける（`+` 連結より不要なアロケーションを抑制）
- 日付フォーマットの例として `time.Time.Format(time.RFC3339)` を使用

## 実行

```bash
go run main.go
```

## 出力例

```
Receipt #A-42 @ 2024-12-24T12:34:56Z
- Coffee          480
- Donut           260
- Tax              50
TOTAL: 790
```
