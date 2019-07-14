package oclcapis

import (
	"reflect"
	"testing"
)

type viafGetWKPTest struct {
	Input      string
	Expected   string
	ShouldFail bool
}

func TestViafGetWKP(t *testing.T) {
	var viafWKPTests = []viafGetWKPTest{
		{
			Input:      "96731408", // JM BONNISSEAU
			Expected:   "Q30084598",
			ShouldFail: false,
		},
		{
			Input:      "",
			Expected:   "",
			ShouldFail: true,
		},
	}
	for _, test := range viafWKPTests {
		actual, err := ViafGetWKP(test.Input)
		if err != nil {
			if test.ShouldFail {
				t.Logf("PASS: got expected error %v", err)
			} else {
				t.Fatalf("FAIL for %s: expected %v, got an error %v", test.Input, test.Expected, err)
			}
		}
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Input, test.Expected, actual)
		}
	}
}
