package main

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcher is an interface that fetches prices for a given symbol
type PriceFetcher interface {
	FetchPrice(ctx context.Context, symbol string) (float64, error)
}

type priceFetcher struct{}

func (p *priceFetcher) FetchPrice(ctx context.Context, symbol string) (float64, error) {
	return MockPriceFetcher(ctx, symbol)
}

var priceMock = map[string]float64{
	"BTC":  100000,
	"ETH":  1000,
	"XRP":  1,
	"SOL":  100,
	"DOT":  10,
	"LINK": 100,
	"UNI":  1000,
}

func MockPriceFetcher(ctx context.Context, symbol string) (float64, error) {
	time.Sleep(100 * time.Millisecond)
	price, ok := priceMock[symbol]
	if !ok {
		return 0.0, fmt.Errorf("price for ticker (%s) is not available", symbol)
	}
	return price, nil
}
