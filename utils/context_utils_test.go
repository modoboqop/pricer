package utils

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunWithTimeout(t *testing.T) {
	t.Run("when context timed out", func(t *testing.T) {
		t.Parallel()
		err := RunWithTimeout(context.Background(), time.Microsecond, func(ctx context.Context) error {
			time.Sleep(time.Millisecond)
			return nil
		})
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "context deadline exceeded")
	})

	t.Run("when context cancelled", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		err := RunWithTimeout(ctx, time.Hour, func(ctx context.Context) error {
			cancel()
			time.Sleep(time.Millisecond)
			return nil
		})
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "context canceled")
	})

	t.Run("when function error", func(t *testing.T) {
		t.Parallel()
		err := RunWithTimeout(context.Background(), time.Hour, func(ctx context.Context) error {
			return errors.New("my personal error")
		})
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "my personal error")
	})

	t.Run("when ok", func(t *testing.T) {
		t.Parallel()
		err := RunWithTimeout(context.Background(), time.Hour, func(ctx context.Context) error {
			return nil
		})
		assert.Nil(t, err)
	})

}
