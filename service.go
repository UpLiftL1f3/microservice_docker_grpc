package main

import (
	"context"
	"fmt"
)

// - PriceFetcher is an interface that can fetch a price
type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

// - priceFetcher implements the PriceFetcher interface
type priceService struct{}

func (s *priceService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

// * MOCK -------------- DATA -------------
// - THIS WILL ME USED TO MIMIC CoinMarketCap current prices
var priceMocks = map[string]float64{
	"BTC": 20000.0,
	"ETH": 200.0,
	"GG":  100000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
