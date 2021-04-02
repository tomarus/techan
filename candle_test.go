package techan

import (
	"testing"
	"time"

	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCandle_AddTrade(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	})

	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(2)) // Open
	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(5)) // High
	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(1)) // Low
	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(3)) // No Diff
	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(3)) // Close

	assert.EqualValues(t, "2", candle.OpenPrice.String())
	assert.EqualValues(t, "5", candle.MaxPrice.String())
	assert.EqualValues(t, "1", candle.MinPrice.String())
	assert.EqualValues(t, "3", candle.ClosePrice.String())
	assert.EqualValues(t, "5", candle.Volume.String())
	assert.EqualValues(t, 5, candle.TradeCount)
}

func TestCandle_String(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	})

	candle.ClosePrice = decimal.NewFromInt(1)
	candle.OpenPrice = decimal.NewFromInt(2)
	candle.MaxPrice = decimal.NewFromInt(3)
	candle.MinPrice = decimal.NewFromInt(0)
	candle.Volume = decimal.NewFromInt(10)

	expected := strings.TrimSpace(fmt.Sprintf(`
Time:	%s
Open:	2.00
Close:	1.00
High:	3.00
Low:	0.00
Volume:	10.00
`, candle.Period))

	assert.EqualValues(t, expected, candle.String())
}
