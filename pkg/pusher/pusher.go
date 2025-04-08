package pusher

import (
	"context"
	"sync"

	"github.com/groundcover-com/metrics/pkg/set"
)

type Pusher struct {
	ctx      context.Context
	sets     []*pushedSet
	setsLock sync.RWMutex
	opts     PusherOptions
}

func NewPusher(ctx context.Context, opts PusherOptions) *Pusher {
	return &Pusher{ctx: ctx, opts: opts, sets: make([]*pushedSet, 0)}
}

func (p *Pusher) AddSet(set *set.Set, opts PushedSetOptions) {
	p.setsLock.Lock()
	defer p.setsLock.Unlock()

	pushedSet := newPushedSet(p.ctx, set, opts, p.opts)
	p.sets = append(p.sets, pushedSet)
	go pushedSet.startLoop()
}

func (p *Pusher) RemoveSet(set *set.Set) {
	p.setsLock.Lock()
	defer p.setsLock.Unlock()

	for i, s := range p.sets {
		if s.set == set {
			s.ctxCancel()
			p.sets = append(p.sets[:i], p.sets[i+1:]...)
		}
	}
}
