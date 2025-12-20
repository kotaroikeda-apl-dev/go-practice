package common

import (
	"errors"
	"fmt"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodPayPay       PaymentMethod = "paypay"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)

const (
	payPayMaxAmount = 1_000_000
	bankTransferFee = 300
)

type CalculationResult struct {
	Fee   int
	Total int
}

var (
	ErrUnsupportedMethod = errors.New("unsupported payment method")
	ErrAmountTooLow      = errors.New("amount must be >= 1")
	ErrPayPayTooHigh     = fmt.Errorf("paypay amount must be <= %d", payPayMaxAmount)
	ErrBankTotalTooLow   = errors.New("bank transfer total must be >= fee")
)

func PercentCeil(amount int, percent float64) int {
	fee := float64(amount) * percent / 100.0
	if fee == float64(int(fee)) {
		return int(fee)
	}
	return int(fee) + 1
}

func ValidateAmount(amount int) error {
	if amount < 1 {
		return ErrAmountTooLow
	}
	return nil
}

func PayPayLimit(amount int) error {
	if amount > payPayMaxAmount {
		return ErrPayPayTooHigh
	}
	return nil
}

func BankTransferFee() int {
	return bankTransferFee
}

