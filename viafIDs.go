package oclcapis

import (
	"errors"
	"fmt"
)

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

	for _, source := range data.Sources.Source {
		s := viafSplitSourceID(source.Text)
		if s == "" {
			continue
		}
		result[s] = source.Nsid
	}
	return result, nil
}
