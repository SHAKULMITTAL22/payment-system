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
func TestPayPalProcessorRefund(t *testing.T) {
	type testCase struct {
		name          string
		transactionID string
		amount        float64
		expectedError error
	}

	tests := []testCase{
		{
			name:          "Valid transaction ID and positive amount",
			transactionID: "TXN12345",
			amount:        100.0,
			expectedError: nil,
		},
		{
			name:          "Empty transaction ID",
			transactionID: "",
			amount:        100.0,
			expectedError: errors.New("invalid transaction ID"),
		},
		{
			name:          "Negative refund amount",
			transactionID: "TXN12345",
			amount:        -50.0,
			expectedError: errors.New("refund amount must be positive"),
		},
		{
			name:          "Zero refund amount",
			transactionID: "TXN12345",
			amount:        0.0,
			expectedError: errors.New("refund amount must be positive"),
		},
		{
			name:          "Boundary case for transaction ID length",
			transactionID: "TXN" + strconv.Itoa(rand.Intn(1000000000)),
			amount:        100.0,
			expectedError: nil,
		},
		{
			name:          "Random valid inputs to test randomness of refund ID",
			transactionID: "TXN" + strconv.Itoa(rand.Intn(1000000)),
			amount:        50.0,
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			processor := PayPalProcessor{}

			refundID, err := processor.Refund(tc.transactionID, tc.amount)

			if tc.expectedError != nil {
				if err == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
				}
				if refundID != "" {
					t.Errorf("Expected empty refund ID, got: %v", refundID)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if refundID == "" {
					t.Errorf("Expected non-empty refund ID, got: %v", refundID)
				}
			}
		})
	}

	t.Run("Randomness of refund ID", func(t *testing.T) {
		processor := PayPalProcessor{}
		transactionID := "TXN12345"
		amount := 100.0
		refundIDs := make(map[string]bool)

		for i := 0; i < 10; i++ {
			refundID, err := processor.Refund(transactionID, amount)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if _, exists := refundIDs[refundID]; exists {
				t.Errorf("Duplicate refund ID found: %v", refundID)
			}
			refundIDs[refundID] = true
		}
	})
}
