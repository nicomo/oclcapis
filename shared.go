package oclcapis

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type batchResult struct {
	input  string
	output string
	err    error
}

// call performs the http GET
// retrieves the response and puts it in a slice of bytes
func callWS(getURL string) ([]byte, error) {
	// get the result from the url
	resp, err := http.Get(getURL)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status not OK: %d %s", resp.StatusCode, resp.Status)
	}

	// put the response into a []byte
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// viafSplitSourceID separates the DNC or SUDOC, etc
// from the actual value
func viafSplitSourceID(sText string) (SourceID, error) {
	var srcID SourceID
	t := strings.Split(sText, "|")
	if len(t) != 2 {
		return srcID, fmt.Errorf("could not properly split source id %s", sText)
	}
	srcID.Src = t[0]
	// remove whitespaces in ID
	srcID.SrcID = strings.Join(strings.Fields(t[1]), "")
	return srcID, nil
}

// batchID takes IDs as input
// + a service (lcn, wkp or translate)
// and retrieves the corresponding IDs (LCN or WKP or viaf ID)
func batchID(service string, input []string) (map[string]string, error) {

	// 2 channels used in a fan out / fan in pattern
	jobs := make(chan []string, len(input))
	results := make(chan batchResult, len(input))
	defer close(results)

	// dispatch jobs to number of workers, capping at 5
	numW := 5
	if len(input) < 5 {
		numW = len(input)
	}

	// This starts up to 5 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= numW; w++ {
		go batchIDWorker(w, jobs, results)
	}

	// Here we send the jobs and then close the
	// channel to indicate that's all the work we have.
	for _, s := range input {
		j := []string{service, s}
		jobs <- j
	}
	close(jobs)

	// fan in the results from the results channel
	var res []batchResult
	for i := 1; i <= numW; i++ {
		res = append(res, <-results)
	}

	m := make(map[string]string)
	for _, r := range res {
		if r.err != nil {
			if len(r.input) == 0 {
				continue
			}
			m[r.input] = fmt.Sprintln(r.err)
			continue
		}
		m[r.input] = r.output
	}
	return m, nil
}

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`.
func batchIDWorker(id int, jobs <-chan []string, results chan<- batchResult) {
	for j := range jobs {
		var br batchResult
		switch j[0] {
		case "lcn":
			output, err := ViafGetLCN(j[1])
			br.input = j[1]
			br.output = output
			br.err = err
		case "wkp":
			output, err := ViafGetWKP(j[1])
			br.input = j[1]
			br.output = output
			br.err = err
		case "translate":
			output, err := ViafTranslate(j[1])
			br.input = j[1]
			br.output = output
			br.err = err
		default:
			br.err = fmt.Errorf("unkown service %s", j[0])
		}

		results <- br
	}
}
