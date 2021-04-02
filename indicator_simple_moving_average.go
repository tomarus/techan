package techan

import big "github.com/shopspring/decimal"

type smaIndicator struct {
	indicator Indicator
	window    int
}

// NewSimpleMovingAverage returns a derivative Indicator which returns the average of the current value and preceding
// values in the given windowSize.
func NewSimpleMovingAverage(indicator Indicator, window int) Indicator {
	return smaIndicator{indicator, window}
}

func (sma smaIndicator) Calculate(index int) big.Decimal {
	if index < sma.window-1 {
		return decimalZERO
	}

	sum := decimalZERO
	for i := index; i > index-sma.window; i-- {
		sum = sum.Add(sma.indicator.Calculate(i))
	}

	result := sum.Div(big.NewFromInt(int64(sma.window)))

	return result
}
