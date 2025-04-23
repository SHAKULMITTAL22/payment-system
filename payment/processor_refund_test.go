package payment

import (
	"testing"
	"github.com/stretchr/testify/assert"
)







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
