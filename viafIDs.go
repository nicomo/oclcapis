package oclcapis

import (
	"errors"
	"fmt"
)

// AllIDsResult returns the input of a request to ViafGetAllIDs
// (e.g. 123456) and the result in the form
// map[DNB]{"TYU6756"}, with the error if any
type AllIDsResult struct {
	input  string
	output map[string]string
	err    error
}

// ViafGetIDs finds all source IDs, e.g. LC, DNB, WKP, etc.
// from a VIAF ID
func ViafGetIDs(input string) (map[string]string, error) {
	if input == "" {
		return nil, errors.New("input cannot be an empty string")
	}
	data, err := viafGetData(input)
	if err != nil {
		return nil, fmt.Errorf("could not get a valid response for %s: %v", input, err)
	}

	result := map[string]string{}

	for _, source := range data.SourceIDs {
		result[source.Src] = source.SrcID
	}
	return result, nil
}

// ViafGetAllIDs finds all source IDs, e.g. LC, DNB, WKP, etc.
// from a bucn of VIAF ID
func ViafGetAllIDs(input []string) ([]AllIDsResult, error) {
	if len(input) == 0 {
		return nil, errors.New("input cannot be an empty string")
	}

	// 2 channels used in a fan out / fan in pattern
	jobs := make(chan string, len(input))
	results := make(chan AllIDsResult, len(input))
	defer close(results)

	// dispatch jobs to number of workers, capping at 5
	numW := 5
	if len(input) < 5 {
		numW = len(input)
	}

	// This starts up to 5 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= numW; w++ {
		go allIDsWorker(w, jobs, results)
	}

	// Here we send the jobs and then close the
	// channel to indicate that's all the work we have.
	for _, s := range input {
		jobs <- s
	}
	close(jobs)

	// fan in the results from the results channel
	var res []AllIDsResult
	for i := 1; i <= numW; i++ {
		res = append(res, <-results)
	}

	return res, nil
}

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`.
func allIDsWorker(id int, jobs <-chan string, results chan<- AllIDsResult) {
	for j := range jobs {
		output, err := ViafGetIDs(j)
		results <- AllIDsResult{
			input:  j,
			output: output,
			err:    err,
		}
	}
}
