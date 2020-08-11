package slices

import "reflect"

// FilterFunc param on filter
type FilterFunc func(entity interface{}, index int, slice interface{}) bool

// Filter slice; filter only true
func Filter(handle FilterFunc, slice interface{}) (interface{}, error) {
	//
	var rets reflect.Value
	if isSlice := IsSlice(slice); isSlice {
		sliceType := reflect.TypeOf(slice)
		sliceValues := reflect.ValueOf(slice)
		sliceLength := sliceValues.Len()

		// theSlice := reflect.ValueOf(slice)
		rets = reflect.MakeSlice(sliceType, 0, sliceValues.Cap())
		for i := 0; i < sliceLength; i++ {
			if val := sliceValues.Index(i); handle(val, i, slice) {
				rets = reflect.Append(rets, val)
			}
		}
	} else {
		return rets, ErrSliceType
	}

	return rets, nil
}
