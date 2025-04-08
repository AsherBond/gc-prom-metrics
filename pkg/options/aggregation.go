package options

import "errors"

const (
	metricAggregationLabel = "__metric_aggregation__"
)

var (
	ErrInvalidMetricAggregation = errors.New("invalid metric aggregation")
)

type MetricAggregation string

const (
	MetricAggregationAvg MetricAggregation = "avg"
	MetricAggregationMin MetricAggregation = "min"
	MetricAggregationMax MetricAggregation = "max"
)

func (aggregation MetricAggregation) Apply(labels map[string]string) {
	labels[metricAggregationLabel] = string(aggregation)
}

func MetricAggregationFromString(str string) (MetricAggregation, error) {
	if str == string(MetricAggregationAvg) {
		return MetricAggregationAvg, nil
	}
	if str == string(MetricAggregationMin) {
		return MetricAggregationMin, nil
	}
	if str == string(MetricAggregationMax) {
		return MetricAggregationMax, nil
	}
	return "", ErrInvalidMetricAggregation
}
