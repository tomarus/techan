package techan

import big "github.com/shopspring/decimal"

type differenceIndicator struct {
	minuend    Indicator
	subtrahend Indicator
}

// NewDifferenceIndicator returns an indicator which returns the difference between one indicator (minuend) and a second
// indicator (subtrahend).
func NewDifferenceIndicator(minuend, subtrahend Indicator) Indicator {
	return differenceIndicator{
		minuend:    minuend,
		subtrahend: subtrahend,
	}
}

func (di differenceIndicator) Calculate(index int) big.Decimal {
	return di.minuend.Calculate(index).Sub(di.subtrahend.Calculate(index))
}
