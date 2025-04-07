package factory

import (
	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/types"
)

func CreateHistogram(name string, labels map[string]string, opts options.Options) *types.Histogram {
	return types.NewHistogram(formatMetric(name, opts.Apply(labels)))
}

func GetOrCreateHistogram(name string, labels map[string]string, opts options.Options) *types.Histogram {
	return types.GetOrCreateHistogram(formatMetric(name, opts.Apply(labels)))
}

func CreateErrorHistogram(name string, labels map[string]string) *types.Histogram {
	return CreateHistogram(name, labels, options.Error)
}

func CreateWarningHistogram(name string, labels map[string]string) *types.Histogram {
	return CreateHistogram(name, labels, options.Warning)
}

func CreateInfoHistogram(name string, labels map[string]string) *types.Histogram {
	return CreateHistogram(name, labels, options.Info)
}

func GetOrCreateErrorHistogram(name string, labels map[string]string) *types.Histogram {
	return GetOrCreateHistogram(name, labels, options.Error)
}

func GetOrCreateWarningHistogram(name string, labels map[string]string) *types.Histogram {
	return GetOrCreateHistogram(name, labels, options.Warning)
}

func GetOrCreateInfoHistogram(name string, labels map[string]string) *types.Histogram {
	return GetOrCreateHistogram(name, labels, options.Info)
}
