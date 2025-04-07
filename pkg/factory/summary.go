package factory

import (
	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/types"
)

func CreateSummary(
	name string,
	labels map[string]string,
	opts options.Options,
	summaryOpts options.SummaryOptions,
) *types.Summary {
	return types.NewSummaryExt(formatMetric(name, opts.Apply(labels)), summaryOpts.Window, summaryOpts.Quantiles)
}

func GetOrCreateSummary(
	name string,
	labels map[string]string,
	opts options.Options,
	summaryOpts options.SummaryOptions,
) *types.Summary {
	return types.GetOrCreateSummaryExt(
		formatMetric(name, opts.Apply(labels)),
		summaryOpts.Window,
		summaryOpts.Quantiles,
	)
}

func CreateErrorSummary(name string, labels map[string]string) *types.Summary {
	return CreateSummary(name, labels, options.Error, options.NewSummaryOptions())
}

func CreateWarningSummary(name string, labels map[string]string) *types.Summary {
	return CreateSummary(name, labels, options.Warning, options.NewSummaryOptions())
}

func CreateInfoSummary(name string, labels map[string]string) *types.Summary {
	return CreateSummary(name, labels, options.Info, options.NewSummaryOptions())
}

func CreateErrorSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *types.Summary {
	return CreateSummary(name, labels, options.Error, opts)
}

func CreateWarningSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *types.Summary {
	return CreateSummary(name, labels, options.Warning, opts)
}

func CreateInfoSummaryWithOptions(name string, labels map[string]string, opts options.SummaryOptions) *types.Summary {
	return CreateSummary(name, labels, options.Info, opts)
}

func GetOrCreateErrorSummary(name string, labels map[string]string) *types.Summary {
	return GetOrCreateSummary(name, labels, options.Error, options.NewSummaryOptions())
}

func GetOrCreateWarningSummary(name string, labels map[string]string) *types.Summary {
	return GetOrCreateSummary(name, labels, options.Warning, options.NewSummaryOptions())
}

func GetOrCreateInfoSummary(name string, labels map[string]string) *types.Summary {
	return GetOrCreateSummary(name, labels, options.Info, options.NewSummaryOptions())
}

func GetOrCreateErrorSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *types.Summary {
	return GetOrCreateSummary(name, labels, options.Error, opts)
}

func GetOrCreateWarningSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *types.Summary {
	return GetOrCreateSummary(name, labels, options.Warning, opts)
}

func GetOrCreateInfoSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *types.Summary {
	return GetOrCreateSummary(name, labels, options.Info, opts)
}
