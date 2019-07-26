package oclcapis

import (
	"errors"
	"fmt"
)

type lcnResult struct {
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
	fmt.Printf("data SourceIDs: %+v\n", data.SourceIDs)
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
// TODO: write the result to a file
func ViafGetLCNs(input []string) (map[string]string, error) {

	// 2 channels used in a fan out / fan in pattern
	jobs := make(chan string, len(input))
	results := make(chan lcnResult, len(input))

	// dispatch jobs to number of workers, capping at 5
	numW := 5
	if len(input) < 5 {
		numW = len(input)
	}

	// This starts up to 5 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= numW; w++ {
		go lcnWorker(w, jobs, results)
	}

	// Here we send the jobs and then close the
	// channel to indicate that's all the work we have.
	for _, s := range input {
		jobs <- s
	}
	close(jobs)

	// fan in the results from the results channel
	var res []lcnResult
	for a := 1; a <= numW; a++ {
		res = append(res, <-results)
	}

	fmt.Println(res)

	m := make(map[string]string)
	for _, r := range res {
		if r.err != nil {
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
func lcnWorker(id int, jobs <-chan string, results chan<- lcnResult) {
	for s := range jobs {
		fmt.Printf("worker %d, job: %s\n", id, s)
		output, err := ViafGetLCN(s)
		fmt.Printf("worker %d, output: %s (err %v)\n", id, output, err)
		if err != nil {
			results <- lcnResult{
				input:  s,
				output: "",
				err:    err,
			}
		}

		results <- lcnResult{
			input:  s,
			output: output,
			err:    nil,
		}
	}
}
