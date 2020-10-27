package utils

import (
	"context"
	"time"

	"github.com/isavinof/pricer/lib/log"
)

// RunWithTimeout run passed function with timeout
func RunWithTimeout(ctx context.Context, timeout time.Duration, f func(ctx context.Context) error) (err error) {
	wait := make(chan error)
	defer close(wait)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	logger := log.FromContext(ctx)
	logger.Info("Run execution")
	go func() {
		wait <- f(ctx)
	}()

	logger.Infof("create timer:%v", timeout)
	timer := time.NewTimer(timeout)
	select {
	case <-ctx.Done():
		logger.Infof("context done %v", ctx.Err())
		return ctx.Err()
	case <-timer.C:
		logger.Infof("timer")
		return context.DeadlineExceeded
	case err = <-wait:
		logger.Infof("waited:%v", err)
		return err
	}
}
