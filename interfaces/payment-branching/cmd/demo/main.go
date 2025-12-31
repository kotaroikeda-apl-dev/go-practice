package main

import (
	"fmt"

	"github.com/example/payment-branching/internal/common"
	"github.com/example/payment-branching/internal/ifstyle"
	"github.com/example/payment-branching/internal/interfacestyle"
)

func main() {
	amount := 12000

	fmt.Println("=== if/switch スタイル ===")
	showResult(ifstyle.Calculate(common.PaymentMethodCreditCard, amount))
	showResult(ifstyle.Calculate(common.PaymentMethodPayPay, amount))
	showResult(ifstyle.Calculate(common.PaymentMethodBankTransfer, amount))

	fmt.Println("=== interface スタイル ===")
	showResult(interfacestyle.Calculate(common.PaymentMethodCreditCard, amount))
	showResult(interfacestyle.Calculate(common.PaymentMethodPayPay, amount))
	showResult(interfacestyle.Calculate(common.PaymentMethodBankTransfer, amount))
}

func showResult(res common.CalculationResult, err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("fee=%d total=%d\n", res.Fee, res.Total)
}
