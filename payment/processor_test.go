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

FUNCTION_DEF=func (p PayPalProcessor) SupportsCurrency(currency string) bool 

 */
func TestPayPalProcessorSupportsCurrency(t *testing.T) {

	type testCase struct {
		name     string
		currency string
		expected bool
	}

	processor := PayPalProcessor{}

	testCases := []testCase{
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
			name:     "Lowercase Supported Currency",
			currency: "usd",
			expected: false,
		},
		{
			name:     "Empty String",
			currency: "",
			expected: false,
		},
		{
			name:     "Numeric String",
			currency: "123",
			expected: false,
		},
		{
			name:     "Long Unrecognized Currency Code",
			currency: "ABCDEFGHIJKLMNOP",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := processor.SupportsCurrency(tc.currency)

			if result != tc.expected {
				t.Errorf("Test '%s' failed: expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Test '%s' passed: expected %v, got %v", tc.name, tc.expected, result)
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
		name      string
		payment   Payment
		expectErr bool
		errMsg    string
	}

	testCases := []testCase{
		{
			name: "Valid Payment",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectErr: false,
		},
		{
			name: "Negative Amount",
			payment: Payment{
				Amount:   -50.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectErr: true,
			errMsg:    "amount must be positive",
		},
		{
			name: "Zero Amount",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectErr: true,
			errMsg:    "amount must be positive",
		},
		{
			name: "Empty Currency",
			payment: Payment{
				Amount:   100.0,
				Currency: "",
				Method:   "PayPal",
			},
			expectErr: true,
			errMsg:    "currency must be specified",
		},
		{
			name: "Invalid Currency and Negative Amount",
			payment: Payment{
				Amount:   -100.0,
				Currency: "",
				Method:   "PayPal",
			},
			expectErr: true,
			errMsg:    "amount must be positive",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processor := PayPalProcessor{}

			err := processor.ValidatePayment(tc.payment)

			if tc.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				} else if err.Error() != tc.errMsg {
					t.Errorf("expected error message '%s', but got '%s'", tc.errMsg, err.Error())
				} else {
					t.Logf("success: received expected error '%s'", err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got '%s'", err.Error())
				} else {
					t.Log("success: no error received as expected")
				}
			}
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
		name           string
		transactionID  string
		expectedError  error
		expectedOutput string
	}

	tests := []testCase{
		{
			name:           "Valid Transaction ID",
			transactionID:  "txn12345",
			expectedError:  nil,
			expectedOutput: "Details of Transaction ID: txn12345",
		},
		{
			name:           "Empty Transaction ID",
			transactionID:  "",
			expectedError:  errors.New("invalid transaction ID"),
			expectedOutput: "",
		},
		{
			name:           "Special Characters in Transaction ID",
			transactionID:  "!@#$%^&*",
			expectedError:  nil,
			expectedOutput: "Details of Transaction ID: !@#$%^&*",
		},
		{
			name:           "Numeric Transaction ID",
			transactionID:  "1234567890",
			expectedError:  nil,
			expectedOutput: "Details of Transaction ID: 1234567890",
		},
		{
			name:           "Long Transaction ID",
			transactionID:  generateLongTransactionID(1000),
			expectedError:  nil,
			expectedOutput: fmt.Sprintf("Details of Transaction ID: %s", generateLongTransactionID(1000)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			processor := PayPalProcessor{}
			result, err := processor.GetTransactionDetails(tc.transactionID)

			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("expected error: %v, got: %v", tc.expectedError, err)
			}

			if result != tc.expectedOutput {
				t.Errorf("expected output: %v, got: %v", tc.expectedOutput, result)
			}

			t.Logf("Test '%s' executed successfully", tc.name)
		})
	}
}

func generateLongTransactionID(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var result []rune
	for i := 0; i < length; i++ {
		result = append(result, chars[rand.Intn(len(chars))])
	}
	return string(result)
}


/*
ROOST_METHOD_HASH=Process_c114677531
ROOST_METHOD_SIG_HASH=Process_893bbedcfe

FUNCTION_DEF=func (p PayPalProcessor) Process(payment Payment) (string, error) 

 */
func TestPayPalProcessorProcess(t *testing.T) {
	tests := []struct {
		name          string
		payment       Payment
		expectedError error
		expectTxnID   bool
	}{
		{
			name: "Valid PayPal Method and Positive Amount",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedError: nil,
			expectTxnID:   true,
		},
		{
			name: "Unsupported Payment Method",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "CreditCard",
			},
			expectedError: errors.New("unsupported payment method"),
			expectTxnID:   false,
		},
		{
			name: "Zero Amount Payment",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedError: errors.New("amount must be greater than zero"),
			expectTxnID:   false,
		},
		{
			name: "Negative Amount Payment",
			payment: Payment{
				Amount:   -50.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedError: errors.New("amount must be greater than zero"),
			expectTxnID:   false,
		},
		{
			name: "Process Payment with Large Amount",
			payment: Payment{
				Amount:   1000000.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectedError: nil,
			expectTxnID:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := PayPalProcessor{}
			txnID, err := processor.Process(tt.payment)

			if tt.expectTxnID {
				if txnID == "" {
					t.Errorf("expected a transaction ID, got an empty string")
				}
			} else {
				if txnID != "" {
					t.Errorf("expected an empty transaction ID, got %s", txnID)
				}
			}

			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}

			if err == nil && tt.expectedError != nil {
				t.Errorf("expected error %v, got no error", tt.expectedError)
			}

			t.Logf("Test %s completed successfully", tt.name)
		})
	}
}


/*
ROOST_METHOD_HASH=Refund_e2af3d6bf2
ROOST_METHOD_SIG_HASH=Refund_4cecc33874

FUNCTION_DEF=func (p PayPalProcessor) Refund(transactionID string, amount float64) (string, error) 

 */
func TestPayPalProcessorRefund(t *testing.T) {

	rand.Seed(1)

	tests := []struct {
		name           string
		transactionID  string
		amount         float64
		expectedErr    error
		expectedRefund bool
	}{
		{
			name:           "Successful refund with valid transaction ID and positive amount",
			transactionID:  "TXN12345",
			amount:         100.0,
			expectedErr:    nil,
			expectedRefund: true,
		},
		{
			name:           "Refund with an empty transaction ID",
			transactionID:  "",
			amount:         100.0,
			expectedErr:    errors.New("invalid transaction ID"),
			expectedRefund: false,
		},
		{
			name:           "Refund with a negative amount",
			transactionID:  "TXN12345",
			amount:         -50.0,
			expectedErr:    errors.New("refund amount must be positive"),
			expectedRefund: false,
		},
		{
			name:           "Refund with a zero amount",
			transactionID:  "TXN12345",
			amount:         0.0,
			expectedErr:    errors.New("refund amount must be positive"),
			expectedRefund: false,
		},
		{
			name:           "Refund with a very large positive amount",
			transactionID:  "TXN12345",
			amount:         1e9,
			expectedErr:    nil,
			expectedRefund: true,
		},
		{
			name:           "Refund with non-numeric characters in transaction ID",
			transactionID:  "TXN-12345",
			amount:         100.0,
			expectedErr:    nil,
			expectedRefund: true,
		},
	}

	processor := PayPalProcessor{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			refundID, err := processor.Refund(tt.transactionID, tt.amount)

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
				} else {
					t.Logf("Success: expected error received: %v", err)
				}
				if refundID != "" {
					t.Errorf("expected empty refund ID, got: %s", refundID)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else {
					t.Logf("Success: no error returned")
				}
				if tt.expectedRefund && refundID == "" {
					t.Errorf("expected non-empty refund ID, got: %s", refundID)
				} else {
					t.Logf("Success: refund ID generated: %s", refundID)
				}
			}
		})
	}
}

