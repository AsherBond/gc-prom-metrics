package factory

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/groundcover-com/metrics/pkg/options"
)

func CreateHistogram(name string, labels map[string]string, opts options.Options) *metrics.Histogram {
	return defaultSet.CreateHistogram(name, labels, opts)
}

func GetOrCreateHistogram(name string, labels map[string]string, opts options.Options) *metrics.Histogram {
	return defaultSet.GetOrCreateHistogram(name, labels, opts)
}

func CreateErrorHistogram(name string, labels map[string]string) *metrics.Histogram {
	return CreateHistogram(name, labels, options.Error)
}

func CreateWarningHistogram(name string, labels map[string]string) *metrics.Histogram {
	return CreateHistogram(name, labels, options.Warning)
}

func CreateInfoHistogram(name string, labels map[string]string) *metrics.Histogram {
	return CreateHistogram(name, labels, options.Info)
}

func GetOrCreateErrorHistogram(name string, labels map[string]string) *metrics.Histogram {
	return GetOrCreateHistogram(name, labels, options.Error)
}

func GetOrCreateWarningHistogram(name string, labels map[string]string) *metrics.Histogram {
	return GetOrCreateHistogram(name, labels, options.Warning)
}

func GetOrCreateInfoHistogram(name string, labels map[string]string) *metrics.Histogram {
	return GetOrCreateHistogram(name, labels, options.Info)
}
