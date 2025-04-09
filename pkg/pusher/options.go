package pusher

import (
	"time"
)

type SetPusherOptions struct {
	Interval *time.Duration
}

func NewSetPusherOptions() SetPusherOptions {
	return SetPusherOptions{}
}

func (o SetPusherOptions) WithInterval(interval time.Duration) SetPusherOptions {
	o.Interval = &interval
	return o
}

func (o SetPusherOptions) CancelInterval() SetPusherOptions {
	o.Interval = nil
	return o
}
