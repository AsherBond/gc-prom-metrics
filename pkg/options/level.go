package options

import "errors"

const (
	metricLevelLabel = "__metric_level__"
)

var (
	ErrInvalidMetricLevel = errors.New("invalid metric level")
)

type MetricLevel string

const (
	MetricLevelError   MetricLevel = "error"
	MetricLevelInfo    MetricLevel = "info"
	MetricLevelWarning MetricLevel = "warning"
)

func (level MetricLevel) Apply(labels map[string]string) {
	labels[metricLevelLabel] = string(level)
}

func MetricLevelFromString(str string) (MetricLevel, error) {
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
