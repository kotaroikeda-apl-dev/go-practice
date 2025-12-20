package ifstyle

import (
	"fmt"

	"github.com/example/payment-branching/internal/common"
)

func Calculate(method common.PaymentMethod, amount int) (common.CalculationResult, error) {
	if err := common.ValidateAmount(amount); err != nil {
		return common.CalculationResult{}, err
	}

	switch method {
	case common.PaymentMethodCreditCard:
		fee := common.PercentCeil(amount, 3.0)
		return common.CalculationResult{Fee: fee, Total: amount + fee}, nil
	case common.PaymentMethodPayPay:
		if err := common.PayPayLimit(amount); err != nil {
			return common.CalculationResult{}, err
		}
		fee := common.PercentCeil(amount, 1.5)
		return common.CalculationResult{Fee: fee, Total: amount + fee}, nil
	case common.PaymentMethodBankTransfer:
		fee := common.BankTransferFee()
		if amount < fee {
			return common.CalculationResult{}, common.ErrBankTotalTooLow
		}
		return common.CalculationResult{Fee: fee, Total: amount + fee}, nil
	default:
		return common.CalculationResult{}, fmt.Errorf("%w: %s", common.ErrUnsupportedMethod, method)
	}
}

