package metrics

import (
	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/set"
)

type FloatCounter = set.FloatCounter

func CreateFloatCounter(name string, labels map[string]string, opts options.Options) *FloatCounter {
	return defaultSet.CreateFloatCounter(name, labels, opts)
}

func GetOrCreateFloatCounter(name string, labels map[string]string, opts options.Options) *FloatCounter {
	return defaultSet.GetOrCreateFloatCounter(name, labels, opts)
}

func CreateErrorFloatCounter(name string, labels map[string]string) *FloatCounter {
	return CreateFloatCounter(name, labels, options.Error)
}

func CreateWarningFloatCounter(name string, labels map[string]string) *FloatCounter {
	return CreateFloatCounter(name, labels, options.Warning)
}

func CreateInfoFloatCounter(name string, labels map[string]string) *FloatCounter {
	return CreateFloatCounter(name, labels, options.Info)
}

func GetOrCreateErrorFloatCounter(name string, labels map[string]string) *FloatCounter {
	return GetOrCreateFloatCounter(name, labels, options.Error)
}

func GetOrCreateWarningFloatCounter(name string, labels map[string]string) *FloatCounter {
	return GetOrCreateFloatCounter(name, labels, options.Warning)
}

func GetOrCreateInfoFloatCounter(name string, labels map[string]string) *FloatCounter {
	return GetOrCreateFloatCounter(name, labels, options.Info)
}
