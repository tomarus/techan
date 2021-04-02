package techan

import "github.com/shopspring/decimal"

type fixedIndicator []float64

// NewFixedIndicator returns an indicator with a fixed set of values that are returned when an index is passed in
func NewFixedIndicator(vals ...float64) Indicator {
	return fixedIndicator(vals)
}

func (fi fixedIndicator) Calculate(index int) decimal.Decimal {
	return decimal.NewFromFloat(fi[index])
}
