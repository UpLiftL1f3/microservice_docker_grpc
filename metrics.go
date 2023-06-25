package main

import (
	"context"
	"fmt"
)

type metricsService struct {
	next PriceService
}

func NewMetricsService(next PriceService) PriceService {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Pushing metrics to Prometheus")
	//your metrics storage. Push to Prometheus (gauge, counters)
	return s.next.FetchPrice(ctx, ticker)
}
