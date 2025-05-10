package payment

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestPayPalProcessorGetTransactionDetails(t *testing.T) {

	processor := PayPalProcessor{}

	tests := []struct {
		name          string
		transactionID string
		expectedError error
		shouldContain string
	}{
		{
			name:          "Valid Transaction ID",
			transactionID: "12345ABC",
			expectedError: nil,
			shouldContain: "Details of Transaction ID: 12345ABC",
		},
		{
			name:          "Empty Transaction ID",
			transactionID: "",
			expectedError: errors.New("invalid transaction ID"),
			shouldContain: "",
		},
		{
			name: "Randomly Generated Transaction ID",
			transactionID: func() string {
				rand.Seed(time.Now().UnixNano())
				return strconv.Itoa(rand.Intn(1000000))
			}(),
			expectedError: nil,
			shouldContain: "Details of Transaction ID:",
		},
		{
			name:          "Transaction ID with Special Characters",
			transactionID: "#$%@!123",
			expectedError: nil,
			shouldContain: "Details of Transaction ID: #$%@!123",
		},
		{
			name:          "Very Long Transaction ID",
			transactionID: strings.Repeat("A", 1000),
			expectedError: nil,
			shouldContain: fmt.Sprintf("Details of Transaction ID: %s", strings.Repeat("A", 1000)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)

			result, err := processor.GetTransactionDetails(tc.transactionID)

			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("Expected error: %v, but got nil", tc.expectedError)
			}

			if !strings.Contains(result, tc.shouldContain) {
				t.Errorf("Expected result to contain: %s, but got: %s", tc.shouldContain, result)
			}

			t.Logf("Success: %s", tc.name)
		})
	}

}
