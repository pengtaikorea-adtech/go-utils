package slices

import "reflect"

// ReduceFunc type reduce function
type ReduceFunc func(subTotal reflect.Value, entity reflect.Value, index int, slice interface{}) (interface{}, error)

// Reduce map/reduce
func Reduce(handle ReduceFunc, slice interface{}, initValue interface{}) (reflect.Value, error) {
	//
	var ret = reflect.New(reflect.TypeOf(initValue))
	reflect.Indirect(ret).Set(reflect.ValueOf(initValue))

	if IsSlice(slice) {
		sliceValues := reflect.ValueOf(slice)
		sliceLength := sliceValues.Len()
		// theSlice := reflect.ValueOf(slice)
		for i := 0; i < sliceLength; i++ {
			if val, err := handle(reflect.Indirect(ret), sliceValues.Index(i), i, slice); err == nil {
				reflect.Indirect(ret).Set(reflect.ValueOf(val))
			} else {
				return reflect.Indirect(ret), err
			}
		}
	} else {
		return reflect.Indirect(ret), ErrSliceType
	}

	return reflect.Indirect(ret), nil

}
