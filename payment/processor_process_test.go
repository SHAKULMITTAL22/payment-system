package payment

import (
	"errors"
	"math/rand"
	"strconv"
	"testing"
)




type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}

func TestPayPalProcessorProcess(t *testing.T) {
	processor := PayPalProcessor{}

	tests := []struct {
		name          string
		input         Payment
		expectedID    string
		expectedError error
	}{
		{
			name: "Successful Payment Processing with PayPal",
			input: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:    "non-empty",
			expectedError: nil,
		},
		{
			name: "Unsupported Payment Method",
			input: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedID:    "",
			expectedError: errors.New("unsupported payment method"),
		},
		{
			name: "Payment with Zero Amount",
			input: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:    "",
			expectedError: errors.New("amount must be greater than zero"),
		},
		{
			name: "Payment with Negative Amount",
			input: Payment{
				Amount:   -5.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:    "",
			expectedError: errors.New("amount must be greater than zero"),
		},
		{
			name: "Random Transaction ID Generation",
			input: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:    "unique",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionID, err := processor.Process(tt.input)

			if tt.expectedID == "non-empty" && transactionID == "" {
				t.Errorf("expected a non-empty transaction ID, got empty")
			}

			if tt.expectedID == "unique" {
				transactionID2, err2 := processor.Process(tt.input)
				if err2 != nil {
					t.Fatalf("unexpected error: %v", err2)
				}
				if transactionID == transactionID2 {
					t.Errorf("expected unique transaction IDs, got identical: %s", transactionID)
				}
			}

			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) || (err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("expected error '%v', got '%v'", tt.expectedError, err)
			}

			if err != nil {
				t.Logf("Operation resulted in expected error: %v", err)
			} else {
				t.Logf("Operation successful with transaction ID: %s", transactionID)
			}
		})
	}
}
