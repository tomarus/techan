package techan

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// Candle represents basic market information for a security over a given time period
type Candle struct {
	Period     TimePeriod
	OpenPrice  decimal.Decimal
	ClosePrice decimal.Decimal
	MaxPrice   decimal.Decimal
	MinPrice   decimal.Decimal
	Volume     decimal.Decimal
	TradeCount uint
}

// NewCandle returns a new *Candle for a given time period
func NewCandle(period TimePeriod) (c *Candle) {
	return &Candle{
		Period:     period,
		OpenPrice:  decimalZERO,
		ClosePrice: decimalZERO,
		MaxPrice:   decimalZERO,
		MinPrice:   decimalZERO,
		Volume:     decimalZERO,
	}
}

// AddTrade adds a trade to this candle. It will determine if the current price is higher or lower than the min or max
// price and increment the tradecount.
func (c *Candle) AddTrade(tradeAmount, tradePrice decimal.Decimal) {
	if c.OpenPrice.IsZero() {
		c.OpenPrice = tradePrice
	}
	c.ClosePrice = tradePrice

	if c.MaxPrice.IsZero() {
		c.MaxPrice = tradePrice
	} else if tradePrice.GreaterThan(c.MaxPrice) {
		c.MaxPrice = tradePrice
	}

	if c.MinPrice.IsZero() {
		c.MinPrice = tradePrice
	} else if tradePrice.LessThan(c.MinPrice) {
		c.MinPrice = tradePrice
	}

	if c.Volume.IsZero() {
		c.Volume = tradeAmount
	} else {
		c.Volume = c.Volume.Add(tradeAmount)
	}

	c.TradeCount++
}

func (c *Candle) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`
Time:	%s
Open:	%s
Close:	%s
High:	%s
Low:	%s
Volume:	%s
	`,
		c.Period,
		c.OpenPrice.StringFixed(2),
		c.ClosePrice.StringFixed(2),
		c.MaxPrice.StringFixed(2),
		c.MinPrice.StringFixed(2),
		c.Volume.StringFixed(2),
	))
}
