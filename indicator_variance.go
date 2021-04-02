package techan

import "github.com/shopspring/decimal"

// NewVarianceIndicator provides a way to find the variance in a base indicator, where variances is the sum of squared
// deviations from the mean at any given index in the time series.
func NewVarianceIndicator(ind Indicator) Indicator {
	return varianceIndicator{
		Indicator: ind,
	}
}

type varianceIndicator struct {
	Indicator Indicator
}

// Calculate returns the Variance for this indicator at the given index
func (vi varianceIndicator) Calculate(index int) decimal.Decimal {
	if index < 1 {
		return decimalZERO
	}

	avgIndicator := NewSimpleMovingAverage(vi.Indicator, index+1)
	avg := avgIndicator.Calculate(index)
	variance := decimalZERO

	for i := 0; i <= index; i++ {
		pow := vi.Indicator.Calculate(i).Sub(avg).Pow(decimal.NewFromInt(2))
		variance = variance.Add(pow)
	}

	return variance.Div(decimal.NewFromFloat(float64(index + 1)))
}
