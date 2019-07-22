package oclcapis

import (
	"reflect"
	"testing"
)

type viafGetIDsTest struct {
	Input      string
	Expected   map[string]string
	ShouldFail bool
}

func TestViafGetIDs(t *testing.T) {
	var viafGetIDsTests = []viafGetIDsTest{
		{
			Input: "96731408", // JM BONNISSEAU
			Expected: map[string]string{
				"SUDOC": "033522448",
				"ISNI":  "0000000068442789",
				"WKP":   "Q30084598",
				"LC":    "n2009050322",
				"BNF":   "12438130",
				"RERO":  "vtls008694187",
				"DNB":   "170346412",
			},
			ShouldFail: false,
		},
		{
			Input:      "",
			Expected:   nil,
			ShouldFail: true,
		},
	}
	for _, test := range viafGetIDsTests {
		actual, err := ViafGetIDs(test.Input)
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
