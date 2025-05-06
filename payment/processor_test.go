package payment

import (
	"testing"
	"errors"
	"math"
	"fmt"
	"strconv"
)








/*
ROOST_METHOD_HASH=SupportsCurrency_875c78ed3e
ROOST_METHOD_SIG_HASH=SupportsCurrency_d2d80e5419

FUNCTION_DEF=func (p PayPalProcessor) SupportsCurrency(currency string) bool 

 */
func TestPayPalProcessorSupportsCurrency(t *testing.T) {

	tests := []struct {
		name     string
		currency string
		expected bool
	}{
		{
			name:     "Supported Currency USD",
			currency: "USD",
			expected: true,
		},
		{
			name:     "Unsupported Currency JPY",
			currency: "JPY",
			expected: false,
		},
		{
			name:     "Empty Currency String",
			currency: "",
			expected: false,
		},
		{
			name:     "Case Sensitivity Check usd",
			currency: "usd",
			expected: false,
		},
		{
			name:     "Currency with Special Characters US$",
			currency: "US$",
			expected: false,
		},
		{
			name:     "Numeric Currency Code 840",
			currency: "840",
			expected: false,
		},
		{
			name:     "Valid Currency with Leading/Trailing Spaces",
			currency: " USD ",
			expected: false,
		},
	}

	processor := PayPalProcessor{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processor.SupportsCurrency(tt.currency)
			if result != tt.expected {
				t.Errorf("SupportsCurrency(%q) = %v; want %v", tt.currency, result, tt.expected)
			} else {
				t.Logf("Test '%s' passed: SupportsCurrency(%q) = %v", tt.name, tt.currency, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidatePayment_cd98909ed8
ROOST_METHOD_SIG_HASH=ValidatePayment_cb0f85252d

FUNCTION_DEF=func (p PayPalProcessor) ValidatePayment(payment Payment) error 

 */
func TestPayPalProcessorValidatePayment(t *testing.T) {
	type testCase struct {
		desc     string
		payment  Payment
		expected error
	}

	tests := []testCase{
		{
			desc: "Valid Payment",
			payment: Payment{
				Amount:   100.00,
				Currency: "USD",
				Method:   "PayPal",
			},
			expected: nil,
		},
		{
			desc: "Negative Payment Amount",
			payment: Payment{
				Amount:   -50.00,
				Currency: "USD",
				Method:   "PayPal",
			},
			expected: errors.New("amount must be positive"),
		},
		{
			desc: "Zero Payment Amount",
			payment: Payment{
				Amount:   0.00,
				Currency: "USD",
				Method:   "PayPal",
			},
			expected: errors.New("amount must be positive"),
		},
		{
			desc: "Missing Currency",
			payment: Payment{
				Amount:   100.00,
				Currency: "",
				Method:   "PayPal",
			},
			expected: errors.New("currency must be specified"),
		},
		{
			desc: "Valid Payment with Maximum Float Amount",
			payment: Payment{
				Amount:   math.MaxFloat64,
				Currency: "USD",
				Method:   "PayPal",
			},
			expected: nil,
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			t.Log("Running test:", tc.desc)
			err := processor.ValidatePayment(tc.payment)

			if tc.expected == nil && err != nil {
				t.Errorf("expected no error, but got: %v", err)
			}

			if tc.expected != nil && err == nil {
				t.Errorf("expected error: %v, but got none", tc.expected)
			}

			if tc.expected != nil && err != nil && tc.expected.Error() != err.Error() {
				t.Errorf("expected error: %v, but got: %v", tc.expected, err)
			}

			t.Logf("Test %s passed", tc.desc)
		})
	}
}


/*
ROOST_METHOD_HASH=GetTransactionDetails_d67406f337
ROOST_METHOD_SIG_HASH=GetTransactionDetails_9463e70b0f

FUNCTION_DEF=func (p PayPalProcessor) GetTransactionDetails(transactionID string) (string, error) 

 */
func TestPayPalProcessorGetTransactionDetails(t *testing.T) {

	tests := []struct {
		name            string
		transactionID   string
		expectedError   error
		expectedDetails string
	}{
		{
			name:            "Valid Transaction ID",
			transactionID:   "123ABC",
			expectedError:   nil,
			expectedDetails: "Details of Transaction ID: 123ABC",
		},
		{
			name:            "Empty Transaction ID",
			transactionID:   "",
			expectedError:   errors.New("invalid transaction ID"),
			expectedDetails: "",
		},
		{
			name:            "Special Characters in Transaction ID",
			transactionID:   "#$%^&*",
			expectedError:   nil,
			expectedDetails: "Details of Transaction ID: #$%^&*",
		},
		{
			name:            "Numeric Transaction ID",
			transactionID:   "1234567890",
			expectedError:   nil,
			expectedDetails: "Details of Transaction ID: 1234567890",
		},
		{
			name:            "Very Long Transaction ID",
			transactionID:   "a" + fmt.Sprintf("%0*d", 1000, 0),
			expectedError:   nil,
			expectedDetails: "Details of Transaction ID: " + "a" + fmt.Sprintf("%0*d", 1000, 0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			processor := PayPalProcessor{}

			details, err := processor.GetTransactionDetails(test.transactionID)

			if test.expectedError != nil && err == nil {
				t.Errorf("expected error %v, got nil", test.expectedError)
			} else if test.expectedError == nil && err != nil {
				t.Errorf("expected no error, got %v", err)
			} else if err != nil && err.Error() != test.expectedError.Error() {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}

			if details != test.expectedDetails {
				t.Errorf("expected details %q, got %q", test.expectedDetails, details)
			}

			t.Logf("Test scenario '%s' executed successfully", test.name)
		})
	}
}


/*
ROOST_METHOD_HASH=Process_c114677531
ROOST_METHOD_SIG_HASH=Process_893bbedcfe

FUNCTION_DEF=func (p PayPalProcessor) Process(payment Payment) (string, error) 

 */
func TestPayPalProcessorProcess(t *testing.T) {
	type testCase struct {
		description   string
		payment       Payment
		expectedErr   error
		expectedTxnID string
	}

	testCases := []testCase{
		{
			description: "Process Payment with Valid PayPal Method and Positive Amount",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedErr:   nil,
			expectedTxnID: "TXN-",
		},
		{
			description: "Process Payment with Unsupported Method",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedErr:   errors.New("unsupported payment method"),
			expectedTxnID: "",
		},
		{
			description: "Process Payment with Zero Amount",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedErr:   errors.New("amount must be greater than zero"),
			expectedTxnID: "",
		},
		{
			description: "Process Payment with Negative Amount",
			payment: Payment{
				Amount:   -50.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedErr:   errors.New("amount must be greater than zero"),
			expectedTxnID: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			processor := PayPalProcessor{}
			txnID, err := processor.Process(tc.payment)

			if tc.expectedErr != nil {
				if err == nil || err.Error() != tc.expectedErr.Error() {
					t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if tc.expectedTxnID != "" {
				if !startsWith(txnID, tc.expectedTxnID) {
					t.Errorf("expected transaction ID to start with %s, got: %s", tc.expectedTxnID, txnID)
				}
			} else if txnID != "" {
				t.Errorf("expected empty transaction ID, got: %s", txnID)
			}

			t.Logf("Test '%s' completed successfully", tc.description)
		})
	}

	t.Run("Process Payment with Large Random Transaction ID", func(t *testing.T) {
		processor := PayPalProcessor{}
		payment := Payment{
			Amount:   100.0,
			Currency: "USD",
			Method:   "PayPal",
		}
		for i := 0; i < 5; i++ {
			txnID, err := processor.Process(payment)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !validateTransactionID(txnID) {
				t.Errorf("invalid transaction ID format or range: %s", txnID)
			}
		}
	})
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func validateTransactionID(txnID string) bool {
	if !startsWith(txnID, "TXN-") {
		return false
	}
	numPart := txnID[4:]
	num, err := strconv.Atoi(numPart)
	if err != nil || num < 0 || num > 999999 {
		return false
	}
	return true
}


/*
ROOST_METHOD_HASH=Refund_e2af3d6bf2
ROOST_METHOD_SIG_HASH=Refund_4cecc33874

FUNCTION_DEF=func (p PayPalProcessor) Refund(transactionID string, amount float64) (string, error) 

 */
func TestPayPalProcessorRefund(t *testing.T) {
	type testCase struct {
		name          string
		transactionID string
		amount        float64
		expectedError error
		validate      func(t *testing.T, refundID string, err error)
	}

	testCases := []testCase{
		{
			name:          "Successful refund with valid transaction ID and positive amount",
			transactionID: "TXN123456",
			amount:        100.0,
			expectedError: nil,
			validate: func(t *testing.T, refundID string, err error) {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if refundID == "" {
					t.Error("Expected non-empty refund ID")
				}
				t.Log("Refund processed successfully with refund ID:", refundID)
			},
		},
		{
			name:          "Error on empty transaction ID",
			transactionID: "",
			amount:        100.0,
			expectedError: errors.New("invalid transaction ID"),
			validate: func(t *testing.T, refundID string, err error) {
				if err == nil || err.Error() != "invalid transaction ID" {
					t.Errorf("Expected error 'invalid transaction ID', got %v", err)
				}
				if refundID != "" {
					t.Error("Expected empty refund ID")
				}
				t.Log("Correctly handled empty transaction ID with error:", err)
			},
		},
		{
			name:          "Error on zero refund amount",
			transactionID: "TXN123456",
			amount:        0.0,
			expectedError: errors.New("refund amount must be positive"),
			validate: func(t *testing.T, refundID string, err error) {
				if err == nil || err.Error() != "refund amount must be positive" {
					t.Errorf("Expected error 'refund amount must be positive', got %v", err)
				}
				if refundID != "" {
					t.Error("Expected empty refund ID")
				}
				t.Log("Correctly handled zero refund amount with error:", err)
			},
		},
		{
			name:          "Error on negative refund amount",
			transactionID: "TXN123456",
			amount:        -50.0,
			expectedError: errors.New("refund amount must be positive"),
			validate: func(t *testing.T, refundID string, err error) {
				if err == nil || err.Error() != "refund amount must be positive" {
					t.Errorf("Expected error 'refund amount must be positive', got %v", err)
				}
				if refundID != "" {
					t.Error("Expected empty refund ID")
				}
				t.Log("Correctly handled negative refund amount with error:", err)
			},
		},
		{
			name:          "Consistent refund ID format",
			transactionID: "TXN123456",
			amount:        100.0,
			expectedError: nil,
			validate: func(t *testing.T, refundID string, err error) {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if len(refundID) <= 4 || refundID[:4] != "REF-" {
					t.Error("Refund ID does not start with 'REF-'")
				}
				t.Log("Refund ID format is consistent:", refundID)
			},
		},
		{
			name:          "Random refund ID uniqueness",
			transactionID: "TXN123456",
			amount:        100.0,
			expectedError: nil,
			validate: func(t *testing.T, refundID string, err error) {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				refundIDs := make(map[string]bool)
				for i := 0; i < 10; i++ {
					refundID, _ := PayPalProcessor{}.Refund("TXN123456", 100.0)
					if refundIDs[refundID] {
						t.Errorf("Duplicate refund ID found: %s", refundID)
					}
					refundIDs[refundID] = true
				}
				t.Log("All generated refund IDs are unique")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			refundID, err := PayPalProcessor{}.Refund(tc.transactionID, tc.amount)
			tc.validate(t, refundID, err)
		})
	}
}

