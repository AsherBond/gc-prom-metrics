package pusher

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SetPusher struct {
	opts     SetPusherOptions
	pushFunc func(context.Context) error

	intervalChangeChan chan *time.Duration
	triggerPushChan    chan struct{}
	channelsLock       sync.Mutex
}

func NewSetPusher(
	opts SetPusherOptions,
	pushFunc func(context.Context) error,
) *SetPusher {
	return &SetPusher{
		opts:     opts,
		pushFunc: pushFunc,
	}
}

func (p *SetPusher) Start(ctx context.Context) error {
	p.initiateChannels()
	defer p.closeChannels()

	for {
		shouldContinue, err := p.loop(ctx)
		if err != nil {
			return err
		}
		if !shouldContinue {
			return nil
		}
	}
}

func (p *SetPusher) TriggerPush() error {
	p.channelsLock.Lock()
	defer p.channelsLock.Unlock()

	ch := p.triggerPushChan
	if ch == nil {
		return fmt.Errorf("trigger push channel is nil")
	}

	select {
	case ch <- struct{}{}:
	default:
		return nil // this isn't an error because pushing is already triggered
	}

	return nil
}

func (p *SetPusher) ChangeInterval(interval *time.Duration) error {
	p.channelsLock.Lock()
	defer p.channelsLock.Unlock()

	p.opts.Interval = interval

	ch := p.intervalChangeChan
	if ch == nil {
		return fmt.Errorf("interval change channel is nil")
	}

	select {
	case ch <- interval:
	default:
		return fmt.Errorf("interval change channel is full")
	}

	return nil
}

func (p *SetPusher) loop(ctx context.Context) (shouldContinue bool, err error) {
	var tickerChannel <-chan time.Time
	if p.opts.Interval != nil {
		ticker := time.NewTicker(*p.opts.Interval)
		defer ticker.Stop()
		tickerChannel = ticker.C
	} else {
		ch := make(chan time.Time)
		defer close(ch)
		tickerChannel = ch
	}

	for {
		select {
		case <-tickerChannel:
			if err := p.pushFunc(ctx); err != nil {
				return false, fmt.Errorf("error interval-pushing metrics: %w", err)
			}
		case <-p.intervalChangeChan:
			return true, nil
		case <-p.triggerPushChan:
			if err := p.pushFunc(ctx); err != nil {
				return false, fmt.Errorf("error trigger-pushing metrics: %w", err)
			}
			return true, nil
		case <-ctx.Done():
			return false, nil
		}
	}
}

func (p *SetPusher) initiateChannels() {
	p.channelsLock.Lock()
	defer p.channelsLock.Unlock()

	p.intervalChangeChan = make(chan *time.Duration, 10)
	p.triggerPushChan = make(chan struct{})
}

func (p *SetPusher) closeChannels() {
	p.channelsLock.Lock()
	defer p.channelsLock.Unlock()

	close(p.intervalChangeChan)
	p.intervalChangeChan = nil

	close(p.triggerPushChan)
	p.triggerPushChan = nil
}
