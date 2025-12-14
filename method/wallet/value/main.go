package main

import (
	"errors"
	"fmt"
	"time"
)

// TransactionType は取引の種類を表す
type TransactionType string

const (
	TransactionTypeDeposit  TransactionType = "入金"
	TransactionTypeWithdraw TransactionType = "出金"
)

// Account は1口座の状態とルールを管理する構造体（ポインタなし版）
// 責務：1口座の状態（残高）とルール（バリデーション）を持つ
// ポインタを使わないため、状態変更時は新しいAccountを返す必要がある
type Account struct {
	balance int // 残高（非公開：必ずmethod経由でアクセス）
}

// NewAccount は新しいAccountを作成する
// コンストラクタパターン：初期状態を保証するため必要
func NewAccount(initialBalance int) (Account, error) {
	if initialBalance < 0 {
		return Account{}, errors.New("初期残高は0以上である必要があります")
	}
	return Account{balance: initialBalance}, nil
}

// Deposit は入金を行う（ポインタなし版）
// 責務：入金に関するルールを適用する（入金額のバリデーション）
// ポインタを使わないため、新しいAccountを返す必要がある
func (a Account) Deposit(amount int) (Account, error) {
	if amount <= 0 {
		return a, errors.New("入金額は0より大きい必要があります")
	}
	// 新しいAccountを作成して返す（元のAccountは変更されない）
	return Account{balance: a.balance + amount}, nil
}

// Withdraw は出金を行う（ポインタなし版）
// 責務：出金に関するルールを適用する（出金額のバリデーション、残高チェック）
// ポインタを使わないため、新しいAccountを返す必要がある
func (a Account) Withdraw(amount int) (Account, error) {
	if amount <= 0 {
		return a, errors.New("出金額は0より大きい必要があります")
	}
	if a.balance < amount {
		return a, errors.New("残高が不足しています")
	}
	// 新しいAccountを作成して返す（元のAccountは変更されない）
	return Account{balance: a.balance - amount}, nil
}

// Balance は現在の残高を取得する
// 責務：残高の取得（balanceフィールドへの直接アクセスを防ぐ）
func (a Account) Balance() int {
	return a.balance
}

// Transaction は入金・出金の履歴を表す構造体
// 不変なデータとして扱うため、フィールドは全て公開
type Transaction struct {
	Type      TransactionType // 取引の種類（入金 or 出金）
	Amount    int             // 金額
	Timestamp time.Time       // 取引時刻
}

// NewTransaction は新しいTransactionを作成する
// コンストラクタパターン：取引時刻を自動設定し、整合性を保つため必要
func NewTransaction(txType TransactionType, amount int) Transaction {
	return Transaction{
		Type:      txType,
		Amount:    amount,
		Timestamp: time.Now(),
	}
}

// Wallet は全体の窓口として機能する構造体（ポインタなし版）
// 責務：外部への窓口（Accountのルールに委譲、履歴管理を担当）
// ポインタを使わないため、状態変更時は新しいWalletを返す必要がある
type Wallet struct {
	account      Account       // 1口座の状態とルール（非公開：Accountのルールに委譲）
	transactions []Transaction // 取引履歴（非公開：Walletが管理）
}

// NewWallet は新しいWalletを作成する
// コンストラクタパターン：AccountとTransactionの初期化を保証するため必要
func NewWallet(initialBalance int) (Wallet, error) {
	account, err := NewAccount(initialBalance)
	if err != nil {
		return Wallet{}, err
	}
	return Wallet{
		account:      account,
		transactions: []Transaction{},
	}, nil
}

// Deposit は入金を行い、取引履歴に追加する（ポインタなし版）
// 責務：Accountにルール適用を委譲し、成功時に履歴を追加する調整役
// ポインタを使わないため、新しいWalletを返す必要がある
func (w Wallet) Deposit(amount int) (Wallet, error) {
	// Accountのルールに委譲（バリデーションと残高更新）
	// ポインタなし版では、新しいAccountが返される
	newAccount, err := w.account.Deposit(amount)
	if err != nil {
		return w, err
	}
	// 履歴管理（Walletの責務）
	// ポインタなし版では、新しいWalletを作成して返す必要がある
	newTransactions := make([]Transaction, len(w.transactions), len(w.transactions)+1)
	copy(newTransactions, w.transactions)
	newTransactions = append(newTransactions, NewTransaction(TransactionTypeDeposit, amount))
	return Wallet{
		account:      newAccount,
		transactions: newTransactions,
	}, nil
}

// Withdraw は出金を行い、取引履歴に追加する（ポインタなし版）
// 責務：Accountにルール適用を委譲し、成功時に履歴を追加する調整役
// ポインタを使わないため、新しいWalletを返す必要がある
func (w Wallet) Withdraw(amount int) (Wallet, error) {
	// Accountのルールに委譲（バリデーションと残高更新）
	// ポインタなし版では、新しいAccountが返される
	newAccount, err := w.account.Withdraw(amount)
	if err != nil {
		return w, err
	}
	// 履歴管理（Walletの責務）
	// ポインタなし版では、新しいWalletを作成して返す必要がある
	newTransactions := make([]Transaction, len(w.transactions), len(w.transactions)+1)
	copy(newTransactions, w.transactions)
	newTransactions = append(newTransactions, NewTransaction(TransactionTypeWithdraw, amount))
	return Wallet{
		account:      newAccount,
		transactions: newTransactions,
	}, nil
}

// Balance は現在の残高を取得する
// 責務：Accountに委譲（Walletは調整役として透過的に提供）
func (w Wallet) Balance() int {
	return w.account.Balance()
}

// Transactions は取引履歴を取得する
// 責務：履歴の取得（Walletが管理する履歴を提供）
// スライスのコピーを返すことで、外部からの履歴改ざんを防ぐ
func (w Wallet) Transactions() []Transaction {
	if len(w.transactions) == 0 {
		return nil
	}
	result := make([]Transaction, len(w.transactions))
	copy(result, w.transactions)
	return result
}

// DisplayHistory は現在の残高と取引履歴を表示する
// 責務：表示ロジック（Walletが調整役として、AccountとTransactionの情報を統合表示）
func (w Wallet) DisplayHistory() {
	fmt.Printf("=== ウォレット残高: %d円 ===\n", w.Balance())
	fmt.Println("取引履歴:")
	if len(w.transactions) == 0 {
		fmt.Println("  取引履歴はありません")
		return
	}
	for i, tx := range w.transactions {
		fmt.Printf("  %d. %s - %d円 (%s)\n",
			i+1,
			tx.Type,
			tx.Amount,
			tx.Timestamp.Format("2006-01-02 15:04:05"),
		)
	}
}

func main() {
	fmt.Println("=== ウォレットシステムのデモ（ポインタなし版） ===\n")
	fmt.Println("※ ポインタを使わないため、各操作で新しいWalletを返す必要があります\n")

	// ウォレットを作成（初期残高1000円）
	wallet, err := NewWallet(1000)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	// 初期状態を表示
	fmt.Println("【初期状態】")
	wallet.DisplayHistory()
	fmt.Println()

	// 入金操作（ポインタなし版では、新しいWalletを受け取る必要がある）
	fmt.Println("【500円入金】")
	wallet, err = wallet.Deposit(500)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		wallet.DisplayHistory()
		fmt.Println()
	}
	if err == nil {
		fmt.Println("入金成功")
		wallet.DisplayHistory()
		fmt.Println()
	}

	// 出金操作（ポインタなし版では、新しいWalletを受け取る必要がある）
	fmt.Println("【300円出金】")
	wallet, err = wallet.Withdraw(300)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		wallet.DisplayHistory()
		fmt.Println()
	}
	if err == nil {
		fmt.Println("出金成功")
		wallet.DisplayHistory()
		fmt.Println()
	}

	// 不正な操作のテスト
	fmt.Println("【不正な操作のテスト】")

	// マイナス入金を試みる
	fmt.Println("1. マイナス入金を試みる:")
	_, err = wallet.Deposit(-100)
	if err != nil {
		fmt.Printf("   エラー（期待通り）: %v\n", err)
	}

	// 残高不足の出金を試みる
	fmt.Println("2. 残高不足の出金を試みる:")
	_, err = wallet.Withdraw(10000)
	if err != nil {
		fmt.Printf("   エラー（期待通り）: %v\n", err)
	}

	// 0円の出金を試みる
	fmt.Println("3. 0円の出金を試みる:")
	_, err = wallet.Withdraw(0)
	if err != nil {
		fmt.Printf("   エラー（期待通り）: %v\n", err)
	}

	fmt.Println()
	wallet.DisplayHistory()
	fmt.Println()

	// 複数の取引を追加（ポインタなし版では、毎回新しいWalletを受け取る必要がある）
	fmt.Println("【複数の取引を追加】")
	wallet, err = wallet.Deposit(200)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}
	wallet, err = wallet.Withdraw(100)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}
	wallet, err = wallet.Deposit(50)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}
	wallet.DisplayHistory()
}
