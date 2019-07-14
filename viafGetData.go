package oclcapis

import (
	"encoding/json"
)

const baseViafURL = "http://www.viaf.org/viaf/"

// ViafData is used to unmarshal
// the response coming from the
// VIAF GetDataInFormat web service
type ViafData struct {
	ViafID   string  `json:"viafID"`
	NameType string  `json:"nameType"`
	Sources  Sources `json:"sources"`
	XLinks   XLinks  `json:"xLinks"`
}

// Source is embedded in ViafData
type Source struct {
	Nsid string `json:"@nsid"`
	Text string `json:"#text"`
}

// Sources is embedded in ViafData
type Sources struct {
	Source []Source `json:"source"`
}

// XLinks is embedded in ViafData
type XLinks struct {
	XLink string `json:"xLink"`
}

// viafGetData hits the OCLC VIAF API
// with a VIAF ID and retrieves
// extra information from the record
func viafGetData(input string) (ViafData, error) {

	// will store result
	var vData ViafData

	// http://www.viaf.org/viaf/96731408/viaf.json
	getURL := baseViafURL + input + "/viaf.json"

	// call WS & put the response into a []byte
	b, err := callWS(getURL)
	if err != nil {
		return vData, err
	}

	// unmarshall
	if err := json.Unmarshal(b, &vData); err != nil {
		return vData, err
	}

	return vData, nil
}
