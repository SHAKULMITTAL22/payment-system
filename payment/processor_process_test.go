package payment

import (
	"errors"
	"math/rand"
	"strconv"
	"testing"
)







func TestPayPalProcessorProcess(t *testing.T) {

	processor := PayPalProcessor{}

	tests := []struct {
		name           string
		payment        Payment
		expectedErr    error
		expectTxnID    bool
		expectedErrMsg string
	}{
		{
			name: "Successful Payment Processing with PayPal",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedErr: nil,
			expectTxnID: true,
		},
		{
			name: "Unsupported Payment Method",
			payment: Payment{
				Amount:   50.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedErr:    errors.New("unsupported payment method"),
			expectTxnID:    false,
			expectedErrMsg: "unsupported payment method",
		},
		{
			name: "Payment with Zero Amount",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedErr:    errors.New("amount must be greater than zero"),
			expectTxnID:    false,
			expectedErrMsg: "amount must be greater than zero",
		},
		{
			name: "Payment with Negative Amount",
			payment: Payment{
				Amount:   -10.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedErr:    errors.New("amount must be greater than zero"),
			expectTxnID:    false,
			expectedErrMsg: "amount must be greater than zero",
		},
		{
			name: "Random Transaction ID Generation",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedErr: nil,
			expectTxnID: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txnID, err := processor.Process(tt.payment)

			if err != nil && err.Error() != tt.expectedErrMsg {
				t.Errorf("expected error '%v', got '%v'", tt.expectedErr, err)
			}

			if tt.expectTxnID && txnID == "" {
				t.Error("expected a non-empty transaction ID, got empty")
			} else if !tt.expectTxnID && txnID != "" {
				t.Error("expected an empty transaction ID, got non-empty")
			}

			t.Logf("Test scenario: %s completed", tt.name)
		})
	}
}
