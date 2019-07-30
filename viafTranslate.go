package oclcapis

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// ViafTranslate takes an ID from an external source
// e.g. DNB, Sudoc, etc
// and retrieves the corresponding VIAF ID
// the input should be an url encoded string, e.g. SUDOC%7c033522448
func ViafTranslate(input string) (string, error) {
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
		return "", fmt.Errorf("could not translate %s: %s", input, resp.Status)
	}

	loc, err := resp.Location()
	if err != nil {
		return "", fmt.Errorf("location URL not found: %v", err)
	}

	return strings.TrimPrefix(loc.Path, "/viaf/"), nil

}

// ViafTranslateBatch takes a string of IDs from external sources
// e.g. DNB, Sudoc, etc
// and retrieves the corresponding VIAF IDs
// the input should be an url encoded string, e.g. SUDOC%7c033522448
func ViafTranslateBatch(input []string) (map[string]string, error) {
	m, err := batchID("translate", input)
	if err != nil {
		return nil, err
	}
	return m, nil
}
