# Package oclcapis

A Go client for _some_ of the _free, public_ Web Services available from OCLC.

See the documentation for this package at <https://godoc.org/github.com/nicomo/oclcapis>

OCLC offers many web services not implemented here, some only accessible to member libraries. The web services are documented at <https://platform.worldcat.org/api-explorer/apis>

## Status

This project is experimental/personal, use at your own risk.

## Covered

- ViafGetIDs finds all source IDs, e.g. LC, DNB, WKP, etc. from a VIAF ID
- ViafGetLCN finds a Library of Congress ID from a VIAF ID
- ViafGetWKP finds a Wikidata ID from a VIAF ID

## Not covered

Everything else, which is a lot.

## To be continued with

- [ViafTranslateSourceID](https://platform.worldcat.org/api-explorer/apis/VIAF/AuthorityCluster/TranslateSourceID): Translate a SourceID (identifier for an original source record at a specific institution) to a VIAF URI. For instance _/viaf/sourceID/DNB%7c1034425390_

- [WorldCatIdentitiesRead](https://platform.worldcat.org/api-explorer/apis/worldcatidentities/identity/Read): Read a single identity record in WorldCat Identities, for example <http://www.worldcat.org/identities/lccn-n79032879/>
