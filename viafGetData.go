package oclcapis

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

const baseViafURL = "http://www.viaf.org/viaf/"

// viafGetData hits the OCLC VIAF API
// with a VIAF ID and retrieves
// extra information from the record
func viafGetData(input string) (ViafData, error) {
	// http://www.viaf.org/viaf/96731408/viaf.json
	getURL := baseViafURL + input + "/viaf.json"

	// call WS & put the response into a []byte
	b, err := callWS(getURL)
	if err != nil {
		log.Fatalln(err)
	}

	// unmarshall
	var resp ViafData
	if err := json.Unmarshal(b, &resp); err != nil {
		return resp, err
	}

	return resp, nil
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

	for _, source := range data.Sources.Source {
		s := viafIDsGetSource(source.Text)
		if s == "" || s != "LC" {
			continue
		}
		return source.Nsid, nil
	}
	return "", fmt.Errorf("could not find a LC Number for %s", input)
}

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
		s := viafIDsGetSource(source.Text)
		if s == "" || s != "WKP" {
			continue
		}
		return source.Nsid, nil
	}
	return "", fmt.Errorf("could not find a Wikipedia Number for %s", input)
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

	for _, source := range data.Sources.Source {
		s := viafIDsGetSource(source.Text)
		if s == "" {
			continue
		}
		result[s] = source.Nsid
	}
	return result, nil
}

func viafIDsGetSource(sText string) string {
	t := strings.Split(sText, "|")
	if len(t) != 2 {
		return ""
	}
	return t[0]
}
