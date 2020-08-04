package binds

import (
	"errors"
	"reflect"
	"strconv"
	"unsafe"
)

// Map - declare data map on any
type Map map[string]string

// BindingKey - binding tag key
const BindingKey = "binds"

// ErrType - type conversion Error
var ErrType = errors.New("Unsupporting type conversion")

type fieldIterationMapFn func(*reflect.StructField, *string, *reflect.Value) error

type fieldIterationReturnFn func(*Map, unsafe.Pointer) error

func fieldIteration(iterFn fieldIterationMapFn) fieldIterationReturnFn {
	return func(data *Map, targetPtr unsafe.Pointer) error {
		targetValue := reflect.ValueOf(targetPtr)
		targetType := reflect.TypeOf(targetValue)
		lenFields := targetValue.NumField()
		for fi := 0; fi < lenFields; fi++ {
			// TODO: iterFn 여기서부터
			fType := targetType.Field(fi)
			bindingKey, found := fType.Tag.Lookup(BindingKey)
			mapValue, exists := (*data)[bindingKey]
			if !(found && exists) {
				continue
			}

			fActual := targetValue.Field(fi)
			if err := iterFn(&fType, &mapValue, &fActual); err != nil {
				return err
			}
		}
		return nil
	}

}

// Populate - map data to instance
var Populate = fieldIteration(populateIterFn)

func populateIterFn(fType *reflect.StructField, mapValue *string, fActual *reflect.Value) error {
	if !fActual.CanSet() {
		return nil
	}

	switch fActual.Kind() {
	case reflect.Bool:
		if x, er := strconv.ParseBool(*mapValue); er == nil {
			fActual.SetBool(x)
		} else {
			return er
		}
	case reflect.Int:
		if x, er := strconv.ParseInt(*mapValue, 10, 64); er == nil {
			fActual.SetInt(x)
		} else {
			return er
		}
	case reflect.Float32:
		if x, er := strconv.ParseFloat(*mapValue, 32); er == nil {
			fActual.SetFloat(x)
		} else {
			return er
		}
	case reflect.Float64:
		if x, er := strconv.ParseFloat(*mapValue, 64); er == nil {
			fActual.SetFloat(x)
		} else {
			return er
		}
	case reflect.String:
		fActual.SetString(*mapValue)
	default:
		return ErrType

	}
	return nil
}

// Mapping - instace to map data
func mappingIterFn(fType *reflect.StructField, mapValue *string, fActual *reflect.Value) error {
	return
}
