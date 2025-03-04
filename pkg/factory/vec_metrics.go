package factory

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/groundcover-com/metrics/pkg/types"
)

const (
	metricLevelLabel = "metric_level"

	// https://github.com/VictoriaMetrics/metrics/blob/v1.24.0/summary.go#L14
	defaultSummaryWindow = 5 * time.Minute
)

var (
	defaultSummaryQuantiles = []float64{0.05, 0.5, 0.95, 1}
	ErrInvalidMetricLevel   = errors.New("invalid metric level")
)

type MetricLevel string

const (
	MetricLevelError   MetricLevel = "error"
	MetricLevelInfo    MetricLevel = "info"
	MetricLevelWarning MetricLevel = "warning"
)

type summaryOptions struct {
	quantiles []float64
	window    time.Duration
}

func NewSummaryOptions() summaryOptions {
	return summaryOptions{quantiles: defaultSummaryQuantiles, window: defaultSummaryWindow}
}

func (options summaryOptions) WithQuantiles(quantiles []float64) summaryOptions {
	options.quantiles = quantiles
	return options
}

func (options summaryOptions) WithWindow(window time.Duration) summaryOptions {
	options.window = window
	return options
}

func ToMetricLevel(str string) (MetricLevel, error) {
	if str == string(MetricLevelError) {
		return MetricLevelError, nil
	}
	if str == string(MetricLevelInfo) {
		return MetricLevelInfo, nil
	}
	if str == string(MetricLevelWarning) {
		return MetricLevelWarning, nil
	}
	return "", ErrInvalidMetricLevel
}

func CreateLeveledCounter(name string, level MetricLevel, labels map[string]string) *types.Counter {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[metricLevelLabel] = string(level)
	return types.NewCounter(formatMetric(name, labels))
}

func CreateLeveledGauge(name string, level MetricLevel, labels map[string]string, f func() float64) *types.Gauge {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[metricLevelLabel] = string(level)
	return types.NewGauge(formatMetric(name, labels), f)
}

func CreateLeveledHistogram(name string, level MetricLevel, labels map[string]string) *types.Histogram {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[metricLevelLabel] = string(level)
	return types.NewHistogram(formatMetric(name, labels))
}

func CreateLeveledSummary(
	name string,
	level MetricLevel,
	labels map[string]string,
	options summaryOptions,
) *types.Summary {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[metricLevelLabel] = string(level)
	return types.NewSummaryExt(formatMetric(name, labels), options.window, options.quantiles)
}

func GetOrCreateLeveledCounter(name string, level MetricLevel, labels map[string]string) *types.Counter {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[metricLevelLabel] = string(level)
	return types.GetOrCreateCounter(formatMetric(name, labels))
}

func GetOrCreateLeveledGauge(
	name string,
	level MetricLevel,
	labels map[string]string,
	f func() float64,
) *types.Gauge {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[metricLevelLabel] = string(level)
	return types.GetOrCreateGauge(formatMetric(name, labels), f)
}

func GetOrCreateLeveledHistogram(name string, level MetricLevel, labels map[string]string) *types.Histogram {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[metricLevelLabel] = string(level)
	return types.GetOrCreateHistogram(formatMetric(name, labels))
}

func GetOrCreateLeveledSummary(
	name string,
	level MetricLevel,
	labels map[string]string,
	options summaryOptions,
) *types.Summary {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[metricLevelLabel] = string(level)
	return types.GetOrCreateSummaryExt(formatMetric(name, labels), options.window, options.quantiles)
}

func CreateErrorCounter(name string, labels map[string]string) *types.Counter {
	return CreateLeveledCounter(name, MetricLevelError, labels)
}

func CreateWarningCounter(name string, labels map[string]string) *types.Counter {
	return CreateLeveledCounter(name, MetricLevelWarning, labels)
}

func CreateInfoCounter(name string, labels map[string]string) *types.Counter {
	return CreateLeveledCounter(name, MetricLevelInfo, labels)
}

func GetOrCreateErrorCounter(name string, labels map[string]string) *types.Counter {
	return GetOrCreateLeveledCounter(name, MetricLevelError, labels)
}

func GetOrCreateWarningCounter(name string, labels map[string]string) *types.Counter {
	return GetOrCreateLeveledCounter(name, MetricLevelWarning, labels)
}

func GetOrCreateInfoCounter(name string, labels map[string]string) *types.Counter {
	return GetOrCreateLeveledCounter(name, MetricLevelInfo, labels)
}

func CreateErrorGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return CreateLeveledGauge(name, MetricLevelError, labels, f)
}

func CreateWarningGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return CreateLeveledGauge(name, MetricLevelWarning, labels, f)
}

func CreateInfoGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return CreateLeveledGauge(name, MetricLevelInfo, labels, f)
}

func GetOrCreateErrorGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return GetOrCreateLeveledGauge(name, MetricLevelError, labels, f)
}

func GetOrCreateWarningGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return GetOrCreateLeveledGauge(name, MetricLevelWarning, labels, f)
}

func GetOrCreateInfoGauge(name string, labels map[string]string, f func() float64) *types.Gauge {
	return GetOrCreateLeveledGauge(name, MetricLevelInfo, labels, f)
}

func CreateErrorHistogram(name string, labels map[string]string) *types.Histogram {
	return CreateLeveledHistogram(name, MetricLevelError, labels)
}

func CreateWarningHistogram(name string, labels map[string]string) *types.Histogram {
	return CreateLeveledHistogram(name, MetricLevelWarning, labels)
}

func CreateInfoHistogram(name string, labels map[string]string) *types.Histogram {
	return CreateLeveledHistogram(name, MetricLevelInfo, labels)
}

func GetOrCreateErrorHistogram(name string, labels map[string]string) *types.Histogram {
	return GetOrCreateLeveledHistogram(name, MetricLevelError, labels)
}

func GetOrCreateWarningHistogram(name string, labels map[string]string) *types.Histogram {
	return GetOrCreateLeveledHistogram(name, MetricLevelWarning, labels)
}

func GetOrCreateInfoHistogram(name string, labels map[string]string) *types.Histogram {
	return GetOrCreateLeveledHistogram(name, MetricLevelInfo, labels)
}

func CreateErrorSummary(name string, labels map[string]string) *types.Summary {
	return CreateLeveledSummary(name, MetricLevelError, labels, NewSummaryOptions())
}

func CreateWarningSummary(name string, labels map[string]string) *types.Summary {
	return CreateLeveledSummary(name, MetricLevelWarning, labels, NewSummaryOptions())
}

func CreateInfoSummary(name string, labels map[string]string) *types.Summary {
	return CreateLeveledSummary(name, MetricLevelInfo, labels, NewSummaryOptions())
}

func CreateErrorSummaryWithOptions(name string, labels map[string]string, options summaryOptions) *types.Summary {
	return CreateLeveledSummary(name, MetricLevelError, labels, options)
}

func CreateWarningSummaryWithOptions(name string, labels map[string]string, options summaryOptions) *types.Summary {
	return CreateLeveledSummary(name, MetricLevelWarning, labels, options)
}

func CreateInfoSummaryWithOptions(name string, labels map[string]string, options summaryOptions) *types.Summary {
	return CreateLeveledSummary(name, MetricLevelInfo, labels, options)
}

func GetOrCreateErrorSummary(name string, labels map[string]string) *types.Summary {
	return GetOrCreateLeveledSummary(name, MetricLevelError, labels, NewSummaryOptions())
}

func GetOrCreateWarningSummary(name string, labels map[string]string) *types.Summary {
	return GetOrCreateLeveledSummary(name, MetricLevelWarning, labels, NewSummaryOptions())
}

func GetOrCreateInfoSummary(name string, labels map[string]string) *types.Summary {
	return GetOrCreateLeveledSummary(name, MetricLevelInfo, labels, NewSummaryOptions())
}

func GetOrCreateErrorSummaryWithOptions(
	name string,
	labels map[string]string,
	options summaryOptions,
) *types.Summary {
	return GetOrCreateLeveledSummary(name, MetricLevelError, labels, options)
}

func GetOrCreateWarningSummaryWithOptions(
	name string,
	labels map[string]string,
	options summaryOptions,
) *types.Summary {
	return GetOrCreateLeveledSummary(name, MetricLevelWarning, labels, options)
}

func GetOrCreateInfoSummaryWithOptions(
	name string,
	labels map[string]string,
	options summaryOptions,
) *types.Summary {
	return GetOrCreateLeveledSummary(name, MetricLevelInfo, labels, options)
}

// Subtract from atomic uint variable with 64 bits.
// Useful for unsigned gauges.
func SubtractAtomicUint64(variable *atomic.Uint64, delta uint64) {
	variable.Add(^uint64(delta - 1))
}
