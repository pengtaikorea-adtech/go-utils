package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	// test sample
	sample := make([]int, 20)
	for i := range sample {
		sample[i] = i
	}

	// take odds
	odds, err := Filter(func(e reflect.Value, i int, s interface{}) bool {
		v := e.Int()
		return 0 < v%2
	}, sample)
	if err != nil {
		t.Error(err)
	}

	t.Error(len(sample) >= len(odds))
	t.Log(odds)
}
