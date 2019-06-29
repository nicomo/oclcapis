package oclcapis

import (
	"io/ioutil"
	"net/http"
)

// call performs the http GET
// retrieves the response and puts it in a slice of bytes
func callWS(getURL string) ([]byte, error) {
	// get the result from the url
	resp, err := http.Get(getURL)
	if err != nil {
		return []byte{}, err
	}

	// put the response into a []byte
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
