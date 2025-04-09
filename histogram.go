package metrics

import (
	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/set"
)

type Histogram = set.Histogram

func CreateHistogram(name string, labels map[string]string, opts options.Options) *Histogram {
	return defaultSet.CreateHistogram(name, labels, opts)
}

func GetOrCreateHistogram(name string, labels map[string]string, opts options.Options) *Histogram {
	return defaultSet.GetOrCreateHistogram(name, labels, opts)
}

func CreateErrorHistogram(name string, labels map[string]string) *Histogram {
	return CreateHistogram(name, labels, options.Error)
}

func CreateWarningHistogram(name string, labels map[string]string) *Histogram {
	return CreateHistogram(name, labels, options.Warning)
}

func CreateInfoHistogram(name string, labels map[string]string) *Histogram {
	return CreateHistogram(name, labels, options.Info)
}

func GetOrCreateErrorHistogram(name string, labels map[string]string) *Histogram {
	return GetOrCreateHistogram(name, labels, options.Error)
}

func GetOrCreateWarningHistogram(name string, labels map[string]string) *Histogram {
	return GetOrCreateHistogram(name, labels, options.Warning)
}

func GetOrCreateInfoHistogram(name string, labels map[string]string) *Histogram {
	return GetOrCreateHistogram(name, labels, options.Info)
}
