## 概要

「if/switch」と「interface」による条件分岐の比較を、決済処理（CreditCard / PayPay / BankTransfer）で示すサンプルです。

## ディレクトリ構成

- `cmd/demo`: デモエントリポイント
- `internal/common`: 共通型・定数
- `internal/ifstyle`: if/switch スタイルの実装とテスト
- `internal/interfacestyle`: interface スタイルの実装とテスト

## 仕様

- 金額 `amount`、支払い方法 `method` を受け取り、`fee` と `total` を返す
- 手数料仕様（コード・記事共通）
  - CreditCard: 3%（端数切り上げ）、amount >= 1
  - PayPay: 1.5%（端数切り上げ）、amount >= 1、かつ 1,000,000 以下
  - BankTransfer: 手数料 300 円固定、amount >= 1 かつ合計が 300 円未満にならない（amount < 300 はエラー）

## 実行

```bash
cd interfaces/payment-branching
go run ./cmd/demo
```

## テスト

```bash
cd interfaces/payment-branching
go test ./...
```



