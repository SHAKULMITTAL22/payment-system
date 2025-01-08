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
		name           string
		value          float64
		currencySymbol string
		expected       string
	}{
		{
			name:           "Scenario 1: Format Positive Value with Default Precision",
			value:          1234.56,
			currencySymbol: "$",
			expected:       "$1,234.56",
		},
		{
			name:           "Scenario 2: Format Negative Value",
			value:          -1234.56,
			currencySymbol: "$",
			expected:       "-$1,234.56",
		},
		{
			name:           "Scenario 3: Format Zero Value",
			value:          0.0,
			currencySymbol: "$",
			expected:       "$0.00",
		},
		{
			name:           "Scenario 4: Format Value with Different Currency Symbol",
			value:          1234.56,
			currencySymbol: "€",
			expected:       "€1,234.56",
		},
		{
			name:           "Scenario 5: Format Large Value",
			value:          1000000000.99,
			currencySymbol: "$",
			expected:       "$1,000,000,000.99",
		},
		{
			name:           "Scenario 6: Format Value with Custom Precision",
			value:          1234.5678,
			currencySymbol: "$",
			expected:       "$1,234.57",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ac := accounting.Accounting{
				Symbol:    tt.currencySymbol,
				Precision: 2,
				Thousand:  ",",
				Decimal:   ".",
			}

			result := ac.FormatMoney(tt.value)

			if result != tt.expected {
				t.Errorf("Test %s failed: expected %s, got %s", tt.name, tt.expected, result)
			} else {
				t.Logf("Test %s succeeded: expected %s, got %s", tt.name, tt.expected, result)
			}
		})
	}
}

