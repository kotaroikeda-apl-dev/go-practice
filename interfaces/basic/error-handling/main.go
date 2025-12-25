package main

import "fmt"

// PaymentError は操作名とメッセージを持つカスタムエラー。
// Error で見やすい整形を行い、呼び出し側は値比較で分岐できる。
type PaymentError struct {
	Op  string
	Msg string
}

func (e PaymentError) Error() string {
	if e.Op == "" {
		return e.Msg
	}
	return fmt.Sprintf("%s: %s", e.Op, e.Msg)
}

var (
	ErrInvalidAmount = PaymentError{Op: "validation", Msg: "amount must be positive"}
	ErrGateway       = PaymentError{Op: "gateway", Msg: "gateway unreachable"}
)

func main() {
	tests := []int{-1, 500, 42} // 不正 / 上流エラー / 正常

	for _, amount := range tests {
		if err := pay(amount); err != nil {
			fmt.Printf("pay(%d) failed: %v\n", amount, err)
			switch err {
			case ErrInvalidAmount:
				fmt.Println(" -> ask user to enter a positive amount")
			case ErrGateway:
				fmt.Println(" -> retry later or use another gateway")
			default:
				fmt.Println(" -> unexpected error")
			}
			continue
		}
		fmt.Printf("pay(%d) ok\n", amount)
	}
}

func pay(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount == 500 { // デモ用に上流エラーを発生
		return ErrGateway
	}
	return nil
}
