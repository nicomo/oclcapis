package oclcapis

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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
