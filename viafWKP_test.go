package oclcapis

import (
	"reflect"
	"testing"
)

type viafGetWKPTest struct {
	Description string
	Input       string
	Expected    string
	ShouldFail  bool
}

type viafGetWKPsTest struct {
	Description string
	Input       []string
	Expected    map[string]string
	ShouldFail  bool
}

func TestViafGetWKP(t *testing.T) {

	viafWKPTests := []viafGetWKPTest{
		{
			Description: "JM BONNISSEAU",
			Input:       "96731408",
			Expected:    "Q30084598",
			ShouldFail:  false,
		},
		{
			Description: "empty input",
			Input:       "",
			Expected:    "",
			ShouldFail:  true,
		},
		{
			Description: "no result",
			Input:       "213067771",
			Expected:    "",
			ShouldFail:  true,
		},
		{
			Description: "C. BROOKE-ROSE",
			Input:       "101833644",
			Expected:    "Q440528",
			ShouldFail:  false,
		},
	}
	for _, test := range viafWKPTests {
		actual, err := ViafGetWKP(test.Input)
		if err != nil {
			if test.ShouldFail {
				t.Logf("PASS %s: got expected error %v", test.Description, err)
			} else {
				t.Fatalf("FAIL %s (input was %s): expected %v, got error %v", test.Description, test.Input, test.Expected, err)
			}
		}
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS %s", test.Description)
		} else {
			t.Fatalf("FAIL %s (input was %s): expected %v, actual result was %v", test.Description, test.Input, test.Expected, actual)
		}
	}
}

func TestViafGetWKPs(t *testing.T) {

	test := viafGetWKPsTest{
		Description: "Concurrent fetching of WKPs",
		Input: []string{
			"96731408",
		},
		Expected:   make(map[string]string),
		ShouldFail: false,
	}
	test.Expected["96731408"] = "Q30084598"

	actual, err := ViafGetWKPs(test.Input)
	if err != nil {
		if test.ShouldFail {
			t.Logf("PASS %s: got expected error %v", test.Description, err)
		} else {
			t.Fatalf("FAIL %s (input was %s): expected %v, got error %v", test.Description, test.Input, test.Expected, err)
		}
	}
	if reflect.DeepEqual(test.Expected, actual) {
		t.Logf("PASS %s", test.Description)
	} else {
		t.Fatalf("FAIL %s (input was %s): expected %v, actual result was %v", test.Description, test.Input, test.Expected, actual)
	}
}
