package main

import (
	"time"

	"github.com/SHAKULMITTAL22/payment-system/payment"
)

// Nested structures for configuration.
type Config struct {
	ServiceSettings ServiceSettings
	APIKeys         map[string]string
}
type ServiceSettings struct {
	RetryPolicy payment.RetryPolicy
	Timeout     time.Duration
}

func check1(payment.Payment) {

}

// Function using nested structs and third-party calls.
// func ProcessWithConfig(ctx context.Context, payment payment.Payment, config Config, logger func(string)) (string, error) {

// 	// Example of using a function parameter for logging.
// 	logger("Starting payment processing")

// 	// Use third-party library to generate a UUID.
// 	txnID := uuid.New().String()

// 	// Simulating API key usage from nested struct.
// 	apiKey, ok := config.APIKeys[payment.Method]
// 	if !ok {
// 		return "", errors.New("unsupported payment method")
// 	}

// 	// Simulated third-party service call.
// 	error := simulateThirdPartyCall(apiKey, payment)
// 	if error != nil {
// 		return "", error
// 	}

// 	logger("Payment processed successfully")

// 	return txnID, nil
// }

// func simulateThirdPartyCall(apiKey string, payment payment.Payment) error {
// 	// Simulated external service interaction.
// 	if rand.Float32() < 0.1 { // Randomly fail to simulate real-world issues.
// 		return errors.New("third-party service call failed")
// 	}
// 	return nil
// }

// Payment represents a base structure for payments.
