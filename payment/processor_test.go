package payment

import (
	"testing"
	"errors"
	"fmt"
	"math/rand"
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
			name:     "Supported Currency",
			currency: "USD",
			expected: true,
		},
		{
			name:     "Unsupported Currency",
			currency: "JPY",
			expected: false,
		},
		{
			name:     "Empty Currency String",
			currency: "",
			expected: false,
		},
		{
			name:     "Case Sensitivity Test",
			currency: "usd",
			expected: false,
		},
		{
			name:     "Currency with Special Characters",
			currency: "US$",
			expected: false,
		},
		{
			name:     "Numeric Currency String",
			currency: "123",
			expected: false,
		},
	}

	payPalProcessor := PayPalProcessor{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := payPalProcessor.SupportsCurrency(tt.currency)

			if result != tt.expected {
				t.Errorf("SupportsCurrency(%s) = %v; want %v", tt.currency, result, tt.expected)
			} else {
				t.Logf("SupportsCurrency(%s) = %v; as expected", tt.currency, result)
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
	tests := []struct {
		name          string
		payment       Payment
		expectedError error
	}{
		{
			name: "Scenario 1: Validate Positive Payment Amount",
			payment: Payment{
				Amount:   100.00,
				Currency: "USD",
				Method:   "Credit Card",
			},
			expectedError: nil,
		},
		{
			name: "Scenario 2: Validate Zero Payment Amount",
			payment: Payment{
				Amount:   0.00,
				Currency: "USD",
				Method:   "Credit Card",
			},
			expectedError: errors.New("amount must be positive"),
		},
		{
			name: "Scenario 3: Validate Negative Payment Amount",
			payment: Payment{
				Amount:   -50.00,
				Currency: "USD",
				Method:   "Credit Card",
			},
			expectedError: errors.New("amount must be positive"),
		},
		{
			name: "Scenario 4: Validate Missing Currency",
			payment: Payment{
				Amount:   100.00,
				Currency: "",
				Method:   "Credit Card",
			},
			expectedError: errors.New("currency must be specified"),
		},
		{
			name: "Scenario 5: Validate Valid Currency and Method",
			payment: Payment{
				Amount:   100.00,
				Currency: "EUR",
				Method:   "Debit Card",
			},
			expectedError: nil,
		},
		{
			name: "Scenario 6: Validate Empty Payment Method",
			payment: Payment{
				Amount:   100.00,
				Currency: "USD",
				Method:   "",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := PayPalProcessor{}
			err := processor.ValidatePayment(tt.payment)

			if err != nil && tt.expectedError == nil || err == nil && tt.expectedError != nil {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			} else if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message: %v, got: %v", tt.expectedError.Error(), err.Error())
			}

			t.Logf("Test %s completed. Expected error: %v, Got: %v", tt.name, tt.expectedError, err)
		})
	}
}


/*
ROOST_METHOD_HASH=GetTransactionDetails_d67406f337
ROOST_METHOD_SIG_HASH=GetTransactionDetails_9463e70b0f

FUNCTION_DEF=func (p PayPalProcessor) GetTransactionDetails(transactionID string) (string, error) 

 */
func TestPayPalProcessorGetTransactionDetails(t *testing.T) {
	type testCase struct {
		name            string
		transactionID   string
		expectedDetails string
		expectError     bool
	}

	tests := []testCase{
		{
			name:            "Valid Transaction ID",
			transactionID:   "12345",
			expectedDetails: "Details of Transaction ID: 12345",
			expectError:     false,
		},
		{
			name:            "Empty Transaction ID",
			transactionID:   "",
			expectedDetails: "",
			expectError:     true,
		},
		{
			name:            "Transaction ID with Special Characters",
			transactionID:   "trx-123!@#",
			expectedDetails: "Details of Transaction ID: trx-123!@#",
			expectError:     false,
		},
		{
			name:            "Long Transaction ID",
			transactionID:   generateLongTransactionID(1000),
			expectedDetails: fmt.Sprintf("Details of Transaction ID: %s", generateLongTransactionID(1000)),
			expectError:     false,
		},
		{
			name:            "Numerical Transaction ID",
			transactionID:   "9876543210",
			expectedDetails: "Details of Transaction ID: 9876543210",
			expectError:     false,
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			details, err := processor.GetTransactionDetails(tc.transactionID)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error but got nil")
				} else {
					t.Logf("Received expected error: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got: %v", err)
				} else if details != tc.expectedDetails {
					t.Errorf("Expected details: %s, but got: %s", tc.expectedDetails, details)
				} else {
					t.Logf("Received expected details: %s", details)
				}
			}
		})
	}
}

func generateLongTransactionID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}


/*
ROOST_METHOD_HASH=Process_c114677531
ROOST_METHOD_SIG_HASH=Process_893bbedcfe

FUNCTION_DEF=func (p PayPalProcessor) Process(payment Payment) (string, error) 

 */
func TestPayPalProcessorProcess(t *testing.T) {
	tests := []struct {
		name        string
		payment     Payment
		expectedID  string
		expectedErr error
	}{
		{
			name: "Valid PayPal Payment Processing",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:  "",
			expectedErr: nil,
		},
		{
			name: "Unsupported Payment Method",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedID:  "",
			expectedErr: errors.New("unsupported payment method"),
		},
		{
			name: "Zero Amount Payment",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:  "",
			expectedErr: errors.New("amount must be greater than zero"),
		},
		{
			name: "Negative Amount Payment",
			payment: Payment{
				Amount:   -50.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:  "",
			expectedErr: errors.New("amount must be greater than zero"),
		},
		{
			name: "Large Amount Payment",
			payment: Payment{
				Amount:   1000000.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:  "",
			expectedErr: nil,
		},
		{
			name: "Random Transaction ID Generation",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedID:  "",
			expectedErr: nil,
		},
		{
			name: "Non-PayPal Method with Valid Amount",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedID:  "",
			expectedErr: errors.New("unsupported payment method"),
		},
		{
			name: "Empty Currency Field",
			payment: Payment{
				Amount:   100.0,
				Currency: "",
				Method:   "PayPal",
			},
			expectedID:  "",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := PayPalProcessor{}
			transactionID, err := processor.Process(tt.payment)

			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("expected error '%v', got '%v'", tt.expectedErr, err)
			}

			if tt.expectedErr == nil && transactionID == "" {
				t.Errorf("expected non-empty transaction ID, got '%v'", transactionID)
			}

			if tt.name == "Random Transaction ID Generation" {

				transactionID2, _ := processor.Process(tt.payment)
				if transactionID == transactionID2 {
					t.Errorf("expected unique transaction IDs, got identical '%v'", transactionID)
				}
			}

			t.Logf("Test %s passed", tt.name)
		})
	}
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
		expectError   bool
		expectedError error
	}

	testCases := []testCase{
		{
			name:          "Valid Transaction ID and Positive Amount",
			transactionID: "12345",
			amount:        100.00,
			expectError:   false,
		},
		{
			name:          "Empty Transaction ID",
			transactionID: "",
			amount:        50.00,
			expectError:   true,
			expectedError: errors.New("invalid transaction ID"),
		},
		{
			name:          "Negative Refund Amount",
			transactionID: "67890",
			amount:        -50.00,
			expectError:   true,
			expectedError: errors.New("refund amount must be positive"),
		},
		{
			name:          "Zero Refund Amount",
			transactionID: "11223",
			amount:        0.00,
			expectError:   true,
			expectedError: errors.New("refund amount must be positive"),
		},
		{
			name:          "Large Refund Amount",
			transactionID: "44556",
			amount:        1000000.00,
			expectError:   false,
		},
	}

	p := PayPalProcessor{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			refundID, err := p.Refund(tc.transactionID, tc.amount)

			if tc.expectError {
				if err == nil {
					t.Errorf("expected error but got none for test case: %s", tc.name)
				} else if err.Error() != tc.expectedError.Error() {
					t.Errorf("expected error: %v, but got: %v for test case: %s", tc.expectedError, err, tc.name)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error but got: %v for test case: %s", err, tc.name)
				}
				if refundID == "" {
					t.Errorf("expected non-empty refund ID but got empty for test case: %s", tc.name)
				}
			}
		})
	}

	t.Run("Randomized Refund ID Generation", func(t *testing.T) {
		transactionID := "99887"
		amount := 200.00
		refundIDs := make(map[string]bool)

		for i := 0; i < 10; i++ {
			refundID, err := p.Refund(transactionID, amount)
			if err != nil {
				t.Fatalf("unexpected error during refund: %v", err)
			}
			if _, exists := refundIDs[refundID]; exists {
				t.Errorf("duplicate refund ID generated: %s", refundID)
			}
			refundIDs[refundID] = true
		}
		t.Log("Randomized Refund ID Generation test passed")
	})
}

