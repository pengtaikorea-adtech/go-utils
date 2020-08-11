package slices

import "reflect"

// ReduceFunc type reduce function
type ReduceFunc func(subTotal interface{}, entity interface{}, index int, slice interface{}) (interface{}, error)

// Reduce map/reduce
func Reduce(handle ReduceFunc, slice interface{}, initValue interface{}) (interface{}, error) {
	//
	var ret = reflect.New(reflect.TypeOf(initValue))
	reflect.Indirect(ret).Set(reflect.ValueOf(initValue))

	if IsSlice(slice) {
		sliceValues := reflect.ValueOf(slice)
		sliceLength := sliceValues.Len()
		// theSlice := reflect.ValueOf(slice)
		for i := 0; i < sliceLength; i++ {
			if val, err := handle(interfaced(ret), sliceValues.Index(i).Interface(), i, slice); err == nil {
				reflect.Indirect(ret).Set(reflect.ValueOf(val))
			} else {
				return interfaced(ret), err
			}
		}
	} else {
		return interfaced(ret), ErrSliceType
	}

	return interfaced(ret), nil
}

func interfaced(ptr reflect.Value) interface{} {
	return reflect.Indirect(ptr).Interface().(interface{})
}
