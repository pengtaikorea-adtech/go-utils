package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEach(t *testing.T) {
	sample := []string{
		"hello",
		"pengtai",
		"korea",
		"adtech",
		"rules",
	}

	Each(func(e reflect.Value, i int, s interface{}) error {
		fmt.Println(e)
		return nil
	}, sample)
}
