package slices

import "reflect"

// ReduceFunc type reduce function
type ReduceFunc func(sum interface{}, entity interface{}, index int, slice interface{}) (interface{}, error)

// Reduce map/reduce
func Reduce(handle ReduceFunc, slice interface{}, initValue interface{}) (interface{}, error) {
	//
	var vtype reflect.Type = reflect.TypeOf(initValue)
	var ret reflect.Value = reflect.New(vtype)
	// initialize with value copy
	ret.Set(reflect.ValueOf(initValue))

	if IsSlice(slice) {
		sliceValues := reflect.ValueOf(slice)
		sliceLength := sliceValues.Len()
		// theSlice := reflect.ValueOf(slice)
		for i := 0; i < sliceLength; i++ {
			if ret, err := handle(ret, sliceValues.Index(i), i, slice); err != nil {
				return ret, err
			}
		}
	} else {
		return ret, ErrSliceType
	}

	return ret, nil

}
