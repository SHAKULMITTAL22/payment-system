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
	tests := []struct {
		name     string
		value    float64
		currency string
		expected string
	}{
		{
			name:     "Format Positive Value with Default Currency Settings",
			value:    123.45,
			currency: "$",
			expected: "$123.45",
		},
		{
			name:     "Format Negative Value with Default Currency Settings",
			value:    -123.45,
			currency: "$",
			expected: "-$123.45",
		},
		{
			name:     "Format Zero Value",
			value:    0.0,
			currency: "$",
			expected: "$0.00",
		},
		{
			name:     "Format Large Value with Thousand Separator",
			value:    1234567.89,
			currency: "$",
			expected: "$1,234,567.89",
		},
		{
			name:     "Format Value with Different Currency Symbol",
			value:    123.45,
			currency: "€",
			expected: "€123.45",
		},
		{
			name:     "Format Value with Custom Precision",
			value:    123.456,
			currency: "$",
			expected: "$123.46",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)

			result := FormatCurrency(tt.value, tt.currency)

			if result != tt.expected {
				t.Errorf("Test failed for %s: expected %s, got %s", tt.name, tt.expected, result)
			} else {
				t.Logf("Test passed for %s: expected %s, got %s", tt.name, tt.expected, result)
			}
		})
	}
}

