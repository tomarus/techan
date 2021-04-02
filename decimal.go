package techan

import (
	"math"

	"github.com/shopspring/decimal"
)

var decimalZERO = decimal.NewFromInt(0)
var decimalONE = decimal.NewFromInt(1)
var decimalTEN = decimal.NewFromInt(10)
var plusInf = decimal.NewFromInt(math.MaxInt64)
var minInf = decimal.NewFromInt(-math.MaxInt64)

// TODO: This is taken from https://github.com/shopspring/decimal/pull/130
// Remove this after it is merged into shopspring/decimal

// SqrtMaxIter sets a limit for number of iterations for the Sqrt function
const SqrtMaxIter = 100000

// Sqrt returns the square root of d, accurate to DivisionPrecision decimal places.
func Sqrt(d decimal.Decimal) decimal.Decimal {
	s, _ := SqrtRound(d, int32(decimal.DivisionPrecision))
	return s
}

// SqrtRound returns the square root of d, accurate to precision decimal places.
// The bool precise returns whether the precision was reached.
func SqrtRound(d decimal.Decimal, precision int32) (decimal.Decimal, bool) {
	maxError := decimal.New(1, -precision)
	one := decimal.NewFromFloat(1)
	var lo decimal.Decimal
	var hi decimal.Decimal
	// Handle cases where d < 0, d = 0, 0 < d < 1, and d > 1
	if d.GreaterThanOrEqual(one) {
		lo = decimalZERO
		hi = d
	} else if d.Equal(one) {
		return one, true
	} else if d.LessThan(decimalZERO) {
		return decimal.NewFromFloat(-1), false // call this an error , cannot take sqrt of neg w/o imaginaries
	} else if d.Equal(decimalZERO) {
		return decimalZERO, true
	} else {
		// d is between 0 and 1. Therefore, 0 < d < Sqrt(d) < 1.
		lo = d
		hi = one
	}
	var mid decimal.Decimal
	for i := 0; i < SqrtMaxIter; i++ {
		mid = lo.Add(hi).Div(decimal.New(2, 0)) //mid = (lo+hi)/2;
		if mid.Mul(mid).Sub(d).Abs().LessThan(maxError) {
			return mid, true
		}
		if mid.Mul(mid).GreaterThan(d) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return mid, false
}
