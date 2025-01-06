package payment

import (
	"errors"
	"testing"
)




type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestPayPalProcessorValidatePayment(t *testing.T) {
	type testCase struct {
		desc     string
		payment  Payment
		expected error
	}

	tests := []testCase{
		{
			desc: "Scenario 1: Validate Positive Payment Amount",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expected: nil,
		},
		{
			desc: "Scenario 2: Validate Payment with Zero Amount",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expected: errors.New("amount must be positive"),
		},
		{
			desc: "Scenario 3: Validate Payment with Negative Amount",
			payment: Payment{
				Amount:   -50.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expected: errors.New("amount must be positive"),
		},
		{
			desc: "Scenario 4: Validate Payment with Empty Currency",
			payment: Payment{
				Amount:   50.0,
				Currency: "",
				Method:   "PayPal",
			},
			expected: errors.New("currency must be specified"),
		},
		{
			desc: "Scenario 5: Validate Payment with Valid Amount and Currency",
			payment: Payment{
				Amount:   200.0,
				Currency: "EUR",
				Method:   "PayPal",
			},
			expected: nil,
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := processor.ValidatePayment(tc.payment)

			if err != nil && tc.expected == nil {
				t.Errorf("expected no error, got %v", err)
			} else if err == nil && tc.expected != nil {
				t.Errorf("expected error %v, got no error", tc.expected)
			} else if err != nil && tc.expected != nil && err.Error() != tc.expected.Error() {
				t.Errorf("expected error %v, got %v", tc.expected, err)
			}

			t.Logf("Test '%s' completed successfully", tc.desc)
		})
	}
}
