package slices

import (
	"errors"
	"reflect"
)

// ElementTypeBool - outElementType bool
var ElementTypeBool = reflect.TypeOf(true)

// ElementTypeInt - outElementType int
var ElementTypeInt = reflect.TypeOf(0)

// ElementTypeString - outElementType string
var ElementTypeString = reflect.TypeOf("")

// ErrSliceType - input element is not a slice type
var ErrSliceType = errors.New("the type is not a slice")

// IsSlice - test if the value is slice
func IsSlice(target interface{}) bool {
	return reflect.TypeOf(target).Kind() == reflect.Slice

}
