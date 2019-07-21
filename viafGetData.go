package oclcapis

import (
	"encoding/json"
	"fmt"
)

const baseViafURL = "http://www.viaf.org/viaf/"

// ViafData is used to unmarshal
// the response coming from the
// VIAF GetDataInFormat web service
type ViafData struct {
	ViafID          string  `json:"viafID"`
	NameType        string  `json:"nameType"`
	Sources         Sources `json:"sources"`
	XLinks          XLinks  `json:"xLinks"`
	WCIdentitiesURL string  // WCIdentitiesURL is the value extracted from XLinks above
}

// Sources is embedded in ViafData
type Sources struct {
	Source []Source `json:"source"`
}

// Source is embedded in ViafData
type Source struct {
	Nsid string `json:"@nsid"`
	Text string `json:"#text"`
}

// XLinks is embedded in ViafData
// the shape of the returned JSON changes
// according to the content :-(
// we have to unmarshal to an empty interface
type XLinks struct {
	XLink interface{} `json:"xLink"`
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

	// inspect XLinks to extract WCIdentitiesURL
	switch t := vData.XLinks.XLink.(type) {
	case string:
		vData.WCIdentitiesURL = t
		vData.XLinks = XLinks{}
	case []interface{}:
		for _, vt := range t {
			switch vtt := vt.(type) {
			case string:
				vData.WCIdentitiesURL = vtt
				vData.XLinks = XLinks{}
			default:
				fmt.Println("skip over XLink non string")
			}
		}

	default:
		fmt.Printf("Xlinks data type not a string: %+v\n", vData.XLinks.XLink)
	}

	return vData, nil
}
