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
                                        {
                                                Src:   "SUDOC",
                                                SrcID: "033522448",
                                        },
                                        {
                                                Src:   "ISNI",
                                                SrcID: "0000000068442789",
                                        },
                                        {
                                                Src:   "WKP",
                                                SrcID: "Q30084598",
                                        },
                                        {
                                                Src:   "LC",
                                                SrcID: "n2009050322",
                                        },
                                        {
                                                Src:   "BNF",
                                                SrcID: "12438130",
                                        },
                                        {
                                                Src:   "RERO",
                                                SrcID: "vtls008694187",
                                        },
                                        {
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
                                        {
                                                Src:   "BNF",
                                                SrcID: "11885831",
                                        },
                                        {
                                                Src:   "DNB",
                                                SrcID: "119196506",
                                        },
                                        {
                                                Src:   "ISNI",
                                                SrcID: "0000000121457956",
                                        },
                                        {
                                                Src:   "LC",
                                                SrcID: "n50048876",
                                        },
                                        {
                                                Src:   "LNB",
                                                SrcID: "LNC10-000175758",
                                        },
                                        {
                                                Src:   "N6I",
                                                SrcID: "vtls000058144",
                                        },
                                        {
                                                Src:   "NKC",
                                                SrcID: "mub2014832218",
                                        },
                                        {
                                                Src:   "NLA",
                                                SrcID: "000035022438",
                                        },
                                        {
                                                Src:   "NTA",
                                                SrcID: "069900833",
                                        },
                                        {
                                                Src:   "NUKAT",
                                                SrcID: "n96001012",
                                        },
                                        {
                                                Src:   "SELIBR",
                                                SrcID: "179329",
                                        },
                                        {
                                                Src:   "SUDOC",
                                                SrcID: "026650649",
                                        },
                                        {
                                                Src:   "WKP",
                                                SrcID: "Q440528",
                                        },
                                        {
                                                Src:   "KRNLK",
                                                SrcID: "KAC200907825",
                                        },
                                        {
                                                Src:   "BIBSYS",
                                                SrcID: "90118757",
                                        },
                                        {
                                                Src:   "NLI",
                                                SrcID: "004009265",
                                        },
                                        {
                                                Src:   "LIH",
                                                SrcID: "LNB:V-374846;=BS",
                                        },
                                        {
                                                Src:   "RERO",
                                                SrcID: "vtls003046752",
                                        },
                                        {
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
                                        {
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