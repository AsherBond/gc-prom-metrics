package set

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/groundcover-com/metrics/pkg/options"
)

type FloatCounter = metrics.FloatCounter

func (set *Set) CreateFloatCounter(name string, labels map[string]string, opts options.Options) *FloatCounter {
	return set.set.NewFloatCounter(formatMetric(name, opts.Apply(labels)))
}

func (set *Set) GetOrCreateFloatCounter(name string, labels map[string]string, opts options.Options) *FloatCounter {
	return set.set.GetOrCreateFloatCounter(formatMetric(name, opts.Apply(labels)))
}

func (set *Set) CreateErrorFloatCounter(name string, labels map[string]string) *FloatCounter {
	return set.CreateFloatCounter(name, labels, options.Error)
}

func (set *Set) CreateWarningFloatCounter(name string, labels map[string]string) *FloatCounter {
	return set.CreateFloatCounter(name, labels, options.Warning)
}

func (set *Set) CreateInfoFloatCounter(name string, labels map[string]string) *FloatCounter {
	return set.CreateFloatCounter(name, labels, options.Info)
}

func (set *Set) GetOrCreateErrorFloatCounter(name string, labels map[string]string) *FloatCounter {
	return set.GetOrCreateFloatCounter(name, labels, options.Error)
}

func (set *Set) GetOrCreateWarningFloatCounter(name string, labels map[string]string) *FloatCounter {
	return set.GetOrCreateFloatCounter(name, labels, options.Warning)
}

func (set *Set) GetOrCreateInfoFloatCounter(name string, labels map[string]string) *FloatCounter {
	return set.GetOrCreateFloatCounter(name, labels, options.Info)
}
