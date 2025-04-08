package set

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/VictoriaMetrics/metrics"
)

var (
	ErrNoRegisteredPushListener = errors.New("no registered push listener")
)

type PushListener interface {
	OnTriggerPush()
	OnChangeInterval(interval *time.Duration)
}

type Set struct {
	set *metrics.Set

	pushListeners     []PushListener
	pushListenersLock sync.RWMutex
}

func NewSet() *Set {
	return &Set{
		set: metrics.NewSet(),
	}
}

func (s *Set) RegisterPushListener(listener PushListener) {
	s.pushListenersLock.Lock()
	defer s.pushListenersLock.Unlock()

	s.pushListeners = append(s.pushListeners, listener)
}

func (s *Set) DeregisterPushListener(listener PushListener) {
	s.pushListenersLock.Lock()
	defer s.pushListenersLock.Unlock()

	for i, l := range s.pushListeners {
		if l == listener {
			s.pushListeners = append(s.pushListeners[:i], s.pushListeners[i+1:]...)
		}
	}
}

func (s *Set) PushMetrics(ctx context.Context, pushURL string, opts *metrics.PushOptions) error {
	return s.set.PushMetrics(ctx, pushURL, opts)
}

func (s *Set) TriggerPush() {
	s.pushListenersLock.RLock()
	defer s.pushListenersLock.RUnlock()

	for _, listener := range s.pushListeners {
		listener.OnTriggerPush()
	}
}

func (s *Set) ChangePushInterval(interval *time.Duration) {
	s.pushListenersLock.RLock()
	defer s.pushListenersLock.RUnlock()

	for _, listener := range s.pushListeners {
		listener.OnChangeInterval(interval)
	}
}
