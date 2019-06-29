package oclcapis

import (
	"reflect"
	"testing"
)

type viafIdentitiesTest struct {
	Input    string
	Expected ViafIdentity
}

func TestGetViafIdentities(t *testing.T) {

	var ViafIdentitiesTests = []viafIdentitiesTest{
		{
			Input: "96731408", // JM BONNISSEAU
			Expected: ViafIdentity{
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
		},
	}

	for _, test := range ViafIdentitiesTests {
		actual := GetViafIdentities(test.Input)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Input, test.Expected, actual)
		}
	}
}
