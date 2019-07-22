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
				SourceIDs: []SourceID{
					SourceID{
						Src:   "SUDOC",
						SrcID: "033522448",
					},
					SourceID{
						Src:   "ISNI",
						SrcID: "0000000068442789",
					},
					SourceID{
						Src:   "WKP",
						SrcID: "Q30084598",
					},
					SourceID{
						Src:   "LC",
						SrcID: "n2009050322",
					},
					SourceID{
						Src:   "BNF",
						SrcID: "12438130",
					},
					SourceID{
						Src:   "RERO",
						SrcID: "vtls008694187",
					},
					SourceID{
						Src:   "DNB",
						SrcID: "170346412",
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
				SourceIDs: []SourceID{
					SourceID{
						Src:   "BNF",
						SrcID: "11885831",
					},
					SourceID{
						Src:   "DNB",
						SrcID: "119196506",
					},
					SourceID{
						Src:   "ISNI",
						SrcID: "0000000121457956",
					},
					SourceID{
						Src:   "LC",
						SrcID: "n50048876",
					},
					SourceID{
						Src:   "LNB",
						SrcID: "LNC10-000175758",
					},
					SourceID{
						Src:   "N6I",
						SrcID: "vtls000058144",
					},
					SourceID{
						Src:   "NKC",
						SrcID: "mub2014832218",
					},
					SourceID{
						Src:   "NLA",
						SrcID: "000035022438",
					},
					SourceID{
						Src:   "NTA",
						SrcID: "069900833",
					},
					SourceID{
						Src:   "NUKAT",
						SrcID: "n96001012",
					},
					SourceID{
						Src:   "SELIBR",
						SrcID: "179329",
					},
					SourceID{
						Src:   "SUDOC",
						SrcID: "026650649",
					},
					SourceID{
						Src:   "WKP",
						SrcID: "Q440528",
					},
					SourceID{
						Src:   "KRNLK",
						SrcID: "KAC200907825",
					},
					SourceID{
						Src:   "BIBSYS",
						SrcID: "90118757",
					},
					SourceID{
						Src:   "NLI",
						SrcID: "004009265",
					},
					SourceID{
						Src:   "LIH",
						SrcID: "LNB:V-374846;=BS",
					},
					SourceID{
						Src:   "RERO",
						SrcID: "vtls003046752",
					},
					SourceID{
						Src:   "PLWABN",
						SrcID: "9810680129805606",
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
		{
			Input: "213067771", // N MORIN
			Expected: ViafData{
				ViafID:   "213067771",
				NameType: "Personal",
				SourceIDs: []SourceID{
					SourceID{
						Src:   "SUDOC",
						SrcID: "075012286",
					},
				},
				WCIdentitiesURL: "https://www.worldcat.org/identities/viaf-213067771",
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
			t.Fatalf("FAIL for %s: expected %+v, actual result was %+v", test.Input, test.Expected, actual)
		}
	}
}
