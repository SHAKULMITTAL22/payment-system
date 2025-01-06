package currency

import (
	"testing"
	"github.com/leekchan/accounting"
)



type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}

func TestFormatCurrency(t *testing.T) {

	type testCase struct {
		value          float64
		currencySymbol string
		expected       string
		description    string
	}

	testCases := []testCase{
		{
			value:          1234.56,
			currencySymbol: "$",
			expected:       "$1,234.56",
			description:    "Format Positive Value with USD Currency",
		},
		{
			value:          -987.65,
			currencySymbol: "€",
			expected:       "-€987.65",
			description:    "Format Negative Value with EUR Currency",
		},
		{
			value:          0.0,
			currencySymbol: "£",
			expected:       "£0.00",
			description:    "Format Zero Value with GBP Currency",
		},
		{
			value:          1000000,
			currencySymbol: "¥",
			expected:       "¥1,000,000",
			description:    "Format Large Positive Value with JPY Currency (No Decimal)",
		},
		{
			value:          0.01,
			currencySymbol: "C$",
			expected:       "C$0.01",
			description:    "Format Small Fractional Value with CAD Currency",
		},
		{
			value:          -5000000,
			currencySymbol: "₹",
			expected:       "-₹5,000,000.00",
			description:    "Format Negative Large Value with INR Currency",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			result := FormatCurrency(tc.value, tc.currencySymbol)

			if result != tc.expected {
				t.Errorf("Test failed for %s: expected %v, got %v", tc.description, tc.expected, result)
			} else {
				t.Logf("Test passed for %s", tc.description)
			}
		})
	}
}
