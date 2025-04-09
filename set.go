package metrics

import (
	"github.com/groundcover-com/metrics/pkg/set"
)

var (
	defaultSet = NewSet()
)

type (
	Set = set.Set
)

func NewSet() *Set {
	return set.NewSet()
}
