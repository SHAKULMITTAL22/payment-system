package currency

import (
	"testing"
	"github.com/leekchan/accounting"
)







func TestFormatCurrency(t *testing.T) {
	type testCase struct {
		name     string
		value    float64
		currency string
		expected string
	}

	testCases := []testCase{
		{
			name:     "Format Positive Value with USD Currency",
			value:    1234.56,
			currency: "$",
			expected: "$1,234.56",
		},
		{
			name:     "Format Negative Value with EUR Currency",
			value:    -987.65,
			currency: "€",
			expected: "-€987.65",
		},
		{
			name:     "Format Zero Value with GBP Currency",
			value:    0.0,
			currency: "£",
			expected: "£0.00",
		},
		{
			name:     "Format Large Positive Value with JPY Currency (No Decimal)",
			value:    1000000,
			currency: "¥",
			expected: "¥1,000,000",
		},
		{
			name:     "Format Small Fractional Value with CAD Currency",
			value:    0.01,
			currency: "C$",
			expected: "C$0.01",
		},
		{
			name:     "Format Negative Large Value with INR Currency",
			value:    -5000000,
			currency: "₹",
			expected: "-₹5,000,000.00",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ac := accounting.Accounting{Symbol: tc.currency, Precision: 2, Thousand: ",", Decimal: "."}
			result := ac.FormatMoney(tc.value)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %s, got %s", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s succeeded: expected %s, got %s", tc.name, tc.expected, result)
			}
		})
	}
}
