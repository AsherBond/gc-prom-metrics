package set

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/groundcover-com/metrics/pkg/options"
)

func (set *Set) CreateSummary(
	name string,
	labels map[string]string,
	opts options.Options,
	summaryOpts options.SummaryOptions,
) *metrics.Summary {
	return set.set.NewSummaryExt(formatMetric(name, opts.Apply(labels)), summaryOpts.Window, summaryOpts.Quantiles)
}

func (set *Set) GetOrCreateSummary(
	name string,
	labels map[string]string,
	opts options.Options,
	summaryOpts options.SummaryOptions,
) *metrics.Summary {
	return set.set.GetOrCreateSummaryExt(
		formatMetric(name, opts.Apply(labels)),
		summaryOpts.Window,
		summaryOpts.Quantiles,
	)
}

func (set *Set) CreateErrorSummary(name string, labels map[string]string) *metrics.Summary {
	return set.CreateSummary(name, labels, options.Error, options.NewSummaryOptions())
}

func (set *Set) CreateWarningSummary(name string, labels map[string]string) *metrics.Summary {
	return set.CreateSummary(name, labels, options.Warning, options.NewSummaryOptions())
}

func (set *Set) CreateInfoSummary(name string, labels map[string]string) *metrics.Summary {
	return set.CreateSummary(name, labels, options.Info, options.NewSummaryOptions())
}

func (set *Set) CreateErrorSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *metrics.Summary {
	return set.CreateSummary(name, labels, options.Error, opts)
}

func (set *Set) CreateWarningSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *metrics.Summary {
	return set.CreateSummary(name, labels, options.Warning, opts)
}

func (set *Set) CreateInfoSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *metrics.Summary {
	return set.CreateSummary(name, labels, options.Info, opts)
}

func (set *Set) GetOrCreateErrorSummary(name string, labels map[string]string) *metrics.Summary {
	return set.GetOrCreateSummary(name, labels, options.Error, options.NewSummaryOptions())
}

func (set *Set) GetOrCreateWarningSummary(name string, labels map[string]string) *metrics.Summary {
	return set.GetOrCreateSummary(name, labels, options.Warning, options.NewSummaryOptions())
}

func (set *Set) GetOrCreateInfoSummary(name string, labels map[string]string) *metrics.Summary {
	return set.GetOrCreateSummary(name, labels, options.Info, options.NewSummaryOptions())
}

func (set *Set) GetOrCreateErrorSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *metrics.Summary {
	return set.GetOrCreateSummary(name, labels, options.Error, opts)
}

func (set *Set) GetOrCreateWarningSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *metrics.Summary {
	return set.GetOrCreateSummary(name, labels, options.Warning, opts)
}

func (set *Set) GetOrCreateInfoSummaryWithOptions(
	name string,
	labels map[string]string,
	opts options.SummaryOptions,
) *metrics.Summary {
	return set.GetOrCreateSummary(name, labels, options.Info, opts)
}
