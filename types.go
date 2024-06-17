package metrics

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/VictoriaMetrics/metrics"
)

// this file exposes the metrics types from VictoriaMetrics/metrics package
// the Counter type is replaced with LazyCounter to avoid registering unused metrics

type Histogram = metrics.Histogram

func NewHistogram(name string) *Histogram {
	return metrics.NewHistogram(name)
}
func GetOrCreateHistogram(name string) *Histogram {
	return metrics.GetOrCreateHistogram(name)
}

type Gauge = metrics.Gauge

func NewGauge(name string, f func() float64) *Gauge {
	return metrics.NewGauge(name, f)
}
func GetOrCreateGauge(name string, f func() float64) *Gauge {
	return metrics.GetOrCreateGauge(name, f)
}

type Summary = metrics.Summary

func NewSummary(name string) *Summary {
	return metrics.NewSummary(name)
}

func NewSummaryExt(name string, window time.Duration, quantiles []float64) *Summary {
	return metrics.NewSummaryExt(name, window, quantiles)
}

func GetOrCreateSummaryExt(name string, window time.Duration, quantiles []float64) *Summary {
	return metrics.GetOrCreateSummaryExt(name, window, quantiles)
}

type FloatCounter = metrics.FloatCounter

func NewFloatCounter(name string) *FloatCounter {
	return metrics.NewFloatCounter(name)
}
func GetOrCreateFloatCounter(name string) *FloatCounter {
	return metrics.GetOrCreateFloatCounter(name)
}

func WritePrometheus(w io.Writer, exposeProcessMetrics bool) {
	metrics.WritePrometheus(w, exposeProcessMetrics)
}

var (
	// a set of all lazy counters. We do not support removing counters from this set.
	lazyCountersSet     = make(map[string]*LazyCounter)
	lazyCountersSetLock sync.Mutex
)

// LazyCounter is a counter that is lazily initialized when it is first used,
// to avoid registering unused metrics.
// It is safe to use from concurrent goroutines.
// Note: a rare race-condition can cause data loss is multiple actions are performed on the counter
// when it is not initialized yet.
type LazyCounter struct {
	name  string
	inner atomic.Pointer[metrics.Counter]
}

type Counter = LazyCounter

// NewCounter creates a new LazyCounter with the given name.
// If a counter with the given name already exists, the program panics.
func NewCounter(name string) *LazyCounter {
	lazyCountersSetLock.Lock()
	defer lazyCountersSetLock.Unlock()

	if _, ok := lazyCountersSet[name]; ok {
		panic(fmt.Errorf("lazy counter with name %s already exists", name))
	}

	return newCounterUnsafe(name)
}

func GetOrCreateCounter(name string) *LazyCounter {
	lazyCountersSetLock.Lock()
	defer lazyCountersSetLock.Unlock()

	if c, ok := lazyCountersSet[name]; ok {
		return c
	}

	return newCounterUnsafe(name)
}

func (mc *LazyCounter) Inc() {
	mc.setActiveIfNeeded()
	mc.inner.Load().Inc()
}

func (mc *LazyCounter) Dec() {
	mc.setActiveIfNeeded()
	mc.inner.Load().Dec()
}

func (mc *LazyCounter) Get() uint64 {
	if !mc.IsActive() {
		return 0

	}
	return mc.inner.Load().Get()
}

func (mc *LazyCounter) Set(n uint64) {
	mc.setActiveIfNeeded()
	mc.inner.Load().Set(n)
}

func (mc *LazyCounter) Add(n int) {
	mc.setActiveIfNeeded()
	mc.inner.Load().Add(n)
}

func (mc *LazyCounter) IsActive() bool {
	return mc.inner.Load() != nil
}

func (mc *LazyCounter) setActiveIfNeeded() {
	if mc.IsActive() {
		return
	}
	counter := metrics.GetOrCreateCounter(mc.name)
	mc.inner.Store(counter)
}

func newCounterUnsafe(name string) *LazyCounter {
	lazyCountersSet[name] = &LazyCounter{name: name, inner: atomic.Pointer[metrics.Counter]{}}
	return lazyCountersSet[name]
}
