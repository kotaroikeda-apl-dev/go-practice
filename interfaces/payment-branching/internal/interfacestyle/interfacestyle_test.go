package interfacestyle

import (
	"errors"
	"testing"

	"github.com/example/payment-branching/internal/common"
)

func TestCalculateConcrete(t *testing.T) {
	tests := []struct {
		name    string
		calc    PaymentMethod
		amount  int
		wantFee int
		wantTot int
		wantErr error
	}{
		{"credit card", CreditCard{}, 1000, 30, 1030, nil},
		{"paypay", PayPay{}, 1000, 15, 1015, nil},
		{"paypay limit", PayPay{}, 1_000_001, 0, 0, common.ErrPayPayTooHigh},
		{"bank", BankTransfer{}, 1000, 300, 1300, nil},
		{"bank too small", BankTransfer{}, 200, 0, 0, common.ErrBankTotalTooLow},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.calc.Calculate(tt.amount)

			if tt.wantErr != nil {
				if err == nil || !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}

			if got.Fee != tt.wantFee || got.Total != tt.wantTot {
				t.Fatalf("fee=%d total=%d want fee=%d total=%d", got.Fee, got.Total, tt.wantFee, tt.wantTot)
			}
		})
	}
}

func TestCalculateDispatcher(t *testing.T) {
	got, err := Calculate(common.PaymentMethodCreditCard, 2000)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got.Fee != 60 || got.Total != 2060 {
		t.Fatalf("unexpected result: %+v", got)
	}
}

type mockCalculator struct {
	fn     func(int) (common.CalculationResult, error)
	called bool
}

func (m *mockCalculator) Calculate(amount int) (common.CalculationResult, error) {
	m.called = true
	return m.fn(amount)
}

func TestMockStyle(t *testing.T) {
	m := &mockCalculator{
		fn: func(amount int) (common.CalculationResult, error) {
			if amount != 123 {
				return common.CalculationResult{}, errors.New("unexpected amount")
			}
			return common.CalculationResult{Fee: 1, Total: 124}, nil
		},
	}
	var _ PaymentMethod = m

	got, err := m.Calculate(123)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !m.called {
		t.Fatalf("mock was not called")
	}
	if got.Total != 124 || got.Fee != 1 {
		t.Fatalf("unexpected result: %+v", got)
	}
}



