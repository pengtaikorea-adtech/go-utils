package slices

import "testing"

func TestFilter(t *testing.T) {
	// test sample
	sample := make([]int, 100)
	for i := range sample {
		sample[i] = i
	}

	// take odds
	rets, err := Filter(func(e interface{}, i int, s interface{}) bool {
		v := e.(int)
		return 0 < v%2
	}, sample)
	odds, _ := rets.([]int)
	if err != nil {
		t.Error(err)
	}

	t.Error(len(sample) >= len(odds))
	t.Log(odds)
}
