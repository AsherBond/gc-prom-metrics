package pusher

import (
	"context"
	"sync"
	"time"

	"github.com/groundcover-com/metrics/pkg/set"
)

type pushedSet struct {
	set        *set.Set
	setOpts    PushedSetOptions
	pusherOpts PusherOptions
	ctx        context.Context
	ctxCancel  context.CancelFunc

	intervalChangeChan chan *time.Duration
	triggerPushChan    chan struct{}
	channelsLock       sync.Mutex
}

func newPushedSet(
	ctx context.Context,
	set *set.Set,
	setOpts PushedSetOptions,
	pusherOpts PusherOptions,
) *pushedSet {
	ctx, ctxCancel := context.WithCancel(ctx)

	return &pushedSet{
		set:        set,
		setOpts:    setOpts,
		pusherOpts: pusherOpts,
		ctx:        ctx,
		ctxCancel:  ctxCancel,
	}
}

func (p *pushedSet) startLoop() {
	p.channelsLock.Lock()
	p.intervalChangeChan = make(chan *time.Duration, 10)
	p.triggerPushChan = make(chan struct{})
	p.channelsLock.Unlock()

	p.set.RegisterPushListener(p)

	defer func() {
		p.set.DeregisterPushListener(p)

		p.channelsLock.Lock()
		defer p.channelsLock.Unlock()

		close(p.intervalChangeChan)
		p.intervalChangeChan = nil

		close(p.triggerPushChan)
		p.triggerPushChan = nil
	}()

	for {
		if !p.loop() {
			return
		}
	}
}

func (p *pushedSet) loop() bool {
	var tickerChannel <-chan time.Time
	if p.setOpts.Interval != nil {
		ticker := time.NewTicker(*p.setOpts.Interval)
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
			p.set.PushMetrics(p.ctx, p.pusherOpts.URL, &p.pusherOpts.PushOptions)
		case <-p.intervalChangeChan:
			return true
		case <-p.triggerPushChan:
			p.set.PushMetrics(p.ctx, p.pusherOpts.URL, &p.pusherOpts.PushOptions)
			return true
		case <-p.ctx.Done():
			return false
		}
	}
}

func (p *pushedSet) OnTriggerPush() {
	p.channelsLock.Lock()
	defer p.channelsLock.Unlock()

	ch := p.triggerPushChan
	if ch == nil {
		return
	}

	select {
	case ch <- struct{}{}:
	default:
	}
}

func (p *pushedSet) OnChangeInterval(interval *time.Duration) {
	p.channelsLock.Lock()
	defer p.channelsLock.Unlock()

	p.setOpts.Interval = interval

	ch := p.intervalChangeChan
	if ch == nil {
		return
	}

	select {
	case ch <- interval:
	default:
	}
}
