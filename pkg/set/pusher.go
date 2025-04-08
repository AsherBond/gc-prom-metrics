package set

import (
	"context"
	"time"

	"github.com/groundcover-com/metrics/pkg/options"
)

type pushedSet struct {
	set        *Set
	setOpts    options.PushedSetOptions
	pusherOpts options.PusherOptions
	ctx        context.Context
	ctxCancel  context.CancelFunc

	intervalChangeChan chan *time.Duration
	triggerPushChan    chan struct{}
}

func newPushedSet(
	ctx context.Context,
	set *Set,
	setOpts options.PushedSetOptions,
	pusherOpts options.PusherOptions,
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
	p.intervalChangeChan = make(chan *time.Duration)
	defer close(p.intervalChangeChan)

	p.triggerPushChan = make(chan struct{})
	defer close(p.triggerPushChan)

	p.set.RegisterPushListener(p)
	defer p.set.DeregisterPushListener(p)

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
			p.set.set.PushMetrics(p.ctx, p.pusherOpts.URL, &p.pusherOpts.PushOptions)
		case <-p.intervalChangeChan:
			return true
		case <-p.triggerPushChan:
			p.set.set.PushMetrics(p.ctx, p.pusherOpts.URL, &p.pusherOpts.PushOptions)
			return true
		case <-p.ctx.Done():
			return false
		}
	}
}

func (p *pushedSet) OnTriggerPush() {
	select {
	case p.triggerPushChan <- struct{}{}:
	default:
	}
}
func (p *pushedSet) OnChangeInterval(interval *time.Duration) {
	p.setOpts.Interval = interval

	select {
	case p.intervalChangeChan <- interval:
	default:
	}
}

type Pusher struct {
	ctx  context.Context
	sets []*pushedSet
	opts options.PusherOptions
}

func NewPusher(ctx context.Context, opts options.PusherOptions) *Pusher {
	return &Pusher{ctx: ctx, opts: opts, sets: make([]*pushedSet, 0)}
}

func (p *Pusher) AddSet(set *Set, opts options.PushedSetOptions) {
	pushedSet := newPushedSet(p.ctx, set, opts, p.opts)
	p.sets = append(p.sets, pushedSet)
	pushedSet.startLoop()
}

func (p *Pusher) RemoveSet(set *Set) {
	for i, s := range p.sets {
		if s.set == set {
			s.ctxCancel()
			p.sets = append(p.sets[:i], p.sets[i+1:]...)
		}
	}
}
