package slices

import "testing"

func testSlice(t *testing.T, arg interface{}, elem string) {
	expected := 0 < len(elem)
	if ok := IsSlice(arg); ok == expected {
		if expected {
			t.Logf("[]%s slice passed", elem)
		} else {
			t.Log("non-slice passed")
		}
	} else {
		t.Errorf("expecting %t but %t, %s", expected, ok, arg.(string))
	}
}

func TestSliceIsSlice(t *testing.T) {
	intSlice := []int{}
	stringSlice := []string{"hello", "hi"}
	dynamicInstance := struct {
		A string
		B int
	}{"world", 0}
	any := false

	testSlice(t, intSlice, "int")
	testSlice(t, stringSlice, "string")
	testSlice(t, dynamicInstance, "")
	testSlice(t, any, "")
	testSlice(t, testSlice, "")

}
