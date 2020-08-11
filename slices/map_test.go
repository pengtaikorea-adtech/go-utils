package slices

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

func TestSliceMap(t *testing.T) {
	// build random int
	bytes := make([]byte, 10)
	io.ReadFull(rand.Reader, bytes)
	sample := make([]int, len(bytes))
	for i, b := range bytes {
		sample[i] = int(b)
	}

	puts, err := Map(func(entity interface{}, index int, slice interface{}) (interface{}, error) {
		return fmt.Sprintf("%d", entity), nil
	}, sample, ElementTypeString)

	if err == nil {
		t.Log(puts)
	} else {
		t.Error(err)
	}

}
