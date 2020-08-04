package binds

import (
	"reflect"
	"testing"
)

type testBaseStruct struct {
	Name  string `name:"Name"`
	Value int    `v:"value" name:"Value" type:"int"`
	Flag  bool
}

type testCollection struct {
	testBaseStruct
	List []string
	Hash map[string]interface{}
}

type testNestedInstance struct {
	testCollection
}

var testBases = []testBaseStruct{
	{"Hello", 1, true},
	{"World", 2, false},
	{"Hi", 3, false},
}

// TestBindingStruct test golang reflect binding
func TestBindingStruct(t *testing.T) {

	for _, base := range testBases {
		// retrieve reflect value
		rv := reflect.TypeOf(base)
		for fi := 0; fi < rv.NumField(); fi++ {
			rf := rv.Field(fi)
			t.Logf("%s %s %s", rf.Name, rf.Type, rf.Tag)
		}
	}

}
