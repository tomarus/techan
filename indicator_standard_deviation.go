package techan

import (
	big "github.com/shopspring/decimal"
)

// NewStandardDeviationIndicator calculates the standard deviation of a base indicator.
// See https://www.investopedia.com/terms/s/standarddeviation.asp
func NewStandardDeviationIndicator(ind Indicator) Indicator {
	return standardDeviationIndicator{
		indicator: NewVarianceIndicator(ind),
	}
}

type standardDeviationIndicator struct {
	indicator Indicator
}

// Calculate returns the standard deviation of a base indicator
func (sdi standardDeviationIndicator) Calculate(index int) big.Decimal {
	return Sqrt(sdi.indicator.Calculate(index))
}
