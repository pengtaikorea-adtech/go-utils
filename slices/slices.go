package slices

import (
	"reflect"
)

// ElementTypeBool - outElementType bool
var ElementTypeBool = reflect.TypeOf(true)

// ElementTypeByte - outElementType byte
var ElementTypeByte = reflect.TypeOf(byte(0))

// ElementTypeInt - outElementType int
var ElementTypeInt = reflect.TypeOf(0)

// ElementTypeInt8 - outElementType int8
var ElementTypeInt8 = reflect.TypeOf(int8(0))

// ElementTypeInt16 - outElementType int16
var ElementTypeInt16 = reflect.TypeOf(int16(0))

// ElementTypeInt32 - outElementType int32
var ElementTypeInt32 = reflect.TypeOf(int32(0))

// ElementTypeInt64 - outElementType int64
var ElementTypeInt64 = reflect.TypeOf(int64(0))

// ElementTypeUInt8 - outElementType uint8
var ElementTypeUInt8 = reflect.TypeOf(uint8(0))

// ElementTypeUInt16 - outElementType uint16
var ElementTypeUInt16 = reflect.TypeOf(uint16(0))

// ElementTypeUInt32 - outElementType uint32
var ElementTypeUInt32 = reflect.TypeOf(uint32(0))

// ElementTypeUInt64 - outElementType uint64
var ElementTypeUInt64 = reflect.TypeOf(uint64(0))

// ElementTypeFloat32 - outElementType float32
var ElementTypeFloat32 = reflect.TypeOf(float32(0))

// ElementTypeFloat64 - outElementType float64
var ElementTypeFloat64 = reflect.TypeOf(float64(0))

// ElementTypeComplex64 - outElementType complex
var ElementTypeComplex64 = reflect.TypeOf(complex64(0))

// ElementTypeComplex128 - outElementType complex
var ElementTypeComplex128 = reflect.TypeOf(complex128(0))

// ElementTypeString - outElementType complex
var ElementTypeString = reflect.TypeOf("")

// IsSlice - test if the value is slice
func IsSlice(target interface{}) bool {
	return reflect.TypeOf(target).Kind() == reflect.Slice

}
