package currency

import "github.com/leekchan/accounting"

func FormatCurrency(value float64, currency string) string {
	ac := accounting.Accounting{Symbol: currency, Precision: 2}
	return ac.FormatMoney(value)
}
