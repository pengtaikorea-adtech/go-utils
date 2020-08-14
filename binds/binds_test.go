package binds

import "testing"

// short form dict
type dictionary map[string]interface{}

type Tester struct {
	Text   string `binds:"text"`
	Number int    `binds:"num"`
	Flag   bool   `binds:"flag"`
	Ignore string
}

type ExpectedTester struct {
	dict dictionary
	inst Tester
}

var samples = []ExpectedTester{
	{
		dictionary{"text": "hi", "num": 1, "flag": true},
		Tester{Text: "hi", Number: 1, Flag: true},
	},
	{
		dictionary{"text": "hello", "num": 2},
		Tester{Text: "hello", Number: 2},
	},
	{
		dictionary{"text": "world"},
		Tester{Text: "world"},
	},
	{
		dictionary{"text": "hi", "num": 4, "flag": false, "ignore": "what?"},
		Tester{Text: "hi", Number: 4, Flag: false},
	},
}

func TestCachemap(t *testing.T) {
	tmpl := Tester{}

	_, mapp := getMapDictionary(tmpl)
	expected := []string{
		"text", "num", "flag", "",
	}
	for i, k := range mapp {
		if k != expected[i] {
			t.Errorf("%d expected %s, but %s", i, expected[i], k)
		}
	}
}

func TestPopulate(t *testing.T) {
	//
	tmpl := Tester{}

	for _, smp := range samples {
		pop := Populate(smp.dict, tmpl)
		if pop == nil {
			t.Error("pop nil")
		} else {
			// log expected
			t.Log(smp.inst)
			// log actual
			t.Log(pop)
		}
	}
}

func TestSerialize(t *testing.T) {
	for i, smp := range samples {
		dict := Serialize(smp.inst)
		if dict == nil {
			t.Error("ser nil")
		} else {
			for k, v := range dict {
				if exp, exists := smp.dict[k]; exists && exp != v {
					t.Errorf("%d %s expected %s but %s", i, k, exp, v)
				}
			}
		}
	}
}
