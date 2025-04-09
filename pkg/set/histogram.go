package set

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/groundcover-com/metrics/pkg/options"
)

type Histogram = metrics.Histogram

func (set *Set) CreateHistogram(name string, labels map[string]string, opts options.Options) *Histogram {
	return set.set.NewHistogram(formatMetric(name, opts.Apply(labels)))
}

func (set *Set) GetOrCreateHistogram(name string, labels map[string]string, opts options.Options) *Histogram {
	return set.set.GetOrCreateHistogram(formatMetric(name, opts.Apply(labels)))
}

func (set *Set) CreateErrorHistogram(name string, labels map[string]string) *Histogram {
	return set.CreateHistogram(name, labels, options.Error)
}

func (set *Set) CreateWarningHistogram(name string, labels map[string]string) *Histogram {
	return set.CreateHistogram(name, labels, options.Warning)
}

func (set *Set) CreateInfoHistogram(name string, labels map[string]string) *Histogram {
	return set.CreateHistogram(name, labels, options.Info)
}

func (set *Set) GetOrCreateErrorHistogram(name string, labels map[string]string) *Histogram {
	return set.GetOrCreateHistogram(name, labels, options.Error)
}

func (set *Set) GetOrCreateWarningHistogram(name string, labels map[string]string) *Histogram {
	return set.GetOrCreateHistogram(name, labels, options.Warning)
}

func (set *Set) GetOrCreateInfoHistogram(name string, labels map[string]string) *Histogram {
	return set.GetOrCreateHistogram(name, labels, options.Info)
}
