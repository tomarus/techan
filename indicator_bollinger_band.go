package techan

import (
	"github.com/shopspring/decimal"
)

type bbandIndicator struct {
	ma     Indicator
	stdev  Indicator
	muladd decimal.Decimal
}

// NewBollingerUpperBandIndicator a a derivative indicator which returns the upper bound of a bollinger band
// on the underlying indicator
func NewBollingerUpperBandIndicator(indicator Indicator, window int, sigma float64) Indicator {
	return bbandIndicator{
		ma:     NewSimpleMovingAverage(indicator, window),
		stdev:  NewWindowedStandardDeviationIndicator(indicator, window),
		muladd: decimal.NewFromFloat(sigma),
	}
}

// NewBollingerLowerBandIndicator returns a a derivative indicator which returns the lower bound of a bollinger band
// on the underlying indicator
func NewBollingerLowerBandIndicator(indicator Indicator, window int, sigma float64) Indicator {
	return bbandIndicator{
		ma:     NewSimpleMovingAverage(indicator, window),
		stdev:  NewWindowedStandardDeviationIndicator(indicator, window),
		muladd: decimal.NewFromFloat(-sigma),
	}
}

func (bbi bbandIndicator) Calculate(index int) decimal.Decimal {
	return bbi.ma.Calculate(index).Add(bbi.stdev.Calculate(index).Mul(bbi.muladd))
}
