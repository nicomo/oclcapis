package oclcapis

import (
	"reflect"
	"testing"
)

type viafTranslateTest struct {
	Description string
	Input       string
	Expected    string
	ShouldFail  bool
}

type viafTranslateBatchTest struct {
	Description string
	Input       []string
	Expected    map[string]string
	ShouldFail  bool
}

func TestViafTranslate(t *testing.T) {
	var viafTranslateTests = []viafTranslateTest{
		{
			Description: "empty input",
			Input:       "",
			Expected:    "",
			ShouldFail:  true,
		},
		{
			Description: "JM BONNISSEAU",
			Input:       "SUDOC%7c033522448",
			Expected:    "96731408",
			ShouldFail:  false,
		},
		{
			Description: "No result",
			Input:       "WRONGSTRING", // does not exist
			Expected:    "",
			ShouldFail:  true,
		},
	}

	for _, test := range viafTranslateTests {
		actual, err := ViafTranslate(test.Input)
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

func TestViafTranslateBatch(t *testing.T) {

	test := viafTranslateBatchTest{
		Description: "Concurrent translating of IDs",
		Input: []string{
			"SUDOC%7c033522448",
			"NoResultString",
		},
		Expected:   make(map[string]string),
		ShouldFail: false,
	}
	test.Expected["SUDOC%7c033522448"] = "96731408"
	test.Expected["NoResultString"] = "could not translate NoResultString: 404 Not Found\n"

	actual, err := ViafTranslateBatch(test.Input)
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
