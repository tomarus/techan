package techan

import (
	"math"

	"github.com/shopspring/decimal"
)

type aroonIndicator struct {
	indicator Indicator
	window    int
	direction decimal.Decimal
	lowIndex  int
}

func (ai *aroonIndicator) Calculate(index int) decimal.Decimal {
	if index < ai.window-1 {
		return decimalZERO
	}

	oneHundred := decimalTEN.Mul(decimalTEN)
	ai.lowIndex = ai.findLowIndex(index)
	pSince := decimal.NewFromFloat(float64(index - ai.lowIndex))
	windowAsDecimal := decimal.NewFromFloat(float64(ai.window))

	return windowAsDecimal.Sub(pSince).Div(windowAsDecimal).Mul(oneHundred)
}

func (ai aroonIndicator) findLowIndex(index int) int {
	if ai.lowIndex < 1 || ai.lowIndex < index-ai.window {
		lv := decimal.NewFromFloat(math.MaxFloat64)
		lowIndex := -1
		for i := (index + 1) - ai.window; i <= index; i++ {
			value := ai.indicator.Calculate(i).Mul(ai.direction)
			if value.LessThan(lv) {
				lv = value
				lowIndex = i
			}
		}

		return lowIndex
	}

	v1 := ai.indicator.Calculate(index).Mul(ai.direction)
	v2 := ai.indicator.Calculate(ai.lowIndex).Mul(ai.direction)

	if v1.LessThan(v2) {
		return index
	}

	return ai.lowIndex
}

// NewAroonUpIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the highest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a HighPriceIndicator or a derivative thereof
func NewAroonUpIndicator(indicator Indicator, window int) Indicator {
	return &aroonIndicator{
		indicator: indicator,
		window:    window,
		direction: decimalONE.Neg(),
		lowIndex:  -1,
	}
}

// NewAroonDownIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the lowest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a LowPriceIndicator or a derivative thereof
func NewAroonDownIndicator(indicator Indicator, window int) Indicator {
	return &aroonIndicator{
		indicator: indicator,
		window:    window,
		direction: decimalONE,
		lowIndex:  -1,
	}
}
