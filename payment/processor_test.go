package payment

import (
	"testing"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)








/*
ROOST_METHOD_HASH=SupportsCurrency_875c78ed3e
ROOST_METHOD_SIG_HASH=SupportsCurrency_d2d80e5419


 */
func TestPayPalProcessorSupportsCurrency(t *testing.T) {
	type testCase struct {
		name           string
		currency       string
		expectedResult bool
	}

	testCases := []testCase{
		{
			name:           "Supported Currency - USD",
			currency:       "USD",
			expectedResult: true,
		},
		{
			name:           "Unsupported Currency - JPY",
			currency:       "JPY",
			expectedResult: false,
		},
		{
			name:           "Empty Currency String",
			currency:       "",
			expectedResult: false,
		},
		{
			name:           "Lowercase Currency Code - usd",
			currency:       "usd",
			expectedResult: false,
		},
		{
			name:           "Numeric Currency Code - 123",
			currency:       "123",
			expectedResult: false,
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := processor.SupportsCurrency(tc.currency)
			if actualResult != tc.expectedResult {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expectedResult, actualResult)
			} else {
				t.Logf("Test %s succeeded: currency %s correctly returned %v", tc.name, tc.currency, actualResult)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidatePayment_cd98909ed8
ROOST_METHOD_SIG_HASH=ValidatePayment_cb0f85252d


 */
func TestPayPalProcessorValidatePayment(t *testing.T) {
	tests := []struct {
		name          string
		payment       Payment
		expectedError error
	}{
		{
			name: "Scenario 1: Validate Positive Payment Amount",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedError: nil,
		},
		{
			name: "Scenario 2: Validate Payment with Zero Amount",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedError: errors.New("amount must be positive"),
		},
		{
			name: "Scenario 3: Validate Payment with Negative Amount",
			payment: Payment{
				Amount:   -50.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedError: errors.New("amount must be positive"),
		},
		{
			name: "Scenario 4: Validate Payment with Empty Currency",
			payment: Payment{
				Amount:   100.0,
				Currency: "",
				Method:   "CreditCard",
			},
			expectedError: errors.New("currency must be specified"),
		},
		{
			name: "Scenario 5: Validate Payment with Valid Amount and Currency",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := PayPalProcessor{}
			err := processor.ValidatePayment(tt.payment)

			if err != nil && tt.expectedError == nil {
				t.Errorf("Expected no error, but got: %v", err)
				t.Log("Failure: Expected no error for valid payment but received an error indicating a logic issue.")
			} else if err == nil && tt.expectedError != nil {
				t.Errorf("Expected error: %v, but got nil", tt.expectedError)
				t.Log("Failure: Expected an error indicating invalid payment, but received none, showing a validation oversight.")
			} else if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedError, err)
				t.Log("Failure: Error message mismatch, indicating possible incorrect error handling in ValidatePayment.")
			} else {
				t.Logf("Success: %s", tt.name)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=GetTransactionDetails_d67406f337
ROOST_METHOD_SIG_HASH=GetTransactionDetails_9463e70b0f


 */
func TestPayPalProcessorGetTransactionDetails(t *testing.T) {

	tests := []struct {
		name           string
		transactionID  string
		expectedError  error
		expectedOutput string
	}{
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
			transactionID:  strconv.Itoa(rand.Intn(1000000)),
			expectedError:  nil,
			expectedOutput: fmt.Sprintf("Details of Transaction ID: %s", strconv.Itoa(rand.Intn(1000000))),
		},
		{
			name:           "Transaction ID with Special Characters",
			transactionID:  "#@!TransactionID",
			expectedError:  nil,
			expectedOutput: "Details of Transaction ID: #@!TransactionID",
		},
		{
			name:           "Very Long Transaction ID",
			transactionID:  generateLongString(1000),
			expectedError:  nil,
			expectedOutput: fmt.Sprintf("Details of Transaction ID: %s", generateLongString(1000)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := PayPalProcessor{}

			result, err := p.GetTransactionDetails(tt.transactionID)

			if tt.expectedError != nil {
				if err == nil || err.Error() != tt.expectedError.Error() {
					t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
				}
				t.Logf("Test %s: Passed with expected error: %v", tt.name, err)
			} else {
				if result != tt.expectedOutput || err != nil {
					t.Errorf("expected output: %v, got: %v, expected error: %v, got: %v", tt.expectedOutput, result, tt.expectedError, err)
				}
				t.Logf("Test %s: Passed with expected output: %v", tt.name, result)
			}
		})
	}
}

func generateLongString(length int) string {

	longString := ""
	for i := 0; i < length; i++ {
		longString += "A"
	}
	return longString
}


/*
ROOST_METHOD_HASH=Process_c114677531
ROOST_METHOD_SIG_HASH=Process_893bbedcfe


 */
func TestPayPalProcessorProcess(t *testing.T) {
	type testCase struct {
		name          string
		payment       Payment
		expectedID    string
		expectedError error
	}

	tests := []testCase{
		{
			name: "Successful Payment Processing with PayPal",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:    "non-empty",
			expectedError: nil,
		},
		{
			name: "Unsupported Payment Method",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedID:    "",
			expectedError: errors.New("unsupported payment method"),
		},
		{
			name: "Payment with Zero Amount",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:    "",
			expectedError: errors.New("amount must be greater than zero"),
		},
		{
			name: "Payment with Negative Amount",
			payment: Payment{
				Amount:   -10.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:    "",
			expectedError: errors.New("amount must be greater than zero"),
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			transactionID, err := processor.Process(tc.payment)

			if tc.expectedError != nil {
				if err == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("expected error %v, got %v", tc.expectedError, err)
				}
				if transactionID != tc.expectedID {
					t.Errorf("expected transaction ID %v, got %v", tc.expectedID, transactionID)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if transactionID == "" {
					t.Error("expected a non-empty transaction ID, got empty")
				}
			}

			t.Logf("Test '%s' executed successfully", tc.name)
		})
	}

	t.Run("Random Transaction ID Generation", func(t *testing.T) {
		payment1 := Payment{
			Amount:   50.0,
			Currency: "USD",
			Method:   "PayPal",
		}
		payment2 := Payment{
			Amount:   75.0,
			Currency: "USD",
			Method:   "PayPal",
		}

		transactionID1, err1 := processor.Process(payment1)
		transactionID2, err2 := processor.Process(payment2)

		if err1 != nil || err2 != nil {
			t.Fatalf("expected no error, got %v and %v", err1, err2)
		}

		if transactionID1 == "" || transactionID2 == "" {
			t.Fatal("expected non-empty transaction IDs for both payments")
		}

		if transactionID1 == transactionID2 {
			t.Error("expected different transaction IDs, got the same")
		}

		t.Log("Random Transaction ID Generation test executed successfully")
	})
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
		expectedError error
		expectRefund  bool
	}

	testCases := []testCase{
		{
			name:          "Scenario 1: Valid transaction ID and positive amount",
			transactionID: "TXN123456",
			amount:        100.0,
			expectedError: nil,
			expectRefund:  true,
		},
		{
			name:          "Scenario 2: Empty transaction ID",
			transactionID: "",
			amount:        100.0,
			expectedError: errors.New("invalid transaction ID"),
			expectRefund:  false,
		},
		{
			name:          "Scenario 3: Negative refund amount",
			transactionID: "TXN123456",
			amount:        -50.0,
			expectedError: errors.New("refund amount must be positive"),
			expectRefund:  false,
		},
		{
			name:          "Scenario 4: Zero refund amount",
			transactionID: "TXN123456",
			amount:        0.0,
			expectedError: errors.New("refund amount must be positive"),
			expectRefund:  false,
		},
		{
			name:          "Scenario 5: Boundary case for transaction ID length",
			transactionID: "TXN" + strconv.Itoa(rand.Intn(1000000)),
			amount:        50.0,
			expectedError: nil,
			expectRefund:  true,
		},
		{
			name:          "Scenario 6: Random valid inputs for refund ID uniqueness",
			transactionID: "TXN123456",
			amount:        75.0,
			expectedError: nil,
			expectRefund:  true,
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			refundID, err := processor.Refund(tc.transactionID, tc.amount)
			if tc.expectedError != nil {
				if err == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
				}
				if tc.expectRefund && refundID != "" {
					t.Errorf("Expected empty refund ID, got: %s", refundID)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if tc.expectRefund && refundID == "" {
					t.Errorf("Expected non-empty refund ID, got empty")
				}
			}
			t.Logf("Test %s succeeded", tc.name)
		})

		if tc.name == "Scenario 6: Random valid inputs for refund ID uniqueness" && tc.expectRefund {
			refundIDs := make(map[string]bool)
			for i := 0; i < 10; i++ {
				refundID, err := processor.Refund(tc.transactionID, tc.amount)
				if err != nil {
					t.Fatalf("Unexpected error during uniqueness check: %v", err)
				}
				if _, exists := refundIDs[refundID]; exists {
					t.Errorf("Duplicate refund ID found: %s", refundID)
				} else {
					refundIDs[refundID] = true
				}
			}
			t.Logf("Unique refund IDs generated for Scenario 6")
		}
	}
}

