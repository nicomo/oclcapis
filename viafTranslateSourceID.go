package oclcapis

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// ViafTranslateSourceID takes an ID from an external source
// e.g. DNB, Sudoc, etc
// and retrieves the corresponding VIAF ID
// the input should be an url encoded string, e.g. SUDOC%7c033522448
func ViafTranslateSourceID(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be an empty string")
	}

	// e.g. http://www.viaf.org/viaf/sourceID/SUDOC%7c033522448/
	getURL := baseViafURL + "SourceID/" + input

	// we need a client because we want to control for HTTP Status 301
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(getURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusMovedPermanently {
		return "", fmt.Errorf("Could not translate %s: %s", input, resp.Status)
	}

	loc, err := resp.Location()
	if err != nil {
		return "", fmt.Errorf("location URL not found%v", err)
	}

	return strings.TrimPrefix(loc.Path, "/viaf/"), nil

}
