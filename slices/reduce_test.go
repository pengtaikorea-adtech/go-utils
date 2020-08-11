package slices

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	sample := make([]int, 100)
	for i := range sample {
		sample[i] = i + 1
	}

	total, err := Reduce(func(t reflect.Value, e reflect.Value, i int, s interface{}) (interface{}, error) {
		// initValue and the return type should be matched
		return int(t.Int() + e.Int()), nil
	}, sample, 0)
	if err != nil {
		t.Error(err)
	} else if total.Int() != 5050 {
		t.Error("total not exact")
	}
}
