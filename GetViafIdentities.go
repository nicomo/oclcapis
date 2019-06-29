package oclcapis

import (
	"encoding/json"
	"fmt"
	"log"
)

const baseViafURL = "http://www.viaf.org/viaf/"

// GetViafIdentities hits the OCLC VIAF API
// with a VIAF ID and retrieves
// lccn and wikipedia IDs from the record
func GetViafIdentities(input string) ViafIdentity {
	// http://www.viaf.org/viaf/96731408/viaf.json
	getURL := baseViafURL + input + "/viaf.json"

	// call WS & put the response into a []byte
	b, err := callWS(getURL)
	if err != nil {
		log.Fatalln(err)
	}

	// unmarshall
	var resp ViafIdentity
	if err := json.Unmarshal(b, &resp); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	return resp
}
