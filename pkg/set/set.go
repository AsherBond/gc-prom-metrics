package set

import (
	"context"
	"errors"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/groundcover-com/metrics/pkg/pusher"
)

var (
	ErrPusherNotInitialized = errors.New("pusher not initialized")
)

type PushListener interface {
	OnTriggerPush()
	OnChangeInterval(interval *time.Duration)
}

type Set struct {
	set    *metrics.Set
	pusher *pusher.SetPusher
}

func NewSet() *Set {
	return &Set{
		set: metrics.NewSet(),
	}
}

func (s *Set) InitPush(
	ctx context.Context,
	pushURL string,
	interval time.Duration,
	opts *metrics.PushOptions,
) error {
	s.pusher = pusher.NewSetPusher(
		pusher.SetPusherOptions{
			Interval: &interval,
		},
		func(ctx context.Context) error {
			return s.set.PushMetrics(ctx, pushURL, opts)
		},
	)
	return s.pusher.Start(ctx)
}

func (s *Set) TriggerPush() error {
	pusher := s.pusher
	if pusher == nil {
		return ErrPusherNotInitialized
	}
	return pusher.TriggerPush()
}

func (s *Set) ChangePushInterval(interval time.Duration) error {
	pusher := s.pusher
	if pusher == nil {
		return ErrPusherNotInitialized
	}
	return pusher.ChangeInterval(&interval)
}

func (s *Set) CancelPushInterval() error {
	pusher := s.pusher
	if pusher == nil {
		return ErrPusherNotInitialized
	}
	return pusher.ChangeInterval(nil)
}
