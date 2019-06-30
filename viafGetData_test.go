package oclcapis

import (
	"reflect"
	"testing"
)

type viafGetDataTest struct {
	Input      string
	Expected   ViafData
	ShouldFail bool
}

type viafGetLCNTest struct {
	Input      string
	Expected   string
	ShouldFail bool
}

type viafGetWKPTest struct {
	Input      string
	Expected   string
	ShouldFail bool
}

type viafGetIDsTest struct {
	Input      string
	Expected   map[string]string
	ShouldFail bool
}

func TestViafGetData(t *testing.T) {

	var viafDataTests = []viafGetDataTest{
		{
			Input: "96731408", // JM BONNISSEAU
			Expected: ViafData{
				ViafID:   "96731408",
				NameType: "Personal",
				Sources: Sources{
					Source: []Source{
						{
							Nsid: "033522448",
							Text: "SUDOC|033522448",
						},
						{
							Nsid: "0000000068442789",
							Text: "ISNI|0000000068442789",
						},
						{
							Nsid: "Q30084598",
							Text: "WKP|Q30084598",
						},
						{
							Nsid: "n2009050322",
							Text: "LC|n 2009050322",
						},
						{
							Nsid: "http://catalogue.bnf.fr/ark:/12148/cb12438130w",
							Text: "BNF|12438130",
						},
						{
							Nsid: "vtls008694187",
							Text: "RERO|vtls008694187",
						},
						{
							Nsid: "http://d-nb.info/gnd/170346412",
							Text: "DNB|170346412",
						},
					},
				},
				XLinks: XLinks{
					XLink: "https://www.worldcat.org/identities/lccn-n2009050322",
				},
			},
			ShouldFail: false,
		},
	}

	for _, test := range viafDataTests {
		actual, err := viafGetData(test.Input)
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

func TestViafGetLCN(t *testing.T) {
	var viafLCNTests = []viafGetLCNTest{
		{
			Input:      "96731408", // JM BONNISSEAU
			Expected:   "n2009050322",
			ShouldFail: false,
		},
	}
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

func TestViafGetWKP(t *testing.T) {
	var viafWKPTests = []viafGetWKPTest{
		{
			Input:      "96731408", // JM BONNISSEAU
			Expected:   "Q30084598",
			ShouldFail: false,
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

func TestViafGetIDs(t *testing.T) {
	var viafGetIDsTests = []viafGetIDsTest{
		{
			Input: "96731408", // JM BONNISSEAU
			Expected: map[string]string{
				"SUDOC": "033522448",
				"ISNI":  "0000000068442789",
				"WKP":   "Q30084598",
				"LC":    "n2009050322",
				"BNF":   "http://catalogue.bnf.fr/ark:/12148/cb12438130w",
				"RERO":  "vtls008694187",
				"DNB":   "http://d-nb.info/gnd/170346412",
			},
			ShouldFail: false,
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
