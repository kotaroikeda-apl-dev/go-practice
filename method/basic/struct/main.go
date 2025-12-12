package main

import (
	"errors"
	"fmt"
)

// Wallet は残高を持つウォレット口座
type Wallet struct {
	Owner   string
	Balance int
}

// Deposit は入金を行うメソッド（値レシーバーで新しい残高を返す）
func (w Wallet) Deposit(amount int) int {
	return w.Balance + amount
}

// Withdraw は出金を行うメソッド（値レシーバーで新しい残高とエラーを返す）
func (w Wallet) Withdraw(amount int) (int, error) {
	if w.Balance < amount {
		return w.Balance, errors.New("残高不足です")
	}
	return w.Balance - amount, nil
}

// GetBalance は残高を取得するメソッド
func (w Wallet) GetBalance() int {
	return w.Balance
}

// String はWalletを文字列として表現するメソッド
// fmtパッケージ（fmt.Printf、fmt.Printlnなど）で%sや%vを使ってWalletを出力する際に自動的に呼び出される
func (w Wallet) String() string {
	return fmt.Sprintf("%sのウォレット: 残高 %d円", w.Owner, w.Balance)
}

func main() {
	fmt.Println("=== Go Methods ===")

	// ウォレットを作成
	wallet := Wallet{Owner: "田中太郎", Balance: 1000}
	fmt.Printf("初期状態: %s\n", wallet)

	// 入金（値レシーバーで新しい残高を返す）
	wallet.Balance = wallet.Deposit(500)
	fmt.Printf("500円入金後: %s\n", wallet)

	// 出金（成功）
	var err error
	wallet.Balance, err = wallet.Withdraw(300)
	if err != nil {
		fmt.Printf("出金エラー: %v\n", err)
	}
	if err == nil {
		fmt.Printf("300円出金後: %s\n", wallet)
	}

	// 残高確認
	fmt.Printf("現在の残高: %d円\n", wallet.GetBalance())

	// 出金（残高不足でエラー）
	wallet.Balance, err = wallet.Withdraw(2000)
	if err != nil {
		fmt.Printf("出金エラー: %v\n", err)
		fmt.Printf("残高は変更されていない: %s\n", wallet)
	}
	if err == nil {
		fmt.Printf("2000円出金後: %s\n", wallet)
	}

	// 再度出金（成功）
	wallet.Balance, err = wallet.Withdraw(500)
	if err != nil {
		fmt.Printf("出金エラー: %v\n", err)
		return
	}
	fmt.Printf("500円出金後: %s\n", wallet)
}
