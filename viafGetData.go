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
	ViafID          string    `json:"viafID"`
	NameType        string    `json:"nameType"`
	Sources         Sources   `json:"sources"`
	SourceIDs       SourceIDs // SourceIDs is the value extracted from Sources above
	XLinks          XLinks    `json:"xLinks"`
	WCIdentitiesURL string    // WCIdentitiesURL is the value extracted from XLinks above
}

// Sources is used to unmarshal
// data about other sources ID
// into an interface (JSON structure subject to
// change on the OCLC end)
type Sources struct {
	Source interface{} `json:"source"`
}

// SourceIDs stores multiple IDs
type SourceIDs []SourceID

// SourceID store a single
// source ID, e.g. SUDOC / 1234567
type SourceID struct {
	Src   string `json:"src"`
	SrcID string `json:"src_id"`
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

	// inspect Sources to extract source IDs
	vData.SourceIDs = []SourceID{}
	switch srcs := vData.Sources.Source.(type) {
	case []interface{}:
		for _, src := range srcs {
			m := src.(map[string]interface{})
			if s, ok := m["#text"].(string); ok {
				srcID, err := viafSplitSourceID(s)
				if err != nil {
					fmt.Println(err)
					continue
				}
				vData.SourceIDs = append(vData.SourceIDs, srcID)
			}
		}
	case map[string]interface{}:
		if s, ok := srcs["#text"].(string); ok {
			srcID, err := viafSplitSourceID(s)
			if err != nil {
				fmt.Println(err)
			}
			vData.SourceIDs = append(vData.SourceIDs, srcID)
		}
	default:
		fmt.Printf("unknown interface type %v - (%T)\n", srcs, srcs)
	}
	vData.Sources = Sources{}

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
