package example

import (
	"strconv"
	"time"

	"github.com/sdcoffey/techan"
	"github.com/shopspring/decimal"
)

// BasicEma is an example of how to create a basic Exponential moving average indicator
// based on the close prices of a timeseries from your exchange of choice.
func BasicEma() techan.Indicator {
	series := techan.NewTimeSeries()

	// fetch this from your preferred exchange
	dataset := [][]string{
		// Timestamp, Open, Close, High, Low, volume
		{"1234567", "1", "2", "3", "5", "6"},
	}

	for _, datum := range dataset {
		start, _ := strconv.ParseInt(datum[0], 10, 64)
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

		candle := techan.NewCandle(period)
		candle.OpenPrice, _ = decimal.NewFromString(datum[1])
		candle.ClosePrice, _ = decimal.NewFromString(datum[2])
		candle.MaxPrice, _ = decimal.NewFromString(datum[3])
		candle.MinPrice, _ = decimal.NewFromString(datum[4])

		series.AddCandle(candle)
	}

	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewEMAIndicator(closePrices, 10) // Create an exponential moving average with a window of 10

	return movingAverage
}
