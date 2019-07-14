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
		{
			Input:      "",
			Expected:   ViafData{},
			ShouldFail: true,
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
