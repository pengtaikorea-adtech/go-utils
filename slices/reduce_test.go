package slices

import (
	"testing"
)

func TestReduce(t *testing.T) {
	sample := make([]int, 100)
	for i := range sample {
		sample[i] = i + 1
	}

	total, err := Reduce(func(t interface{}, e interface{}, i int, s interface{}) (interface{}, error) {
		// initValue and the return type should be matched
		subTotal := t.(int)
		entity := e.(int)
		return subTotal + entity, nil
	}, sample, 0)
	if err != nil {
		t.Error(err)
	} else if total.(int) != 5050 {
		t.Error("total not exact")
	}
}
