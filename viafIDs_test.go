package oclcapis

import (
	"errors"
	"reflect"
	"testing"
)

type viafGetIDsTest struct {
	Description string
	Input       string
	Expected    map[string]string
	ShouldFail  bool
}

type ViafGetAllIDsTest struct {
	Description string
	Input       []string
	Expected    []AllIDsResult
	ShouldFail  bool
}

var viafGetIDsTests = []viafGetIDsTest{
	{
		Description: "JM BONNISSEAU",
		Input:       "96731408",
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
		Description: "empty input",
		Input:       "",
		Expected:    nil,
		ShouldFail:  true,
	},
}

func TestViafGetIDs(t *testing.T) {
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

func TestViafGetIDsBatch(t *testing.T) {
	test := ViafGetAllIDsTest{
		Description: "Concurrent fetching of source IDs",
		Input: []string{
			"96731408",
			"",
		},
		Expected: []AllIDsResult{
			{
				input: "96731408",
				output: map[string]string{
					"SUDOC": "033522448",
					"ISNI":  "0000000068442789",
					"WKP":   "Q30084598",
					"LC":    "n2009050322",
					"BNF":   "12438130",
					"RERO":  "vtls008694187",
					"DNB":   "170346412",
				},
				err: nil,
			},
			{
				input: "",
				err:   errors.New("input cannot be an empty string"),
			},
		},
		ShouldFail: false,
	}

	actual, err := ViafGetAllIDs(test.Input)
	// general error
	if err != nil {
		if test.ShouldFail {
			t.Logf("PASS %s: got expected error %v", test.Description, err)
		} else {
			t.Fatalf("FAIL %s (input was %s): expected %v, got error %v", test.Description, test.Input, test.Expected, err)
		}
	}

	// checks for each specific result
	for _, a := range actual {

		// check for specific errors
		if a.err != nil {
			for _, e := range test.Expected {
				if a.input == e.input {
					if e.err != nil {
						t.Logf("PASS %s: got expected error %v", a.input, a.err)
						continue
					} else {
						t.Fatalf("FAIL %s: expected %v, got error %v", e.input, e.output, a.err)
					}
				}
			}
		}

		// checks for results
		for _, e := range test.Expected {
			if a.input == e.input {
				if reflect.DeepEqual(e.output, a.output) {
					t.Logf("PASS %s", e.input)
				} else {
					t.Fatalf("FAIL %s: expected %v, actual result was %v", e.input, e.output, a.output)
				}
			}
		}

	}

}
