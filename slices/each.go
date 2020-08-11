package slices

import "reflect"

// EachFunc param on ForEach
type EachFunc func(entity reflect.Value, index int, slice interface{}) error

// Each slice; run ForEach
func Each(handle EachFunc, slice interface{}) error {
	if IsSlice(slice) {
		sliceValues := reflect.ValueOf(slice)
		sliceLength := sliceValues.Len()
		for i := 0; i < sliceLength; i++ {
			if err := handle(sliceValues.Index(i), i, slice); err != nil {
				return err
			}
		}
	} else {
		return ErrSliceType
	}
	return nil
}
