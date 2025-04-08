package pusher

import (
	"time"

	"github.com/VictoriaMetrics/metrics"
)

type PusherOptions struct {
	URL         string
	PushOptions metrics.PushOptions
}

func NewPusherOptions() PusherOptions {
	return PusherOptions{}
}

func (o PusherOptions) WithPushOptions(pushOptions metrics.PushOptions) PusherOptions {
	o.PushOptions = pushOptions
	return o
}

func (o PusherOptions) WithURL(url string) PusherOptions {
	o.URL = url
	return o
}

type PushedSetOptions struct {
	Interval *time.Duration
}

func NewPushedSetOptions() PushedSetOptions {
	return PushedSetOptions{}
}

func (o PushedSetOptions) WithInterval(interval time.Duration) PushedSetOptions {
	o.Interval = &interval
	return o
}
