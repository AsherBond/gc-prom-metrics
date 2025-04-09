package metrics

import (
	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/set"
)

func CreateSummary(
	name string,
	labels map[string]string,
	opts options.Options,
	summaryOpts options.SummaryOptions,
) *set.Summary {
	return defaultSet.CreateSummary(name, labels, opts, summaryOpts)
}

func GetOrCreateSummary(
	name string,
	labels map[string]string,
	opts options.Options,
	summaryOpts options.SummaryOptions,
) *set.Summary {
	return defaultSet.GetOrCreateSummary(name, labels, opts, summaryOpts)
}

func CreateErrorSummary(name string, labels map[string]string) *set.Summary {
	return CreateSummary(name, labels, options.Error, options.NewSummaryOptions())
}

func CreateWarningSummary(name string, labels map[string]string) *set.Summary {
	return CreateSummary(name, labels, options.Warning, options.NewSummaryOptions())
}

func CreateInfoSummary(name string, labels map[string]string) *set.Summary {
	return CreateSummary(name, labels, options.Info, options.NewSummaryOptions())
}

func CreateErrorSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *set.Summary {
	return CreateSummary(name, labels, options.Error, opts)
}

func CreateWarningSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *set.Summary {
	return CreateSummary(name, labels, options.Warning, opts)
}

func CreateInfoSummaryWithOptions(name string, labels map[string]string, opts options.SummaryOptions) *set.Summary {
	return CreateSummary(name, labels, options.Info, opts)
}

func GetOrCreateErrorSummary(name string, labels map[string]string) *set.Summary {
	return GetOrCreateSummary(name, labels, options.Error, options.NewSummaryOptions())
}

func GetOrCreateWarningSummary(name string, labels map[string]string) *set.Summary {
	return GetOrCreateSummary(name, labels, options.Warning, options.NewSummaryOptions())
}

func GetOrCreateInfoSummary(name string, labels map[string]string) *set.Summary {
	return GetOrCreateSummary(name, labels, options.Info, options.NewSummaryOptions())
}

func GetOrCreateErrorSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *set.Summary {
	return GetOrCreateSummary(name, labels, options.Error, opts)
}

func GetOrCreateWarningSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *set.Summary {
	return GetOrCreateSummary(name, labels, options.Warning, opts)
}

func GetOrCreateInfoSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *set.Summary {
	return GetOrCreateSummary(name, labels, options.Info, opts)
}
