package metrics

import "github.com/groundcover-com/metrics/pkg/options"

var (
	Info       = options.Info
	Warning    = options.Warning
	Error      = options.Error
	InfoAvg    = options.InfoAvg
	InfoMin    = options.InfoMin
	InfoMax    = options.InfoMax
	ErrorAvg   = options.ErrorAvg
	ErrorMin   = options.ErrorMin
	ErrorMax   = options.ErrorMax
	WarningAvg = options.WarningAvg
	WarningMin = options.WarningMin
	WarningMax = options.WarningMax
)

type (
	Options        = options.Options
	SummaryOptions = options.SummaryOptions
)

func NewOptions() Options {
	return options.NewOptions()
}

func NewSummaryOptions() SummaryOptions {
	return options.NewSummaryOptions()
}
