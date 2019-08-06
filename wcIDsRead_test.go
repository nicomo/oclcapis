package oclcapis

import (
	"reflect"
	"testing"
)

type wciTest struct {
	Description string
	Input       string
	Expected    WCIdentity
	ShouldFail  bool
}

type wciBatchTest struct {
	Description string
	Input       []string
	Expected    []WCIReadResult
	ShouldFail  bool
}

var wciTests = []wciTest{
	{
		Description: "Valid, JM BONNISSEAU",
		Input:       "lccn-n2009050322",
		Expected: WCIdentity{
			Pnkey: "lccn-n2009050322",
			AudLevel: AudLevel{
				Avg: Avg{
					Text:  "",
					Level: "0.93",
				},
			},
			NameInfo: NameInfo{
				Type: "personal",
				RawName: RawName{
					Suba: "Bonnisseau, Jean-Marc",
				},
				Languages: Languages{
					Count: "2",
					Lang: []Lang{
						{
							Code:  "eng",
							Count: "25",
						},
						{
							Code:  "fre",
							Count: "7",
						},
					},
				},
				BirthDate: "1957",
				Dates: Dates{
					Date: []Date{
						{
							Text:  "1985",
							Count: "1",
						},
						{
							Text:  "1986",
							Count: "6",
						},
						{
							Text:  "1987",
							Count: "5",
						},
						{
							Text:  "1988",
							Count: "9",
						},
						{
							Text:  "1990",
							Count: "1",
						},
						{
							Text:  "1992",
							Count: "1",
						},
						{
							Text:  "1994",
							Count: "1",
						},
						{
							Text:  "1995",
							Count: "1",
						},
						{
							Text:  "1996",
							Count: "1",
						},
						{
							Text:  "1997",
							Count: "4",
						},
						{
							Text:  "1998",
							Count: "2",
						},
						{
							Text:  "2000",
							Count: "4",
						},
						{
							Text:  "2001",
							Count: "1",
						},
						{
							Text:  "2005",
							Count: "1",
						},
						{
							Text:  "2007",
							Count: "3",
						},
						{
							Text:  "2008",
							Count: "1",
						},
						{
							Text:  "2012",
							Count: "1",
						},
						{
							Text:  "2013",
							Count: "4",
						},
						{
							Text:  "2015",
							Count: "1",
						},
						{
							Text:  "2016",
							Count: "1",
						},
						{
							Text:  "2018",
							Count: "1",
						},
					},
				},
				TotalHoldings: "77",
				WorkCount:     "42",
				RecordCount:   "54",
			},
		},
		ShouldFail: false,
	},
	{
		Description: "Returns 404, resource does not exist",
		Input:       "lccn-no201312602",
		Expected:    WCIdentity{},
		ShouldFail:  true,
	},
	{
		Description: "Invalid input",
		Input:       "",
		Expected:    WCIdentity{},
		ShouldFail:  true,
	},
}

func TestWCIRead(t *testing.T) {
	for _, test := range wciTests {
		actual, err := WCIRead(test.Input)
		if err != nil {
			if test.ShouldFail {
				t.Logf("PASS: got expected error %v", err)
			} else {
				t.Fatalf("FAIL for %s: expected %v, got an error %v", test.Input, test.Expected, err)
			}
		}

		// not testing on the whole struct. Can change at the server end.
		if reflect.DeepEqual(test.Expected.NameInfo.RawName.Suba, actual.NameInfo.RawName.Suba) && reflect.DeepEqual(test.Expected.Pnkey, actual.Pnkey) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %+v", test.Input, test.Expected.NameInfo, actual.NameInfo)
		}
	}
}

func TestWCIBatchRead(t *testing.T) {
	wciBatchTests := []wciBatchTest{}
	wbt := wciBatchTest{
		Description: "batch request of WorldCat Identities, happy path",
	}
	for _, w := range wciTests {
		wbt.Input = append(wbt.Input, w.Input)
		wbt.Expected = append(wbt.Expected,
			WCIReadResult{
				Input:  w.Input,
				Output: w.Expected,
			},
		)
	}
	wciBatchTests = append(wciBatchTests, wbt)

	for _, test := range wciBatchTests {
		actual, err := WCIBatchRead(test.Input)
		if err != nil {
			if test.ShouldFail {
				t.Logf("PASS: got expected error %v", err)
			} else {
				t.Fatalf("FAIL for %s: expected %v, got an error %v", test.Input, test.Expected, err)
			}
		}

		for _, a := range actual {
			for _, e := range test.Expected {
				if a.Input == e.Input {
					if reflect.DeepEqual(e.Output.NameInfo.RawName.Suba, a.Output.NameInfo.RawName.Suba) && reflect.DeepEqual(e.Output.Pnkey, a.Output.Pnkey) {
						t.Logf("PASS: got %v", a.Output)
					} else {
						t.Fatalf("FAIL for %s: expected %v, actual result was %+v", e.Input, e.Output, a.Output)
					}
				}
			}
		}
	}
}
