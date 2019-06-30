package oclcapis

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
