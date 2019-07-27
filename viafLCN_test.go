package oclcapis

import (
	"reflect"
	"testing"
)

type viafGetLCNTest struct {
	Description string
	Input       string
	Expected    string
	ShouldFail  bool
}

type viafGetLCNsTest struct {
	Description string
	Input       []string
	Expected    map[string]string
	ShouldFail  bool
}

var viafLCNTests = []viafGetLCNTest{
	{
		Description: "happy path, JM BONNISSEAU",
		Input:       "96731408",
		Expected:    "n2009050322",
		ShouldFail:  false,
	},
	{
		Description: "fail, no result",
		Input:       "213067771",
		Expected:    "",
		ShouldFail:  true,
	},
	{
		Description: "happy path, C. BROOKE-ROSE",
		Input:       "101833644",
		Expected:    "n50048876",
		ShouldFail:  false,
	},
	{
		Description: "fail, empty input",
		Input:       "",
		Expected:    "",
		ShouldFail:  true,
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

func TestViafGetLCNs(t *testing.T) {

	test := viafGetLCNsTest{
		Description: "Concurrent fetching of LCNs",
		Input: []string{
			"96731408",
			"213067771",
			"101833644",
		},
		Expected:   make(map[string]string),
		ShouldFail: false,
	}
	test.Expected["96731408"] = "n2009050322"
	test.Expected["213067771"] = "could not find a LC Number for 213067771\n"
	test.Expected["101833644"] = "n50048876"

	actual, err := ViafGetLCNs(test.Input)
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
