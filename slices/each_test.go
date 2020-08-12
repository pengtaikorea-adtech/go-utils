package slices

import (
	"crypto/rand"
	"fmt"
	"io"
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

// TestEmbedEachContext - test if embeding scope available
func TestEmbedEachContext(t *testing.T) {
	values := []int{}
	sample := make([]byte, 100)
	io.ReadFull(rand.Reader, sample)

	Each(func(e interface{}, i int, s interface{}) error {
		v := int(e.(byte))
		values = append(values, v)
		return nil
	}, sample)

	for i := range sample {
		if values[i] != int(sample[i]) {
			t.Errorf("expected %d but %d", int(sample[i]), values[i])
		}
	}
}
