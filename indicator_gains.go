package techan

import "github.com/shopspring/decimal"

type gainLossIndicator struct {
	Indicator
	coefficient decimal.Decimal
}

// NewGainIndicator returns a derivative indicator that returns the gains in the underlying indicator in the last bar,
// if any. If the delta is negative, zero is returned
func NewGainIndicator(indicator Indicator) Indicator {
	return gainLossIndicator{
		Indicator:   indicator,
		coefficient: decimalONE,
	}
}

// NewLossIndicator returns a derivative indicator that returns the losses in the underlying indicator in the last bar,
// if any. If the delta is positive, zero is returned
func NewLossIndicator(indicator Indicator) Indicator {
	return gainLossIndicator{
		Indicator:   indicator,
		coefficient: decimalONE.Neg(),
	}
}

func (gli gainLossIndicator) Calculate(index int) decimal.Decimal {
	if index == 0 {
		return decimalZERO
	}

	delta := gli.Indicator.Calculate(index).Sub(gli.Indicator.Calculate(index - 1)).Mul(gli.coefficient)
	if delta.GreaterThan(decimalZERO) {
		return delta
	}

	return decimalZERO
}

type cumulativeIndicator struct {
	Indicator
	window int
	mult   decimal.Decimal
}

// NewCumulativeGainsIndicator returns a derivative indicator which returns all gains made in a base indicator for a given
// window.
func NewCumulativeGainsIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      decimalONE,
	}
}

// NewCumulativeLossesIndicator returns a derivative indicator which returns all losses in a base indicator for a given
// window.
func NewCumulativeLossesIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      decimalONE.Neg(),
	}
}

func (ci cumulativeIndicator) Calculate(index int) decimal.Decimal {
	total := decimal.NewFromFloat(0.0)

	for i := Max(1, index-(ci.window-1)); i <= index; i++ {
		diff := ci.Indicator.Calculate(i).Sub(ci.Indicator.Calculate(i - 1))
		if diff.Mul(ci.mult).GreaterThan(decimalZERO) {
			total = total.Add(diff.Abs())
		}
	}

	return total
}

type percentChangeIndicator struct {
	Indicator
}

// NewPercentChangeIndicator returns a derivative indicator which returns the percent change (positive or negative)
// made in a base indicator up until the given indicator
func NewPercentChangeIndicator(indicator Indicator) Indicator {
	return percentChangeIndicator{indicator}
}

func (pgi percentChangeIndicator) Calculate(index int) decimal.Decimal {
	if index == 0 {
		return decimalZERO
	}

	cp := pgi.Indicator.Calculate(index)
	cplast := pgi.Indicator.Calculate(index - 1)
	return cp.Div(cplast).Sub(decimalONE)
}
