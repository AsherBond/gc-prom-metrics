package set

import (
	"sync/atomic"

	"github.com/VictoriaMetrics/metrics"
	"github.com/groundcover-com/metrics/pkg/options"
)

type Gauge = metrics.Gauge

func (set *Set) CreateGauge(
	name string,
	labels map[string]string,
	f func() float64,
	opts options.Options,
) *Gauge {
	return set.set.NewGauge(formatMetric(name, opts.Apply(labels)), f)
}

func (set *Set) GetOrCreateGauge(
	name string,
	labels map[string]string,
	f func() float64,
	opts options.Options,
) *Gauge {
	return set.set.GetOrCreateGauge(formatMetric(name, opts.Apply(labels)), f)
}

func (set *Set) CreateErrorGauge(
	name string,
	labels map[string]string,
	f func() float64,
) *Gauge {
	return set.CreateGauge(name, labels, f, options.Error)
}

func (set *Set) CreateWarningGauge(
	name string,
	labels map[string]string,
	f func() float64,
) *Gauge {
	return set.CreateGauge(name, labels, f, options.Warning)
}

func (set *Set) CreateInfoGauge(
	name string,
	labels map[string]string,
	f func() float64,
) *Gauge {
	return set.CreateGauge(name, labels, f, options.Info)
}

func (set *Set) GetOrCreateErrorGauge(
	name string,
	labels map[string]string,
	f func() float64,
) *Gauge {
	return set.GetOrCreateGauge(name, labels, f, options.Error)
}

func (set *Set) GetOrCreateWarningGauge(
	name string,
	labels map[string]string,
	f func() float64,
) *Gauge {
	return set.GetOrCreateGauge(name, labels, f, options.Warning)
}

func (set *Set) GetOrCreateInfoGauge(
	name string,
	labels map[string]string,
	f func() float64,
) *Gauge {
	return set.GetOrCreateGauge(name, labels, f, options.Info)
}

// Subtract from atomic uint variable with 64 bits.
// Useful for unsigned gauges.
func SubtractAtomicUint64(variable *atomic.Uint64, delta uint64) {
	variable.Add(^uint64(delta - 1))
}
