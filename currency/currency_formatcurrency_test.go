package currency

import (
	"fmt"
	"testing"
)

func TestFormatCurrency(t *testing.T) {
	type testScenario struct {
		description string
		value       float64
		currency    string
		expected    string
	}

	// Define test scenarios
	scenarios := []testScenario{
		{
			description: "Scenario 1: Format Positive Value with USD Currency",
			value:       1234.56,
			currency:    "$",
			expected:    "$1,234.56",
		},
		{
			description: "Scenario 2: Format Negative Value with EUR Currency",
			value:       -987.65,
			currency:    "€",
			expected:    "-€987.65",
		},
		{
			description: "Scenario 3: Format Zero Value with GBP Currency",
			value:       0.0,
			currency:    "£",
			expected:    "£0.00",
		},
		{
			description: "Scenario 4: Format Large Positive Value with JPY Currency (No Decimal)",
			value:       1000000,
			currency:    "¥",
			expected:    "¥1,000,000",
		},
		{
			description: "Scenario 5: Format Small Fractional Value with CAD Currency",
			value:       0.01,
			currency:    "C$",
			expected:    "C$0.01",
		},
		{
			description: "Scenario 6: Format Negative Large Value with INR Currency",
			value:       -5000000,
			currency:    "₹",
			expected:    "-₹5,000,000.00",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			t.Log(fmt.Sprintf("Executing: %s", scenario.description)) // Logging scenario details

			// Act
			result := FormatCurrency(scenario.value, scenario.currency)

			// Assert
			if result != scenario.expected {
				t.Errorf("Failed: %s - Expected: %s, Got: %s", scenario.description, scenario.expected, result)
			} else {
				t.Log(fmt.Sprintf("Success: %s - Result: %s", scenario.description, result))
			}
		})
	}
}
