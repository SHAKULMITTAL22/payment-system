package main

import (
	"fmt"

	"github.com/SHAKULMITTAL22/payment-system/currency"
	"github.com/SHAKULMITTAL22/payment-system/payment"
)

func main() {
	pay := payment.Payment{
		Amount:   100.5,
		Currency: "USD",
		Method:   "PayPal",
	}

	processor := payment.PayPalProcessor{}
	transactionID, err := processor.Process(pay)
	if err != nil {
		fmt.Println("Error processing payment:", err)
		return
	}

	formattedAmount := currency.FormatCurrency(pay.Amount, pay.Currency)
	fmt.Printf("Payment of %s successful with transaction ID: %s\n", formattedAmount, transactionID)
}
