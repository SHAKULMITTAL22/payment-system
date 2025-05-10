package payment

import (
	"testing"
	"errors"
)








/*
ROOST_METHOD_HASH=SupportsCurrency_875c78ed3e
ROOST_METHOD_SIG_HASH=SupportsCurrency_d2d80e5419

FUNCTION_DEF=func (p PayPalProcessor) SupportsCurrency(currency string) bool 

 */
func TestPayPalProcessorSupportsCurrency(t *testing.T) {

	processor := PayPalProcessor{}

	testCases := []struct {
		name           string
		currency       string
		expectedResult bool
	}{
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
			name:           "Case Sensitivity Test - usd",
			currency:       "usd",
			expectedResult: false,
		},
		{
			name:           "Numeric Currency Code - 123",
			currency:       "123",
			expectedResult: false,
		},
		{
			name:           "Special Characters in Currency Code - US$",
			currency:       "US$",
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := processor.SupportsCurrency(tc.currency)

			if result != tc.expectedResult {
				t.Errorf("Test failed for %s: expected %v, got %v", tc.name, tc.expectedResult, result)
			} else {
				t.Logf("Test passed for %s: expected %v, got %v", tc.name, tc.expectedResult, result)
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
		name        string
		payment     Payment
		expectErr   bool
		expectedMsg string
	}

	testCases := []testCase{
		{
			name: "Valid Payment",
			payment: Payment{
				Amount:   100.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectErr:   false,
			expectedMsg: "",
		},
		{
			name: "Zero Amount",
			payment: Payment{
				Amount:   0.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectErr:   true,
			expectedMsg: "amount must be positive",
		},
		{
			name: "Negative Amount",
			payment: Payment{
				Amount:   -50.0,
				Currency: "USD",
				Method:   "PayPal",
			},
			expectErr:   true,
			expectedMsg: "amount must be positive",
		},
		{
			name: "Empty Currency",
			payment: Payment{
				Amount:   100.0,
				Currency: "",
				Method:   "PayPal",
			},
			expectErr:   true,
			expectedMsg: "currency must be specified",
		},
		{
			name: "Missing Currency and Negative Amount",
			payment: Payment{
				Amount:   -100.0,
				Currency: "",
				Method:   "PayPal",
			},
			expectErr: true,
		},
	}

	processor := PayPalProcessor{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := processor.ValidatePayment(tc.payment)

			if tc.expectErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				} else if tc.expectedMsg != "" && err.Error() != tc.expectedMsg {
					t.Errorf("expected error message '%s', got '%s'", tc.expectedMsg, err.Error())
				} else {
					t.Logf("success: received expected error '%s'", err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got '%s'", err.Error())
				} else {
					t.Log("success: no error as expected")
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Process_c114677531
ROOST_METHOD_SIG_HASH=Process_893bbedcfe

FUNCTION_DEF=func (p PayPalProcessor) Process(payment Payment) (string, error) 

 */
func TestPayPalProcessorProcess(t *testing.T) {
	type args struct {
		payment Payment
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		expectedErrMsg string
	}{
		{
			name: "Successful Payment Processing with Valid PayPal Method",
			args: args{
				payment: Payment{
					Amount:   100.0,
					Currency: "USD",
					Method:   "PayPal",
				},
			},
			wantErr:        false,
			expectedErrMsg: "",
		},
		{
			name: "Unsupported Payment Method",
			args: args{
				payment: Payment{
					Amount:   100.0,
					Currency: "USD",
					Method:   "CreditCard",
				},
			},
			wantErr:        true,
			expectedErrMsg: "unsupported payment method",
		},
		{
			name: "Zero Payment Amount",
			args: args{
				payment: Payment{
					Amount:   0.0,
					Currency: "USD",
					Method:   "PayPal",
				},
			},
			wantErr:        true,
			expectedErrMsg: "amount must be greater than zero",
		},
		{
			name: "Negative Payment Amount",
			args: args{
				payment: Payment{
					Amount:   -50.0,
					Currency: "USD",
					Method:   "PayPal",
				},
			},
			wantErr:        true,
			expectedErrMsg: "amount must be greater than zero",
		},
		{
			name: "Multiple Consecutive Transactions",
			args: args{
				payment: Payment{
					Amount:   100.0,
					Currency: "USD",
					Method:   "PayPal",
				},
			},
			wantErr:        false,
			expectedErrMsg: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PayPalProcessor{}
			transactionID, err := p.Process(tt.args.payment)

			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.expectedErrMsg {
				t.Errorf("Process() error message = %v, expectedErrMsg %v", err.Error(), tt.expectedErrMsg)
			}
			if !tt.wantErr && transactionID == "" {
				t.Errorf("Process() transactionID is empty, expected a valid transaction ID")
			}

			if err == nil {
				t.Logf("Process() succeeded with transactionID = %v", transactionID)
			}
		})
	}

	t.Run("Transaction ID Uniqueness", func(t *testing.T) {
		p := PayPalProcessor{}
		transactionIDs := make(map[string]bool)

		for i := 0; i < 100; i++ {
			transactionID, err := p.Process(Payment{Amount: 100.0, Currency: "USD", Method: "PayPal"})
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if _, exists := transactionIDs[transactionID]; exists {
				t.Errorf("Duplicate transaction ID found: %v", transactionID)
			}
			transactionIDs[transactionID] = true
		}
	})
}


/*
ROOST_METHOD_HASH=Refund_e2af3d6bf2
ROOST_METHOD_SIG_HASH=Refund_4cecc33874

FUNCTION_DEF=func (p PayPalProcessor) Refund(transactionID string, amount float64) (string, error) 

 */
func TestPayPalProcessorRefund(t *testing.T) {
	type testCase struct {
		description    string
		transactionID  string
		amount         float64
		expectedError  error
		expectedFormat string
	}

	testCases := []testCase{
		{
			description:    "Scenario 1: Successful refund with valid transaction ID and positive amount",
			transactionID:  "validTransaction123",
			amount:         100.0,
			expectedError:  nil,
			expectedFormat: "REF-",
		},
		{
			description:    "Scenario 2: Error on empty transaction ID",
			transactionID:  "",
			amount:         50.0,
			expectedError:  errors.New("invalid transaction ID"),
			expectedFormat: "",
		},
		{
			description:    "Scenario 3: Error on zero refund amount",
			transactionID:  "validTransaction456",
			amount:         0.0,
			expectedError:  errors.New("refund amount must be positive"),
			expectedFormat: "",
		},
		{
			description:    "Scenario 4: Error on negative refund amount",
			transactionID:  "validTransaction789",
			amount:         -10.0,
			expectedError:  errors.New("refund amount must be positive"),
			expectedFormat: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			processor := PayPalProcessor{}
			refundID, err := processor.Refund(tc.transactionID, tc.amount)

			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("expected error: %v, got: %v", tc.expectedError, err)
			}

			if tc.expectedFormat != "" && len(refundID) > 0 {
				if !assertRefundIDFormat(refundID, tc.expectedFormat) {
					t.Errorf("refund ID %s does not match expected format %s", refundID, tc.expectedFormat)
				}
			}

			t.Logf("Test case '%s' completed successfully.", tc.description)
		})
	}

	t.Run("Scenario 5: Consistent refund ID format", func(t *testing.T) {
		processor := PayPalProcessor{}
		transactionID := "consistentFormatTransaction"
		amount := 25.0

		for i := 0; i < 5; i++ {
			refundID, _ := processor.Refund(transactionID, amount)
			if !assertRefundIDFormat(refundID, "REF-") {
				t.Errorf("refund ID %s does not match expected format REF-", refundID)
			}
		}
		t.Log("Refund ID format is consistent across multiple calls.")
	})

	t.Run("Scenario 6: Randomness of refund ID", func(t *testing.T) {
		processor := PayPalProcessor{}
		transactionID := "randomnessTransaction"
		amount := 30.0
		refundIDs := make(map[string]bool)

		for i := 0; i < 100; i++ {
			refundID, _ := processor.Refund(transactionID, amount)
			if refundIDs[refundID] {
				t.Errorf("duplicate refund ID found: %s", refundID)
			}
			refundIDs[refundID] = true
		}
		t.Log("Refund IDs are unique across multiple calls.")
	})
}

func assertRefundIDFormat(refundID, expectedFormat string) bool {
	return len(refundID) > len(expectedFormat) && refundID[:len(expectedFormat)] == expectedFormat
}

