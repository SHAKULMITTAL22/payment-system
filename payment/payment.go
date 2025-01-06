package payment

type Payment struct {
	Amount   float64
	Currency string
	Method   string
}

type Processor interface {
	Process(payment Payment) (string, error)
}
