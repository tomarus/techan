package techan

import (
	"github.com/shopspring/decimal"
)

type windowedStandardDeviationIndicator struct {
	Indicator
	movingAverage Indicator
	window        int
}

// NewWindowedStandardDeviationIndicator returns a indicator which calculates the standard deviation of the underlying
// indicator over a window
func NewWindowedStandardDeviationIndicator(ind Indicator, window int) Indicator {
	return windowedStandardDeviationIndicator{
		Indicator:     ind,
		movingAverage: NewSimpleMovingAverage(ind, window),
		window:        window,
	}
}

func (sdi windowedStandardDeviationIndicator) Calculate(index int) decimal.Decimal {
	avg := sdi.movingAverage.Calculate(index)
	variance := decimalZERO
	for i := Max(0, index-sdi.window+1); i <= index; i++ {
		pow := sdi.Indicator.Calculate(i).Sub(avg).Pow(decimal.NewFromInt(2))
		variance = variance.Add(pow)
	}
	realwindow := Min(sdi.window, index+1)

	tmp := variance.Div(decimal.NewFromFloat(float64(realwindow)))
	return Sqrt(tmp)
}
