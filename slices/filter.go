package slices

import (
	"fmt"
	"reflect"
)

// FilterFunc param on filter
type FilterFunc func(entity reflect.Value, index int, slice interface{}) bool

// Filter slice; filter only true
func Filter(handle FilterFunc, slice interface{}) (reflect.Value, error) {
	//
	var rets reflect.Value
	if isSlice := IsSlice(slice); isSlice {
		sliceType := reflect.TypeOf(slice)
		sliceValues := reflect.ValueOf(slice)
		sliceLength := sliceValues.Len()

		// theSlice := reflect.ValueOf(slice)
		rets = reflect.MakeSlice(sliceType, 0, sliceValues.Cap())
		for i := 0; i < sliceLength; i++ {
			val := sliceValues.Index(i)
			filter := handle(val, i, slice)
			// fmt.Println(val)
			// fmt.Println(filter)
			if filter {
				rets = reflect.Append(rets, val)
			}
			fmt.Println(rets)
		}
	} else {
		return rets, ErrSliceType
	}
	return rets, nil
}
