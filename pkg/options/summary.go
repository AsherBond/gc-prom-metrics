package options

import "time"

const (
	// https://github.com/VictoriaMetrics/metrics/blob/v1.24.0/summary.go#L14
	defaultSummaryWindow = 5 * time.Minute
)

var (
	defaultSummaryQuantiles = []float64{0.05, 0.5, 0.95, 1}
)

type SummaryOptions struct {
	Quantiles []float64
	Window    time.Duration
}

func NewSummaryOptions() SummaryOptions {
	return SummaryOptions{Quantiles: defaultSummaryQuantiles, Window: defaultSummaryWindow}
}

func (options SummaryOptions) WithQuantiles(quantiles []float64) SummaryOptions {
	options.Quantiles = quantiles
	return options
}

func (options SummaryOptions) WithWindow(window time.Duration) SummaryOptions {
	options.Window = window
	return options
}
