package systemsignals

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/isavinof/pricer/log"
)

// SignalListener listen system signals
type SignalListener struct {
	lock     sync.Mutex
	signals  chan os.Signal
	handlers map[os.Signal][]func()
}

// NewSignalListener only initializations without subscription
// use SubscribeToShutdownSignals for subscribe
func NewSignalListener() SignalListener {
	return SignalListener{signals: make(chan os.Signal, 1), handlers: map[os.Signal][]func(){}}
}

// SubscribeToShutdownSignals now only SIGINT and SIGTERM listening
func (s *SignalListener) SubscribeToShutdownSignals(stopFunction func()) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.handlers[syscall.SIGINT] = append(s.handlers[syscall.SIGINT], stopFunction)
	s.handlers[syscall.SIGTERM] = append(s.handlers[syscall.SIGTERM], stopFunction)
	signal.Notify(s.signals, syscall.SIGINT, syscall.SIGTERM)
}

// ListenBlocked returning only after ctx.Done or received signal
func (s *SignalListener) ListenBlocked(ctx context.Context) {
	logger := log.FromContext(ctx)
	for {
		select {
		case sig := <-s.signals:
			logger.Infof("System signal received:%v", sig.String())
			func() {
				s.lock.Lock()
				defer s.lock.Unlock()
				for _, handler := range s.handlers[sig] {
					handler()
				}
			}()
		case <-ctx.Done():
			logger.Infof("System signal context done. Stop listen")
			return
		}
	}
}
