package oclcapis

import (
	"errors"
	"fmt"
)

// ViafGetWKP finds a Wikipedia ID
// from a VIAF ID
func ViafGetWKP(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be an empty string")
	}
	data, err := viafGetData(input)
	if err != nil {
		return "", fmt.Errorf("could not get a valid response for %s: %v", input, err)
	}

	for _, source := range data.Sources.Source {
		s := viafSplitSourceID(source.Text)
		if s == "" || s != "WKP" {
			continue
		}
		return source.Nsid, nil
	}
	return "", fmt.Errorf("could not find a Wikipedia Number for %s", input)
}
