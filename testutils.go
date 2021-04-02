package techan

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"strconv"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var candleIndex int
var mockedTimeSeries = mockTimeSeriesFl(
	64.75, 63.79, 63.73,
	63.73, 63.55, 63.19,
	63.91, 63.85, 62.95,
	63.37, 61.33, 61.51)

func randomTimeSeries(size int) *TimeSeries {
	vals := make([]string, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		val := rand.Float64() * 100
		if i == 0 {
			vals[i] = fmt.Sprint(val)
		} else {
			last, _ := strconv.ParseFloat(vals[i-1], 64)
			if i%2 == 0 {
				vals[i] = fmt.Sprint(last + (val / 10))
			} else {
				vals[i] = fmt.Sprint(last - (val / 10))
			}
		}
	}

	return mockTimeSeries(vals...)
}

func mockTimeSeriesOCHL(values ...[]float64) *TimeSeries {
	ts := NewTimeSeries()
	for i, ochl := range values {
		candle := NewCandle(NewTimePeriod(time.Unix(int64(i), 0), time.Second))
		candle.OpenPrice = decimal.NewFromFloat(ochl[0])
		candle.ClosePrice = decimal.NewFromFloat(ochl[1])
		candle.MaxPrice = decimal.NewFromFloat(ochl[2])
		candle.MinPrice = decimal.NewFromFloat(ochl[3])
		candle.Volume = decimal.NewFromFloat(float64(i))

		ts.AddCandle(candle)
	}

	return ts
}

func mockTimeSeries(values ...string) *TimeSeries {
	ts := NewTimeSeries()
	for _, val := range values {
		maxp, _ := decimal.NewFromString(val)
		minp, _ := decimal.NewFromString(val)
		candle := NewCandle(NewTimePeriod(time.Unix(int64(candleIndex), 0), time.Second))
		candle.OpenPrice, _ = decimal.NewFromString(val)
		candle.ClosePrice, _ = decimal.NewFromString(val)
		candle.MaxPrice = maxp.Add(decimalONE)
		candle.MinPrice = minp.Sub(decimalONE)
		candle.Volume, _ = decimal.NewFromString(val)

		ts.AddCandle(candle)

		candleIndex++
	}

	return ts
}

func mockTimeSeriesFl(values ...float64) *TimeSeries {
	strVals := make([]string, len(values))

	for i, val := range values {
		strVals[i] = fmt.Sprint(val)
	}

	return mockTimeSeries(strVals...)
}

func decimalEquals(t *testing.T, expected float64, actual decimal.Decimal) {
	actf, _ := actual.Float64()
	assert.Equal(t, fmt.Sprintf("%.4f", expected), fmt.Sprintf("%.4f", actf))
}

func dump(indicator Indicator) (values []float64) {
	precision := 4.0
	m := math.Pow(10, precision)

	defer func() {
		recover()
	}()

	var index int
	for {
		res, _ := indicator.Calculate(index).Float64()
		values = append(values, math.Round(res*m)/m)
		index++
	}

	return
}

func indicatorEquals(t *testing.T, expected []float64, indicator Indicator) {
	actualValues := dump(indicator)
	assert.EqualValues(t, expected, actualValues)
}
