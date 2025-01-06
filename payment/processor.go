package payment

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

// PayPalProcessor is a struct that processes payments using PayPal.
type PayPalProcessor struct{}

// Process handles the processing of a payment.
func (p PayPalProcessor) Process(payment Payment) (string, error) {
	if payment.Method != "PayPal" {
		return "", errors.New("unsupported payment method")
	}

	if payment.Amount <= 0 {
		return "", errors.New("amount must be greater than zero")
	}

	transactionID := "TXN-" + strconv.Itoa(rand.Intn(1000000))
	return transactionID, nil
}

// ValidatePayment checks the validity of a payment before processing.
func (p PayPalProcessor) ValidatePayment(payment Payment) error {
	if payment.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	if payment.Currency == "" {
		return errors.New("currency must be specified")
	}

	return nil
}

// Refund processes a refund for a given transaction ID.
func (p PayPalProcessor) Refund(transactionID string, amount float64) (string, error) {
	if transactionID == "" {
		return "", errors.New("invalid transaction ID")
	}

	if amount <= 0 {
		return "", errors.New("refund amount must be positive")
	}

	refundID := "REF-" + strconv.Itoa(rand.Intn(1000000))
	return refundID, nil
}

// GetTransactionDetails fetches the details of a given transaction.
func (p PayPalProcessor) GetTransactionDetails(transactionID string) (string, error) {
	if transactionID == "" {
		return "", errors.New("invalid transaction ID")
	}

	// For the sake of example, let's assume these details are fetched from a database.
	details := fmt.Sprintf("Details of Transaction ID: %s", transactionID)
	return details, nil
}

// SupportsCurrency checks if a particular currency is supported.
func (p PayPalProcessor) SupportsCurrency(currency string) bool {
	supportedCurrencies := []string{"USD", "EUR", "GBP"}
	for _, c := range supportedCurrencies {
		if c == currency {
			return true
		}
	}
	return false
}
