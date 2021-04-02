package techan

import "github.com/shopspring/decimal"

type keltnerChannelIndicator struct {
	ema    Indicator
	atr    Indicator
	mul    decimal.Decimal
	window int
}

func NewKeltnerChannelUpperIndicator(series *TimeSeries, window int) Indicator {
	return keltnerChannelIndicator{
		atr:    NewAverageTrueRangeIndicator(series, window),
		ema:    NewEMAIndicator(NewClosePriceIndicator(series), window),
		mul:    decimalONE,
		window: window,
	}
}

func NewKeltnerChannelLowerIndicator(series *TimeSeries, window int) Indicator {
	return keltnerChannelIndicator{
		atr:    NewAverageTrueRangeIndicator(series, window),
		ema:    NewEMAIndicator(NewClosePriceIndicator(series), window),
		mul:    decimalONE.Neg(),
		window: window,
	}
}

func (kci keltnerChannelIndicator) Calculate(index int) decimal.Decimal {
	if index <= kci.window-1 {
		return decimalZERO
	}

	coefficient := decimal.NewFromInt(2).Mul(kci.mul)

	return kci.ema.Calculate(index).Add(kci.atr.Calculate(index).Mul(coefficient))
}
