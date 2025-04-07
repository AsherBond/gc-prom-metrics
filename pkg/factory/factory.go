package factory

import (
	"sync/atomic"

	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/types"
)

func CreateCounter(name string, labels map[string]string, opts options.Options) *types.Counter {
	return types.NewCounter(formatMetric(name, opts.Apply(labels)))
}

func CreateGauge(name string, labels map[string]string, f func() float64, opts options.Options) *types.Gauge {
	return types.NewGauge(formatMetric(name, opts.Apply(labels)), f)
}

func CreateHistogram(name string, labels map[string]string, opts options.Options) *types.Histogram {
	return types.NewHistogram(formatMetric(name, opts.Apply(labels)))
}

func CreateSummary(
	name string,
	labels map[string]string,
	opts options.Options,
	summaryOpts options.SummaryOptions,
) *types.Summary {
	return types.NewSummaryExt(formatMetric(name, opts.Apply(labels)), summaryOpts.Window, summaryOpts.Quantiles)
}

func GetOrCreateCounter(name string, labels map[string]string, opts options.Options) *types.Counter {
	return types.GetOrCreateCounter(formatMetric(name, opts.Apply(labels)))
}

func GetOrCreateGauge(
	name string,
	labels map[string]string,
	f func() float64,
	opts options.Options,
) *types.Gauge {
	return types.GetOrCreateGauge(formatMetric(name, opts.Apply(labels)), f)
}

func GetOrCreateHistogram(name string, labels map[string]string, opts options.Options) *types.Histogram {
	return types.GetOrCreateHistogram(formatMetric(name, opts.Apply(labels)))
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

func CreateErrorGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return CreateGauge(name, labels, f, options.Error)
}

func CreateWarningGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return CreateGauge(name, labels, f, options.Warning)
}

func CreateInfoGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return CreateGauge(name, labels, f, options.Info)
}

func GetOrCreateErrorGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return GetOrCreateGauge(name, labels, f, options.Error)
}

func GetOrCreateWarningGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return GetOrCreateGauge(name, labels, f, options.Warning)
}

func GetOrCreateInfoGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return GetOrCreateGauge(name, labels, f, options.Info)
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

// Subtract from atomic uint variable with 64 bits.
// Useful for unsigned gauges.
func SubtractAtomicUint64(variable *atomic.Uint64, delta uint64) {
	variable.Add(^uint64(delta - 1))
}
