package oclcapis

import (
	"encoding/xml"
	"errors"
	"fmt"
)

// WCIdentity stores the unmarshalled result
// from the WorldCat Identities web service
type WCIdentity struct {
	Pnkey    string   `xml:"pnkey"`
	AudLevel AudLevel `xml:"audLevel"`
	NameInfo NameInfo `xml:"nameInfo"`
}

// AudLevel is embedded in WCIdentity
// represents the "audience level"
// of the author's works
type AudLevel struct {
	Avg Avg `xml:"avg"`
}

// Avg is embedded in AudLevel
type Avg struct {
	Text  string `xml:",chardata"`
	Level string `xml:"level"`
}

// NameInfo is embedded in WCIdentity
type NameInfo struct {
	Type          string    `xml:"type,attr"`
	RawName       RawName   `xml:"rawName"`
	Languages     Languages `xml:"languages"`
	BirthDate     string    `xml:"birthDate"`
	Dates         Dates     `xml:"dates"`
	TotalHoldings string    `xml:"totalHoldings"`
	WorkCount     string    `xml:"workCount"`
	RecordCount   string    `xml:"recordCount"`
}

// RawName is embedded in NameInfo
type RawName struct {
	Suba string `xml:"suba"`
}

// Languages is embedded in NameInfo
type Languages struct {
	Count string `xml:"count,attr"`
	Lang  []Lang `xml:"lang"`
}

// Lang is embedded in Languages
type Lang struct {
	Code  string `xml:"code,attr"`
	Count string `xml:"count,attr"`
}

// Dates is embedded in NameInfo
type Dates struct {
	Date []Date `xml:"date"`
}

// Date is embedded in Dates
type Date struct {
	Text  string `xml:",chardata"`
	Count string `xml:"count,attr"`
}

const baseWCIdentitiesURL = "http://www.worldcat.org/identities/"

// WCIdentitiesRead calls the WorldCat Identities web service
// for a given LC number (i.e. authority number)
func WCIdentitiesRead(input string) (WCIdentity, error) {
	// will store the result
	var wci WCIdentity

	if input == "" {
		return wci, errors.New("input cannot be an empty string")
	}

	// e.g. http://www.worldcat.org/identities/lccn-n2009050322/
	getURL := baseWCIdentitiesURL + input + "/"

	// call WS & put the response into a []byte
	b, err := callWS(getURL)
	if err != nil {
		return wci, err
	}

	// unmarshall
	if err := xml.Unmarshal(b, &wci); err != nil {
		return wci, fmt.Errorf("could not unmarshall response: %v", err)
	}

	return wci, nil

}
