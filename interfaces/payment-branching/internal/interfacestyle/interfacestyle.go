package interfacestyle

import (
	"fmt"

	"github.com/example/payment-branching/internal/common"
)

type PaymentMethod interface {
	Calculate(amount int) (common.CalculationResult, error)
}

type CreditCard struct{}

func (CreditCard) Calculate(amount int) (common.CalculationResult, error) {
	if err := common.ValidateAmount(amount); err != nil {
		return common.CalculationResult{}, err
	}
	fee := common.PercentCeil(amount, 3.0)
	return common.CalculationResult{Fee: fee, Total: amount + fee}, nil
}

type PayPay struct{}

func (PayPay) Calculate(amount int) (common.CalculationResult, error) {
	if err := common.ValidateAmount(amount); err != nil {
		return common.CalculationResult{}, err
	}
	if err := common.PayPayLimit(amount); err != nil {
		return common.CalculationResult{}, err
	}
	fee := common.PercentCeil(amount, 1.5)
	return common.CalculationResult{Fee: fee, Total: amount + fee}, nil
}

type BankTransfer struct{}

func (BankTransfer) Calculate(amount int) (common.CalculationResult, error) {
	if err := common.ValidateAmount(amount); err != nil {
		return common.CalculationResult{}, err
	}
	fee := common.BankTransferFee()
	if amount < fee {
		return common.CalculationResult{}, common.ErrBankTotalTooLow
	}
	return common.CalculationResult{Fee: fee, Total: amount + fee}, nil
}

func Calculate(method common.PaymentMethod, amount int) (common.CalculationResult, error) {
	calculators := map[common.PaymentMethod]PaymentMethod{
		common.PaymentMethodCreditCard:   CreditCard{},
		common.PaymentMethodPayPay:       PayPay{},
		common.PaymentMethodBankTransfer: BankTransfer{},
	}

	calc, ok := calculators[method]
	if !ok {
		return common.CalculationResult{}, fmt.Errorf("%w: %s", common.ErrUnsupportedMethod, method)
	}
	return calc.Calculate(amount)
}

