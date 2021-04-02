package techan

import "github.com/shopspring/decimal"

type stopLossRule struct {
	Indicator
	tolerance decimal.Decimal
}

// NewStopLossRule returns a new rule that is satisfied when the given loss tolerance (a percentage) is met or exceeded.
// Loss tolerance should be a value between -1 and 1.
func NewStopLossRule(series *TimeSeries, lossTolerance float64) Rule {
	return stopLossRule{
		Indicator: NewClosePriceIndicator(series),
		tolerance: decimal.NewFromFloat(lossTolerance),
	}
}

func (slr stopLossRule) IsSatisfied(index int, record *TradingRecord) bool {
	if !record.CurrentPosition().IsOpen() {
		return false
	}

	openPrice := record.CurrentPosition().CostBasis()
	loss := slr.Indicator.Calculate(index).Div(openPrice).Sub(decimalONE)
	return loss.LessThanOrEqual(slr.tolerance)
}
