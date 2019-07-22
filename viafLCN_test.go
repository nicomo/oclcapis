package oclcapis

import (
	"reflect"
	"testing"
)

type viafGetLCNTest struct {
	Input      string
	Expected   string
	ShouldFail bool
}

var viafLCNTests = []viafGetLCNTest{
	{
		Input:      "96731408", // JM BONNISSEAU
		Expected:   "n2009050322",
		ShouldFail: false,
	},
	{
		Input:      "213067771", // N MORIN, NO RESULT
		Expected:   "",
		ShouldFail: true,
	},
	{
		Input:      "101833644", // C. BROOKE-ROSE
		Expected:   "n50048876",
		ShouldFail: false,
	},
	{
		Input:      "", // WRONG INPUT
		Expected:   "",
		ShouldFail: true,
	},
}

func TestViafGetLCN(t *testing.T) {

	for _, test := range viafLCNTests {
		actual, err := ViafGetLCN(test.Input)
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
