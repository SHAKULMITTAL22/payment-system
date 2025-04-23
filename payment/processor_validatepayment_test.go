package payment

import (
	"errors"
	"testing"
)







func TestPayPalProcessorValidatePayment(t *testing.T) {

	testCases := []struct {
		name          string
		payment       Payment
		expectedError error
	}{
		{
			name: "Positive Payment Amount",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedError: nil,
		},
		{
			name: "Zero Payment Amount",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedError: errors.New("amount must be positive"),
		},
		{
			name: "Negative Payment Amount",
			payment: Payment{
				Amount:   -50.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedError: errors.New("amount must be positive"),
		},
		{
			name: "Empty Currency",
			payment: Payment{
				Amount:   100.0,
				Currency: "",
				Method:   "PayPal",
			},
			expectedError: errors.New("currency must be specified"),
		},
		{
			name: "Valid Amount and Currency",
			payment: Payment{
				Amount:   200.0,
				Currency: "EUR",
				Method:   "PayPal",
			},
			expectedError: nil,
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)

			err := processor.ValidatePayment(tc.payment)

			if err != nil && tc.expectedError == nil {
				t.Errorf("Expected no error, but got: %v", err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("Expected error: %v, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedError, err)
			}

			if err == nil {
				t.Logf("Success: No error as expected for test case: %s", tc.name)
			} else {
				t.Logf("Failure: Expected error: %v, got: %v for test case: %s", tc.expectedError, err, tc.name)
			}
		})
	}
}
