package options

var (
	Info       = NewOptions().WithLevel(MetricLevelInfo)
	Warning    = NewOptions().WithLevel(MetricLevelWarning)
	Error      = NewOptions().WithLevel(MetricLevelError)
	InfoAvg    = Info.WithAggregation(MetricAggregationAvg)
	InfoMin    = Info.WithAggregation(MetricAggregationMin)
	InfoMax    = Info.WithAggregation(MetricAggregationMax)
	ErrorAvg   = Error.WithAggregation(MetricAggregationAvg)
	ErrorMin   = Error.WithAggregation(MetricAggregationMin)
	ErrorMax   = Error.WithAggregation(MetricAggregationMax)
	WarningAvg = Warning.WithAggregation(MetricAggregationAvg)
	WarningMin = Warning.WithAggregation(MetricAggregationMin)
	WarningMax = Warning.WithAggregation(MetricAggregationMax)
)

type Options struct {
	Level       MetricLevel
	Aggregation *MetricAggregation
}

func NewOptions() Options {
	return Options{
		Level:       MetricLevelInfo,
		Aggregation: nil,
	}
}

func (o Options) WithLevel(level MetricLevel) Options {
	o.Level = level
	return o
}

func (o Options) WithAggregation(aggregation MetricAggregation) Options {
	o.Aggregation = &aggregation
	return o
}

func (o Options) Apply(labels map[string]string) map[string]string {
	if labels == nil {
		labels = make(map[string]string)
	}

	o.Level.Apply(labels)
	if o.Aggregation != nil {
		o.Aggregation.Apply(labels)
	}

	return labels
}
