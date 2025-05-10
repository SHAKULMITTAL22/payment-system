package payment

import "testing"



type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestPayPalProcessorSupportsCurrency(t *testing.T) {
	type testCase struct {
		name     string
		currency string
		expected bool
	}

	tests := []testCase{
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
			name:     "Lowercase Supported Currency Code",
			currency: "usd",
			expected: false,
		},
		{
			name:     "Numeric Currency Code",
			currency: "123",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			processor := PayPalProcessor{}
			result := processor.SupportsCurrency(tc.currency)

			if result != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed: got expected result %v", tc.name, tc.expected)
			}
		})
	}
}
