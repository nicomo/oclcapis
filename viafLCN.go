package oclcapis

import (
	"errors"
	"fmt"
)

type batchResult struct {
	input  string
	output string
	err    error
}

// ViafGetLCN finds a Library of Congress ID
// from a VIAF ID
func ViafGetLCN(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be an empty string")
	}
	data, err := viafGetData(input)
	if err != nil {
		return "", fmt.Errorf("could not get a valid response for %s: %v", input, err)
	}

	for _, src := range data.SourceIDs {
		if src.Src != "LC" {
			continue
		}
		return src.SrcID, nil
	}
	return "", fmt.Errorf("could not find a LC Number for %s", input)
}

// ViafGetLCNs finds Library of Congress IDs
// from a slice of VIAF IDs in batches
func ViafGetLCNs(input []string) (map[string]string, error) {
	m, err := batchID("lcn", input)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// batchID takes VIAF IDs as input
// + a service (either LCN or WKP)
// and retrieves the corresponding IDs (LCN or WKP)
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
		default:
			br.err = fmt.Errorf("unkown service %s", j[0])
		}

		results <- br
	}
}
