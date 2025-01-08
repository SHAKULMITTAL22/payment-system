package currency

import (
	"testing"
	"github.com/leekchan/accounting"
)








/*
ROOST_METHOD_HASH=FormatCurrency_cc3bbdadf3
ROOST_METHOD_SIG_HASH=FormatCurrency_d831836e3e

FUNCTION_DEF=func FormatCurrency(value float64, currency string) string 

 */
func TestFormatCurrency(t *testing.T) {
	type testCase struct {
		value    float64
		currency string
		expected string
	}

	tests := []testCase{
		{
			value:    1234.56,
			currency: "$",
			expected: "$1,234.56",
		},
		{
			value:    -987.65,
			currency: "€",
			expected: "-€987.65",
		},
		{
			value:    0.0,
			currency: "£",
			expected: "£0.00",
		},
		{
			value:    1234567890.0,
			currency: "¥",
			expected: "¥1,234,567,890.00",
		},
		{
			value:    0.01,
			currency: "BTC",
			expected: "BTC0.01",
		},
		{
			value:    123.45,
			currency: "",
			expected: "123.45",
		},
	}

	for _, tc := range tests {
		t.Run(tc.expected, func(t *testing.T) {

			ac := accounting.Accounting{
				Symbol:    tc.currency,
				Precision: 2,
				Thousand:  ",",
				Decimal:   ".",
			}

			result := ac.FormatMoney(tc.value)

			if result != tc.expected {
				t.Errorf("expected %s but got %s", tc.expected, result)
			} else {
				t.Logf("success: value %f with currency %s formatted correctly as %s", tc.value, tc.currency, result)
			}
		})
	}
}

