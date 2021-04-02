package techan

import (
	"github.com/shopspring/decimal"
)

type trendLineIndicator struct {
	indicator Indicator
	window    int
}

// NewTrendlineIndicator returns an indicator whose output is the slope of the trend
// line given by the values in the window.
func NewTrendlineIndicator(indicator Indicator, window int) Indicator {
	return trendLineIndicator{
		indicator: indicator,
		window:    window,
	}
}

func (tli trendLineIndicator) Calculate(index int) decimal.Decimal {
	window := Min(index+1, tli.window)

	values := make([]decimal.Decimal, window)

	for i := 0; i < window; i++ {
		values[i] = tli.indicator.Calculate(index - (window - 1) + i)
	}

	n := decimalONE.Mul(decimal.NewFromFloat(float64(window)))
	ab := sumXy(values).Mul(n).Sub(sumX(values).Mul(sumY(values)))
	cd := sumX2(values).Mul(n).Sub(sumX(values).Pow(decimal.NewFromInt(2)))

	return ab.Div(cd)
}

func sumX(decimals []decimal.Decimal) (s decimal.Decimal) {
	s = decimalZERO

	for i := range decimals {
		s = s.Add(decimal.NewFromFloat(float64(i)))
	}

	return s
}

func sumY(decimals []decimal.Decimal) (b decimal.Decimal) {
	b = decimalZERO
	for _, d := range decimals {
		b = b.Add(d)
	}

	return
}

func sumXy(decimals []decimal.Decimal) (b decimal.Decimal) {
	b = decimalZERO

	for i, d := range decimals {
		b = b.Add(d.Mul(decimal.NewFromFloat(float64(i))))
	}

	return
}

func sumX2(decimals []decimal.Decimal) decimal.Decimal {
	b := decimalZERO

	for i := range decimals {
		b = b.Add(decimal.NewFromFloat(float64(i)).Pow(decimal.NewFromInt(2)))
	}

	return b
}
