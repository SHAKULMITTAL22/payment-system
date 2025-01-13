package payment

import (
	"time"

	"github.com/govalues/decimal"
)

type Processor interface {
	Process(payment Payment) (string, error)
}

type RetryPolicy struct {
	MaxRetries int
	Backoff    time.Duration
}

type Payment struct {
	Amount      decimal.Decimal
	Currency    string
	Method      string
	ReferenceID string
	Timestamp   time.Time
	Metadata    [][]interface{} // Modified to use [][]interface{}
}
