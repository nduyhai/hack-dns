package main

import (
	"context"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"net/http/httptrace"
	"time"
)

const (
	defaultTimeout = 100
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	ps := pool.New().WithContext(context.Background())
	for i := 0; i < 100; i++ {
		ps.Go(func(ctx context.Context) error {
			return remoteCall(ctx, client, logger)
		})
	}
	err := ps.Wait()
	if err != nil {
		logger.Error("wait err")
	}
}

func remoteCall(ctx context.Context, client *http.Client, logger *zap.Logger) error {
	n := rand.Intn(defaultTimeout)
	childCtx, cancelFunc := context.WithTimeout(ctx, time.Duration(n)*time.Millisecond)
	defer cancelFunc()

	trace := &httptrace.ClientTrace{
		DNSDone: func(info httptrace.DNSDoneInfo) {
			logger.Info("done looking up dns", zap.Any("dns_info", info), zap.Int("ms", n))
		},
	}

	traceCtx := httptrace.WithClientTrace(childCtx, trace)

	request, err := http.NewRequestWithContext(traceCtx, "GET", "http://jupiter:80", nil)
	if err != nil {
		logger.Debug("create http request", zap.Error(err))
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		logger.Debug("do http request", zap.Error(err))
		return err
	}

	return nil
}
