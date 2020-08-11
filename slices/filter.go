package slices

import (
	"reflect"
)

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

		rets = reflect.MakeSlice(sliceType, 0, sliceValues.Cap())
		for i := 0; i < sliceLength; i++ {
			val := sliceValues.Index(i)
			filter := handle(val.Interface(), i, slice)
			if filter {
				rets = reflect.Append(rets, val)
			}
		}
	} else {
		return rets.Interface(), ErrSliceType
	}
	return rets.Interface(), nil
}
