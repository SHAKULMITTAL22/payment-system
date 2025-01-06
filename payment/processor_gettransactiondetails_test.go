package payment

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)







func TestPayPalProcessorGetTransactionDetails(t *testing.T) {
	type testCase struct {
		name           string
		transactionID  string
		expectedError  error
		expectedOutput string
	}

	testCases := []testCase{
		{
			name:           "Valid Transaction ID",
			transactionID:  "12345",
			expectedError:  nil,
			expectedOutput: "Details of Transaction ID: 12345",
		},
		{
			name:           "Empty Transaction ID",
			transactionID:  "",
			expectedError:  errors.New("invalid transaction ID"),
			expectedOutput: "",
		},
		{
			name:           "Randomly Generated Transaction ID",
			transactionID:  strconv.Itoa(rand.Intn(100000)),
			expectedError:  nil,
			expectedOutput: fmt.Sprintf("Details of Transaction ID: %s", strconv.Itoa(rand.Intn(100000))),
		},
		{
			name:           "Transaction ID with Special Characters",
			transactionID:  "#Special@ID!",
			expectedError:  nil,
			expectedOutput: "Details of Transaction ID: #Special@ID!",
		},
		{
			name:           "Very Long Transaction ID",
			transactionID:  generateLongString(1000),
			expectedError:  nil,
			expectedOutput: fmt.Sprintf("Details of Transaction ID: %s", generateLongString(1000)),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)

			p := PayPalProcessor{}
			output, err := p.GetTransactionDetails(tc.transactionID)

			if err != nil && tc.expectedError == nil {
				t.Errorf("Expected no error, but got %v", err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("Expected error %v, but got none", tc.expectedError)
			} else if err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
			}

			if output != tc.expectedOutput {
				t.Errorf("Expected output %v, but got %v", tc.expectedOutput, output)
			}
		})
	}
}
func generateLongString(length int) string {
	const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}
