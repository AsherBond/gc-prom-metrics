package set

import (
	"github.com/VictoriaMetrics/metrics"
)

type Set struct {
	set *metrics.Set
}

func NewSet() *Set {
	return &Set{
		set: metrics.NewSet(),
	}
}
