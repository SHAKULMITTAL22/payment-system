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
	type test struct {
		value    float64
		currency string
		expected string
	}

	tests := []test{
		{value: 1234.56, currency: "$", expected: "$1,234.56"},
		{value: -1234.56, currency: "$", expected: "($1,234.56)"},
		{value: 0.00, currency: "$", expected: "$0.00"},
		{value: 1234.56, currency: "€", expected: "€1,234.56"},
		{value: 1234567890.12, currency: "$", expected: "$1,234,567,890.12"},
		{value: 1234.5678, currency: "$", expected: "$1,234.57"},
	}

	for _, tc := range tests {
		t.Run(tc.expected, func(t *testing.T) {

			result := FormatCurrency(tc.value, tc.currency)

			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			} else {
				t.Logf("Successfully formatted %f with currency %s to %s", tc.value, tc.currency, result)
			}
		})
	}
}

