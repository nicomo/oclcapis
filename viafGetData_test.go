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
				//XLinks:          XLinks{},
				WCIdentitiesURL: "https://www.worldcat.org/identities/lccn-n2009050322",
			},
			ShouldFail: false,
		},
		{
			Input: "101833644", // C. BROOKE-ROSE
			Expected: ViafData{
				ViafID:   "101833644",
				NameType: "Personal",
				Sources: Sources{
					Source: []Source{
						{
							Nsid: "http://catalogue.bnf.fr/ark:/12148/cb118858311",
							Text: "BNF|11885831",
						},
						{
							Nsid: "http://d-nb.info/gnd/119196506",
							Text: "DNB|119196506",
						},
						{
							Nsid: "0000000121457956",
							Text: "ISNI|0000000121457956",
						},
						{
							Nsid: "n50048876",
							Text: "LC|n  50048876",
						},
						{
							Nsid: "LNC10-000175758",
							Text: "LNB|LNC10-000175758",
						},
						{
							Nsid: "vtls000058144",
							Text: "N6I|vtls000058144",
						},
						{
							Nsid: "mub2014832218",
							Text: "NKC|mub2014832218",
						},
						{
							Nsid: "000035022438",
							Text: "NLA|000035022438",
						},
						{
							Nsid: "069900833",
							Text: "NTA|069900833",
						},
						{
							Nsid: "vtls000153113",
							Text: "NUKAT|n  96001012",
						},
						{
							Nsid: "31fhh8vm4gcckf3",
							Text: "SELIBR|179329",
						},
						{
							Nsid: "026650649",
							Text: "SUDOC|026650649",
						},
						{
							Nsid: "Q440528",
							Text: "WKP|Q440528",
						},
						{
							Nsid: "KAC200907825",
							Text: "KRNLK|KAC200907825",
						},
						{
							Nsid: "90118757",
							Text: "BIBSYS|90118757",
						},
						{
							Nsid: "004009265",
							Text: "NLI|004009265",
						},
						{
							Nsid: "LNB:V*374846;=BS",
							Text: "LIH|LNB:V-374846;=BS",
						},
						{
							Nsid: "vtls003046752",
							Text: "RERO|vtls003046752",
						},
						{
							Nsid: "9810680129805606",
							Text: "PLWABN|9810680129805606",
						},
					},
				},
				WCIdentitiesURL: "https://www.worldcat.org/identities/lccn-n50048876",
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
			t.Fatalf("FAIL for %s: expected %v, actual result was %+v", test.Input, test.Expected, actual)
		}
	}
}
