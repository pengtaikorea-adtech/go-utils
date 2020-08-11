package slices

import (
	"fmt"
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

	Each(func(e interface{}, i int, s interface{}) error {
		fmt.Println(e)
		return nil
	}, sample)
}
