package options

import (
	"time"

	"github.com/VictoriaMetrics/metrics"
)

type PusherOptions struct {
	Interval    *time.Duration
	URL         string
	PushOptions metrics.PushOptions
}

func NewPusherOptions() PusherOptions {
	return PusherOptions{}
}

func (o PusherOptions) WithInterval(interval time.Duration) PusherOptions {
	o.Interval = &interval
	return o
}

func (o PusherOptions) WithPushOptions(pushOptions metrics.PushOptions) PusherOptions {
	o.PushOptions = pushOptions
	return o
}

func (o PusherOptions) WithURL(url string) PusherOptions {
	o.URL = url
	return o
}
