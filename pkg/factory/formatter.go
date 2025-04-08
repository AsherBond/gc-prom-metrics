package factory

import (
	"slices"
	"strings"

	"github.com/groundcover-com/metrics/pkg/set"
)

var (
	defaultSet = set.NewSet()
)

func formatMetric(metric string, labels map[string]string) string {
	labelsSorted := make([]string, 0, len(labels))
	for labelName := range labels {
		if labelName != "" {
			labelsSorted = append(labelsSorted, labelName)
		}
	}

	slices.Sort(labelsSorted)

	labelPairs := make([]string, 0, len(labels))
	for _, labelName := range labelsSorted {
		labelValue := labels[labelName]
		if labelValue != "" {
			labelPairs = append(labelPairs, labelName+`="`+strings.ReplaceAll(labelValue, `"`, `\"`)+`"`)
		}
	}

	return metric + "{" + strings.Join(labelPairs, ",") + "}"
}
