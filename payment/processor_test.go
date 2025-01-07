package payment

import (
	"testing"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"github.com/stretchr/testify/assert"
)








/*
ROOST_METHOD_HASH=SupportsCurrency_875c78ed3e
ROOST_METHOD_SIG_HASH=SupportsCurrency_d2d80e5419


 */
func TestPayPalProcessorSupportsCurrency(t *testing.T) {
	type testCase struct {
		desc     string
		currency string
		expected bool
	}

	testCases := []testCase{
		{
			desc:     "Scenario 1: Check Support for a Supported Currency",
			currency: "USD",
			expected: true,
		},
		{
			desc:     "Scenario 2: Check Support for an Unsupported Currency",
			currency: "JPY",
			expected: false,
		},
		{
			desc:     "Scenario 3: Check Support for an Empty Currency String",
			currency: "",
			expected: false,
		},
		{
			desc:     "Scenario 4: Check Support for a Lowercase Currency Code",
			currency: "usd",
			expected: false,
		},
		{
			desc:     "Scenario 5: Check Support for a Numeric Currency Code",
			currency: "123",
			expected: false,
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			result := processor.SupportsCurrency(tc.currency)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.desc, tc.expected, result)
			} else {
				t.Logf("Passed %s: expected %v, got %v", tc.desc, tc.expected, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidatePayment_cd98909ed8
ROOST_METHOD_SIG_HASH=ValidatePayment_cb0f85252d


 */
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


/*
ROOST_METHOD_HASH=GetTransactionDetails_d67406f337
ROOST_METHOD_SIG_HASH=GetTransactionDetails_9463e70b0f


 */
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


/*
ROOST_METHOD_HASH=Process_c114677531
ROOST_METHOD_SIG_HASH=Process_893bbedcfe


 */
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


/*
ROOST_METHOD_HASH=Refund_e2af3d6bf2
ROOST_METHOD_SIG_HASH=Refund_4cecc33874


 */
func TestPayPalProcessorRefund(t *testing.T) {
	type testCase struct {
		name          string
		transactionID string
		amount        float64
		expectError   bool
		expectEmptyID bool
	}

	testCases := []testCase{
		{
			name:          "Valid Transaction ID and Positive Amount",
			transactionID: "validID123",
			amount:        100.0,
			expectError:   false,
			expectEmptyID: false,
		},
		{
			name:          "Empty Transaction ID",
			transactionID: "",
			amount:        100.0,
			expectError:   true,
			expectEmptyID: true,
		},
		{
			name:          "Negative Refund Amount",
			transactionID: "validID123",
			amount:        -50.0,
			expectError:   true,
			expectEmptyID: true,
		},
		{
			name:          "Zero Refund Amount",
			transactionID: "validID123",
			amount:        0.0,
			expectError:   true,
			expectEmptyID: true,
		},
		{
			name:          "Boundary Case for Transaction ID Length",
			transactionID: "ID" + string(make([]rune, 255)),
			amount:        100.0,
			expectError:   false,
			expectEmptyID: false,
		},
	}

	payPalProcessor := PayPalProcessor{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			refundID, err := payPalProcessor.Refund(tc.transactionID, tc.amount)

			if tc.expectError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
			}

			if tc.expectEmptyID {
				assert.Empty(t, refundID, "Expected refund ID to be empty")
			} else {
				assert.NotEmpty(t, refundID, "Expected refund ID to be non-empty")
			}

			t.Logf("Test case '%s' completed successfully", tc.name)
		})
	}

	t.Run("Randomness of Refund ID", func(t *testing.T) {
		transactionID := "randomID123"
		amount := 50.0
		refundIDs := make(map[string]bool)

		for i := 0; i < 10; i++ {
			refundID, err := payPalProcessor.Refund(transactionID, amount)
			assert.NoError(t, err, "Expected no error but got one")
			assert.NotEmpty(t, refundID, "Expected refund ID to be non-empty")

			if _, exists := refundIDs[refundID]; exists {
				t.Errorf("Duplicate refund ID found: %s", refundID)
			}
			refundIDs[refundID] = true
		}

		t.Log("All refund IDs were unique across multiple calls")
	})
}

