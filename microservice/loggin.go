package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceFetcher
}

// return a decorated service
func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, symbol string) (price float64, err error) {
	//decoration part
	defer func(begintime time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"path":      ctx.Value("path"),
			"price":     price,
			"err":       err,
			"took":      time.Since(begintime),
		}).Info("fetchPrice")
	}(time.Now())
	// origin part
	return s.next.FetchPrice(ctx, symbol)
}
