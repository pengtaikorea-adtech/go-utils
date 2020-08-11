package slices

import (
	"reflect"
)

// MapFunc for map.
type MapFunc func(entity interface{}, index int, slice interface{}) (interface{}, error)

// Map slice; slice to another slice that of slice => []outElementType;
// length preserved
func Map(handle MapFunc, slice interface{}, outElementType reflect.Type) (interface{}, error) {
	//
	var rets reflect.Value
	if IsSlice(slice) {
		sliceType := reflect.SliceOf(outElementType)
		sliceValues := reflect.ValueOf(slice)
		sliceLength := sliceValues.Len()
		// theSlice := reflect.ValueOf(slice)
		rets = reflect.MakeSlice(sliceType, sliceValues.Len(), sliceValues.Len())
		for i := 0; i < sliceLength; i++ {
			if val, err := handle(sliceValues.Index(i), i, slice); err == nil {
				rets.Index(i).Set(reflect.ValueOf(val))
			} else {
				return rets, err
			}
		}
	} else {
		return rets, ErrSliceType
	}

	return rets, nil

}
