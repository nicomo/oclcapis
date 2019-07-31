package oclcapis

import (
	"encoding/xml"
	"errors"
	"fmt"
)

// WCIReadResult returns the input of a request to WCIBatchRead
// and the result in the form
// WCIdentity, with the error if any
type WCIReadResult struct {
	input  string
	output WCIdentity
	err    error
}

// WCIdentity stores the unmarshalled result
// from the WorldCat Identities web service
type WCIdentity struct {
	Pnkey    string   `xml:"pnkey"`
	AudLevel AudLevel `xml:"audLevel"`
	NameInfo NameInfo `xml:"nameInfo"`
}

// AudLevel is embedded in WCIdentity
// represents the "audience level"
// of the author's works
type AudLevel struct {
	Avg Avg `xml:"avg"`
}

// Avg is embedded in AudLevel
type Avg struct {
	Text  string `xml:",chardata"`
	Level string `xml:"level"`
}

// NameInfo is embedded in WCIdentity
type NameInfo struct {
	Type          string    `xml:"type,attr"`
	RawName       RawName   `xml:"rawName"`
	Languages     Languages `xml:"languages"`
	BirthDate     string    `xml:"birthDate"`
	Dates         Dates     `xml:"dates"`
	TotalHoldings string    `xml:"totalHoldings"`
	WorkCount     string    `xml:"workCount"`
	RecordCount   string    `xml:"recordCount"`
}

// RawName is embedded in NameInfo
type RawName struct {
	Suba string `xml:"suba"`
}

// Languages is embedded in NameInfo
type Languages struct {
	Count string `xml:"count,attr"`
	Lang  []Lang `xml:"lang"`
}

// Lang is embedded in Languages
type Lang struct {
	Code  string `xml:"code,attr"`
	Count string `xml:"count,attr"`
}

// Dates is embedded in NameInfo
type Dates struct {
	Date []Date `xml:"date"`
}

// Date is embedded in Dates
type Date struct {
	Text  string `xml:",chardata"`
	Count string `xml:"count,attr"`
}

const baseWCIdentitiesURL = "http://www.worldcat.org/identities/"

// WCIRead calls the WorldCat Identities web service
// for a given LC number (i.e. authority number)
func WCIRead(input string) (WCIdentity, error) {
	// will store the result
	var wci WCIdentity

	if input == "" {
		return wci, errors.New("input cannot be an empty string")
	}

	// e.g. http://www.worldcat.org/identities/lccn-n2009050322/
	getURL := baseWCIdentitiesURL + input + "/"

	// call WS & put the response into a []byte
	b, err := callWS(getURL)
	if err != nil {
		return wci, err
	}

	// unmarshall
	if err := xml.Unmarshal(b, &wci); err != nil {
		return wci, fmt.Errorf("could not unmarshall response: %v", err)
	}

	return wci, nil

}

// WCIBatchRead manages concurrent calls to WCIRead
func WCIBatchRead(input []string) ([]WCIReadResult, error) {
	if len(input) == 0 {
		return nil, errors.New("input cannot be an empty string")
	}

	// 2 channels used in a fan out / fan in pattern
	jobs := make(chan string, len(input))
	results := make(chan WCIReadResult, len(input))
	defer close(results)

	// dispatch jobs to number of workers, capping at 5
	numW := 5
	if len(input) < 5 {
		numW = len(input)
	}

	// This starts up to 5 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= numW; w++ {
		go wciBatchWorker(w, jobs, results)
	}

	// Here we send the jobs and then close the
	// channel to indicate that's all the work we have.
	for _, s := range input {
		jobs <- s
	}
	close(jobs)

	// fan in the results from the results channel
	var res []WCIReadResult
	for i := 1; i <= len(input); i++ {
		res = append(res, <-results)
	}
	if len(res) == 0 {
		return res, errors.New("could not find any result")
	}
	return res, nil
}

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`.
func wciBatchWorker(id int, jobs <-chan string, results chan<- WCIReadResult) {
	for j := range jobs {
		output, err := WCIRead(j)
		results <- WCIReadResult{
			input:  j,
			output: output,
			err:    err,
		}
	}
}
