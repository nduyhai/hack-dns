package main

import (
	"context"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptrace"
	"time"
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
			childCtx, cancelFunc := context.WithTimeout(ctx, 100*time.Millisecond)
			defer cancelFunc()
			return remoteCall(childCtx, client, logger)
		})
	}
	err := ps.Wait()
	if err != nil {
		logger.Error("wait err", zap.Error(err))
	}
}

func remoteCall(ctx context.Context, client *http.Client, logger *zap.Logger) error {
	trace := &httptrace.ClientTrace{
		DNSDone: func(info httptrace.DNSDoneInfo) {
			logger.Info("done looking up dns", zap.Any("dns_info", info))
		},
	}

	traceCtx := httptrace.WithClientTrace(ctx, trace)

	request, err := http.NewRequestWithContext(traceCtx, "GET", "https://tinhte.vn", nil)
	if err != nil {
		logger.Error("create http request", zap.Error(err))
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		logger.Error("do http request", zap.Error(err))
		return err
	}

	return nil
}
