# ウォレットシステム（ポインタなし版）

このディレクトリでは、**ポインタを使わない**実装でウォレットシステムを実装しています。ポインタあり版（`../pointer/`）と比較することで、ポインタの便利さを理解できます。

## 目的

- ポインタを使わない実装の不便さを体験する
- ポインタあり版と比較して、ポインタの利点を理解する
- 値レシーバーを使った実装の課題を学ぶ

## ポインタなし版の特徴

### 1. 値レシーバーを使用

全てのメソッドが値レシーバーを使用しています：

```go
// ポインタなし版
func (w Wallet) Deposit(amount int) (Wallet, error) {
    // 新しいWalletを作成して返す必要がある
    newAccount, err := w.account.Deposit(amount)
    // ...
    return Wallet{account: newAccount, transactions: newTransactions}, nil
}
```

### 2. 状態変更時に新しい構造体を返す必要がある

ポインタを使わないため、状態を変更するたびに新しい構造体を作成して返す必要があります：

```go
// 使用例
wallet, err = wallet.Deposit(500)  // 毎回新しいWalletを受け取る必要がある
wallet, err = wallet.Withdraw(300) // 毎回代入が必要
```

### 3. スライスのコピーが発生

履歴を追加するたびに、既存のスライスをコピーして新しいスライスを作成する必要があります：

```go
newTransactions := make([]Transaction, len(w.transactions), len(w.transactions)+1)
copy(newTransactions, w.transactions)
newTransactions = append(newTransactions, NewTransaction(...))
```

## ポインタあり版との比較

### コードの比較

| 項目           | ポインタなし版                      | ポインタあり版                   |
| -------------- | ----------------------------------- | -------------------------------- |
| レシーバー     | 値レシーバー `(w Wallet)`           | ポインタレシーバー `(w *Wallet)` |
| 戻り値         | `(Wallet, error)`                   | `error`                          |
| 使用例         | `wallet, err = wallet.Deposit(500)` | `err = wallet.Deposit(500)`      |
| メモリ効率     | 毎回コピーが発生                    | ポインタで効率的                 |
| コードの簡潔性 | 冗長                                | シンプル                         |

### 実装の比較

#### ポインタなし版

```go
func (w Wallet) Deposit(amount int) (Wallet, error) {
    newAccount, err := w.account.Deposit(amount)
    if err != nil {
        return w, err
    }
    newTransactions := make([]Transaction, len(w.transactions), len(w.transactions)+1)
    copy(newTransactions, w.transactions)
    newTransactions = append(newTransactions, NewTransaction(...))
    return Wallet{
        account:      newAccount,
        transactions: newTransactions,
    }, nil
}
```

#### ポインタあり版

```go
func (w *Wallet) Deposit(amount int) error {
    if err := w.account.Deposit(amount); err != nil {
        return err
    }
    w.transactions = append(w.transactions, NewTransaction(...))
    return nil
}
```

## ポインタなし版の不便な点

1. **毎回新しい構造体を返す必要がある**

   - メモリ効率が悪い
   - コードが冗長になる

2. **呼び出し側で毎回代入が必要**

   - `wallet, err = wallet.Deposit(500)` のように毎回代入が必要
   - 代入を忘れると状態が更新されない

3. **スライスのコピーが発生**

   - 履歴追加のたびに既存のスライスをコピーする必要がある
   - パフォーマンスが悪い

4. **コードが複雑になる**
   - 新しい構造体を作成する処理が必要
   - エラーハンドリングが複雑になる

## 実行方法

```bash
go run main.go
```

## 出力例

```
=== ウォレットシステムのデモ（ポインタなし版） ===

※ ポインタを使わないため、各操作で新しいWalletを返す必要があります

【初期状態】
=== ウォレット残高: 1000円 ===
取引履歴:
  取引履歴はありません

【500円入金】
入金成功
=== ウォレット残高: 1500円 ===
取引履歴:
  1. 入金 - 500円 (2025-12-14 07:15:04)

【300円出金】
出金成功
=== ウォレット残高: 1200円 ===
取引履歴:
  1. 入金 - 500円 (2025-12-14 07:15:04)
  2. 出金 - 300円 (2025-12-14 07:15:04)

【不正な操作のテスト】
1. マイナス入金を試みる:
   エラー（期待通り）: 入金額は0より大きい必要があります
2. 残高不足の出金を試みる:
   エラー（期待通り）: 残高が不足しています
3. 0円の出金を試みる:
   エラー（期待通り）: 出金額は0より大きい必要があります

=== ウォレット残高: 1200円 ===
取引履歴:
  1. 入金 - 500円 (2025-12-14 07:15:04)
  2. 出金 - 300円 (2025-12-14 07:15:04)

【複数の取引を追加】
=== ウォレット残高: 1350円 ===
取引履歴:
  1. 入金 - 500円 (2025-12-14 07:15:04)
  2. 出金 - 300円 (2025-12-14 07:15:04)
  3. 入金 - 200円 (2025-12-14 07:15:04)
  4. 出金 - 100円 (2025-12-14 07:15:04)
  5. 入金 - 50円 (2025-12-14 07:15:04)
```

## 学べること

1. **ポインタなし版の不便さ**: 状態変更時に新しい構造体を返す必要がある
2. **メモリ効率**: ポインタを使うことで、コピーを避けられる
3. **コードの簡潔性**: ポインタを使うことで、コードがシンプルになる
4. **実用性**: 実務ではポインタレシーバーを使うことが多い理由を理解できる

## 関連ディレクトリ

- `../pointer/` - ポインタあり版の実装（比較用）
