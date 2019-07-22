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

	for _, src := range data.SourceIDs {
		if src.Src != "WKP" {
			continue
		}
		return src.SrcID, nil
	}

	return "", fmt.Errorf("could not find a Wikipedia Number for %s", input)
}
