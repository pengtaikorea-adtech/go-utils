package slices

// EveryFunc is actually equvalent to FilterFunc
type EveryFunc func(entity interface{}, index int, slice interface{}) bool

// Every - test every entity
func Every(handle EveryFunc, slice interface{}) (bool, error) {
	rs, err := Reduce(func(sub interface{}, e interface{}, i int, s interface{}) (interface{}, error) {
		return (sub.(bool) && handle(e, i, s)), nil
	}, slice, true)

	if err != nil {
		return false, err
	}

	return rs.(bool), nil
}
