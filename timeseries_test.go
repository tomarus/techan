package techan

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTimeSeries_AddCandle(t *testing.T) {
	t.Run("Throws if nil candle passed", func(t *testing.T) {
		ts := NewTimeSeries()
		assert.Panics(t, func() {
			ts.AddCandle(nil)
		})
	})

	t.Run("Adds candle if last is nil", func(t *testing.T) {
		ts := NewTimeSeries()

		candle := NewCandle(NewTimePeriod(time.Now(), time.Minute))
		candle.ClosePrice = decimal.NewFromInt(1)

		ts.AddCandle(candle)

		assert.Len(t, ts.Candles, 1)
	})

	t.Run("Does not add candle if before last candle", func(t *testing.T) {
		ts := NewTimeSeries()

		now := time.Now()
		candle := NewCandle(NewTimePeriod(now, time.Minute))
		candle.ClosePrice = decimal.NewFromInt(1)

		ts.AddCandle(candle)
		then := now.Add(-time.Minute * 10)

		nextCandle := NewCandle(NewTimePeriod(then, time.Minute))
		candle.ClosePrice = decimal.NewFromInt(2)

		ts.AddCandle(nextCandle)

		assert.Len(t, ts.Candles, 1)
		assert.EqualValues(t, now.UnixNano(), ts.Candles[0].Period.Start.UnixNano())
	})
}

func TestTimeSeries_LastCandle(t *testing.T) {
	ts := NewTimeSeries()

	now := time.Now()
	candle := NewCandle(NewTimePeriod(now, time.Minute))
	candle.ClosePrice = decimal.NewFromInt(1)

	ts.AddCandle(candle)

	assert.EqualValues(t, now.UnixNano(), ts.LastCandle().Period.Start.UnixNano())
	f, _ := ts.LastCandle().ClosePrice.Float64()
	assert.EqualValues(t, 1, f)

	next := time.Now().Add(time.Minute)
	newCandle := NewCandle(NewTimePeriod(next, time.Minute))
	newCandle.ClosePrice = decimal.NewFromInt(2)

	ts.AddCandle(newCandle)

	assert.Len(t, ts.Candles, 2)

	assert.EqualValues(t, next.UnixNano(), ts.LastCandle().Period.Start.UnixNano())
	f, _ = ts.LastCandle().ClosePrice.Float64()
	assert.EqualValues(t, 2, f)
}

func TestTimeSeries_LastIndex(t *testing.T) {
	ts := NewTimeSeries()

	candle := NewCandle(NewTimePeriod(time.Now(), time.Minute))
	ts.AddCandle(candle)

	assert.EqualValues(t, 0, ts.LastIndex())

	candle = NewCandle(NewTimePeriod(time.Now().Add(time.Minute), time.Minute))
	ts.AddCandle(candle)

	assert.EqualValues(t, 1, ts.LastIndex())
}
