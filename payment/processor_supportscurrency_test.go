package payment

import "testing"







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
