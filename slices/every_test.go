package slices

import "testing"

func TestEveryone(t *testing.T) {

	// types
	type tester struct {
		handle    EveryFunc
		theSlice  []int
		expecting bool
		errorMsg  string
	}

	biggerThanZero := func(e interface{}, i int, s interface{}) bool {
		return 0 < e.(int)
	}

	isEvenNumber := func(e interface{}, i int, s interface{}) bool {
		return (e.(int)%2 == 0)
	}
	numbers := make([]int, 100)
	evens := make([]int, len(numbers))

	for i := range numbers {
		numbers[i] = i + 1
		evens[i] = (2 * numbers[i])
	}

	testers := []tester{
		{biggerThanZero, numbers, true, "nubmers le 0"},
		{biggerThanZero, evens, true, "evens le 0"},
		{isEvenNumber, numbers, false, "not every numbers are even"},
	}

	for _, ts := range testers {
		if r, e := Every(ts.handle, ts.theSlice); e == nil {
			if r != ts.expecting {
				t.Error(ts.errorMsg)
			}
		} else {
			t.Error("Err!! " + ts.errorMsg)
		}
	}
}
