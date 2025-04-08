package set

import (
	"context"

	"github.com/groundcover-com/metrics/pkg/options"
)

type Pusher struct {
	sets []*Set
}

func NewPusher(ctx context.Context, opts options.PusherOptions) *Pusher {
	return &Pusher{}
}

func (p *Pusher) AddSet(set *Set) {
	p.sets = append(p.sets, set)
}

func (s *Set) loop() {
	select {
	case <-s.pushTriggerChannel:
		break
	}
}
