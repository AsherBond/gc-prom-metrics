package factory

import (
	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/types"
)

func CreateCounter(name string, labels map[string]string, opts options.Options) *types.Counter {
	return types.NewCounter(formatMetric(name, opts.Apply(labels)))
}

func GetOrCreateCounter(name string, labels map[string]string, opts options.Options) *types.Counter {
	return types.GetOrCreateCounter(formatMetric(name, opts.Apply(labels)))
}

func CreateErrorCounter(name string, labels map[string]string) *types.Counter {
	return CreateCounter(name, labels, options.Error)
}

func CreateWarningCounter(name string, labels map[string]string) *types.Counter {
	return CreateCounter(name, labels, options.Warning)
}

func CreateInfoCounter(name string, labels map[string]string) *types.Counter {
	return CreateCounter(name, labels, options.Info)
}

func GetOrCreateErrorCounter(name string, labels map[string]string) *types.Counter {
	return GetOrCreateCounter(name, labels, options.Error)
}

func GetOrCreateWarningCounter(name string, labels map[string]string) *types.Counter {
	return GetOrCreateCounter(name, labels, options.Warning)
}

func GetOrCreateInfoCounter(name string, labels map[string]string) *types.Counter {
	return GetOrCreateCounter(name, labels, options.Info)
}
