package oclcapis

import (
	"reflect"
	"testing"
)

type viafTranslateSourceIDTest struct {
	Input      string
	Expected   string
	ShouldFail bool
}

func TestViafTranslateSourceID(t *testing.T) {
	var viafTranslateTests = []viafTranslateSourceIDTest{
		{
			Input:      "",
			Expected:   "",
			ShouldFail: true,
		},
		{
			Input:      "SUDOC%7c033522448",
			Expected:   "96731408",
			ShouldFail: false,
		},
		{
			Input:      "WRONGSTRING", // does not exist
			Expected:   "",
			ShouldFail: true,
		},
	}

	for _, test := range viafTranslateTests {
		actual, err := ViafTranslateSourceID(test.Input)
		if err != nil {
			if test.ShouldFail {
				t.Logf("PASS: got expected error %v", err)
			} else {
				t.Fatalf("FAIL for %s: expected %v, got an error %v", test.Input, test.Expected, err)
			}
			continue
		}
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %+v", test.Input, test.Expected, actual)
		}
	}

}
