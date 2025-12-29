package ifstyle

import (
	"errors"
	"testing"

	"github.com/example/payment-branching/internal/common"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		method  common.PaymentMethod
		amount  int
		wantFee int
		wantTot int
		wantErr error
	}{
		{"credit card basic", common.PaymentMethodCreditCard, 1000, 30, 1030, nil},
		{"paypay basic", common.PaymentMethodPayPay, 1000, 15, 1015, nil},
		{"bank basic", common.PaymentMethodBankTransfer, 1000, 300, 1300, nil},
		{"paypay upper limit", common.PaymentMethodPayPay, 1_000_001, 0, 0, common.ErrPayPayTooHigh},
		{"too low", common.PaymentMethodCreditCard, 0, 0, 0, common.ErrAmountTooLow},
		{"bank too small", common.PaymentMethodBankTransfer, 200, 0, 0, common.ErrBankTotalTooLow},
		{"unsupported", common.PaymentMethod("unknown"), 100, 0, 0, common.ErrUnsupportedMethod},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.method, tt.amount)

			if tt.wantErr != nil {
				if err == nil || !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.Fee != tt.wantFee || got.Total != tt.wantTot {
				t.Fatalf("got fee=%d total=%d, want fee=%d total=%d", got.Fee, got.Total, tt.wantFee, tt.wantTot)
			}
		})
	}
}



