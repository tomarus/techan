package techan

import (
	"github.com/shopspring/decimal"
)

type trueRangeIndicator struct {
	series *TimeSeries
}

// NewTrueRangeIndicator returns a base indicator
// which calculates the true range at the current point in time for a series
// https://www.investopedia.com/terms/a/atr.asp
func NewTrueRangeIndicator(series *TimeSeries) Indicator {
	return trueRangeIndicator{
		series: series,
	}
}

func (tri trueRangeIndicator) Calculate(index int) decimal.Decimal {
	if index-1 < 0 {
		return decimalZERO
	}

	candle := tri.series.Candles[index]
	previousClose := tri.series.Candles[index-1].ClosePrice

	trueHigh := decimal.Max(candle.MaxPrice, previousClose)
	trueLow := decimal.Min(candle.MinPrice, previousClose)

	return trueHigh.Sub(trueLow)
}
