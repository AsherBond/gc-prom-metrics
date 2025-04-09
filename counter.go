package metrics

import (
	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/set"
)

func NewCounter(name string) *set.Counter {
	return defaultSet.NewCounter(name)
}

func CreateCounter(name string, labels map[string]string, opts options.Options) *set.Counter {
	return defaultSet.CreateCounter(name, labels, opts)
}

func GetOrCreateCounter(name string, labels map[string]string, opts options.Options) *set.Counter {
	return defaultSet.GetOrCreateCounter(name, labels, opts)
}

func CreateErrorCounter(name string, labels map[string]string) *set.Counter {
	return CreateCounter(name, labels, options.Error)
}

func CreateWarningCounter(name string, labels map[string]string) *set.Counter {
	return CreateCounter(name, labels, options.Warning)
}

func CreateInfoCounter(name string, labels map[string]string) *set.Counter {
	return CreateCounter(name, labels, options.Info)
}

func GetOrCreateErrorCounter(name string, labels map[string]string) *set.Counter {
	return GetOrCreateCounter(name, labels, options.Error)
}

func GetOrCreateWarningCounter(name string, labels map[string]string) *set.Counter {
	return GetOrCreateCounter(name, labels, options.Warning)
}

func GetOrCreateInfoCounter(name string, labels map[string]string) *set.Counter {
	return GetOrCreateCounter(name, labels, options.Info)
}
