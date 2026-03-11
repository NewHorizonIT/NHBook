package domain

import (
	"github.com/shopspring/decimal"
)

// Price represents a monetary value
type Price struct {
	value decimal.Decimal
}

func NewPrice(value decimal.Decimal) (Price, error) {
	if value.LessThanOrEqual(decimal.Zero) {
		return Price{}, ErrBookPriceInvalid
	}
	return Price{value: value}, nil
}

func NewPriceFromString(s string) (Price, error) {
	value, err := decimal.NewFromString(s)
	if err != nil {
		return Price{}, ErrBookPriceInvalid
	}
	return NewPrice(value)
}

func NewPriceFromFloat(f float64) (Price, error) {
	return NewPrice(decimal.NewFromFloat(f))
}

func (p Price) Value() decimal.Decimal {
	return p.value
}

func (p Price) String() string {
	return p.value.StringFixed(2)
}

func (p Price) Float64() float64 {
	f, _ := p.value.Float64()
	return f
}

// Stock represents inventory quantity
type Stock int32

func NewStock(quantity int32) (Stock, error) {
	if quantity < 0 {
		return 0, ErrBookStockInvalid
	}
	return Stock(quantity), nil
}

func (s Stock) Value() int32 {
	return int32(s)
}

func (s Stock) CanDeduct(quantity int32) bool {
	return int32(s) >= quantity
}

func (s Stock) Deduct(quantity int32) (Stock, error) {
	if !s.CanDeduct(quantity) {
		return s, ErrInsufficientStock
	}
	return Stock(int32(s) - quantity), nil
}

func (s Stock) Add(quantity int32) Stock {
	return Stock(int32(s) + quantity)
}
