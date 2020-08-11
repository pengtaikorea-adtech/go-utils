package slices

import (
	"testing"
)

func TestFilter(t *testing.T) {
	// test sample
	sample := make([]int, 20)
	for i := range sample {
		sample[i] = i
	}

	// take odds
	ret, err := Filter(func(e interface{}, i int, s interface{}) bool {
		if v, ok := e.(int); ok {
			return 0 < v%2
		}
		return false
	}, sample)

	if err != nil {
		t.Error(err)
	}

	// type conversion please
	if odds, ok := ret.([]int); !ok {
		t.Error("type assertion failed")
	} else if len(sample) <= len(odds) {
		t.Error("filtered length matched?")
	} else {
		t.Log(odds)
	}
}
