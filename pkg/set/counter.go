package set

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/groundcover-com/metrics/pkg/options"
)

type Counter = metrics.Counter

func (set *Set) NewCounter(name string) *Counter {
	return set.set.NewCounter(name)
}

func (set *Set) CreateCounter(name string, labels map[string]string, opts options.Options) *Counter {
	return set.set.NewCounter(formatMetric(name, opts.Apply(labels)))
}

func (set *Set) GetOrCreateCounter(name string, labels map[string]string, opts options.Options) *Counter {
	return set.set.GetOrCreateCounter(formatMetric(name, opts.Apply(labels)))
}

func (set *Set) CreateErrorCounter(name string, labels map[string]string) *Counter {
	return set.CreateCounter(name, labels, options.Error)
}

func (set *Set) CreateWarningCounter(name string, labels map[string]string) *Counter {
	return set.CreateCounter(name, labels, options.Warning)
}

func (set *Set) CreateInfoCounter(name string, labels map[string]string) *Counter {
	return set.CreateCounter(name, labels, options.Info)
}

func (set *Set) GetOrCreateErrorCounter(name string, labels map[string]string) *Counter {
	return set.GetOrCreateCounter(name, labels, options.Error)
}

func (set *Set) GetOrCreateWarningCounter(name string, labels map[string]string) *Counter {
	return set.GetOrCreateCounter(name, labels, options.Warning)
}

func (set *Set) GetOrCreateInfoCounter(name string, labels map[string]string) *Counter {
	return set.GetOrCreateCounter(name, labels, options.Info)
}
