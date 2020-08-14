package binds

import (
	"fmt"
	"reflect"
)

// TagKey "binds" struct field tag keyname for binding
const TagKey = "binds"

var mappingCache = make(map[string][]string, 0)

func registerMapping(typePath string, targetType reflect.Type) {
}

func getMapDictionary(instance interface{}) (reflect.Type, []string) {
	targetType := reflect.TypeOf(instance)
	typePath := fmt.Sprintf("%s/%s", targetType.PkgPath(), targetType.Name())

	if _, exists := mappingCache[typePath]; !exists {
		numFields := targetType.NumField()
		dict := make([]string, numFields)
		for i := 0; i < numFields; i++ {
			field := targetType.Field(i)
			if key, has := field.Tag.Lookup(TagKey); has {
				dict[i] = key
			} else {
				dict[i] = ""
			}
		}
		mappingCache[typePath] = dict
	}

	return targetType, mappingCache[typePath]
}

// Populate create new instance as of type Bindable
func Populate(data map[string]interface{}, templateInstance interface{}) interface{} {
	targetType, dict := getMapDictionary(templateInstance)
	theInstance := reflect.Indirect(reflect.New(targetType))

	for i, key := range dict {
		if val, exists := data[key]; exists {
			theInstance.Field(i).Set(reflect.ValueOf(val))
		}
	}

	return theInstance.Interface()
}

// Serialize build mapping dictionary
func Serialize(instance interface{}) map[string]interface{} {
	_, dict := getMapDictionary(instance)
	targetValue := reflect.ValueOf(instance)

	data := make(map[string]interface{}, 0)

	for i, key := range dict {
		if 0 < len(key) {
			data[key] = targetValue.Field(i).Interface()
		}
	}

	return data
}
