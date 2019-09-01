package models

import (
	"fmt"
	"strings"
)

// A Domain provides information about a library.
// It indicates the:
// ISO language code
// ISO country code
// Realm - an arbitrary unique identifier, e.g. of an individual translator, or a metropolis
type Domain struct {
	Language string
	Country string
	Realm string
}
// Formats a Domain into the form required for an AGES ares file
func (d Domain) ToAres() string {
	return fmt.Sprintf("%s_%s_%s",d.Language, strings.ToUpper(d.Country), d.Realm)
}
// Formats a Domain into the form required for OLW's neo4j database
func (d Domain) ToNeo() string {
	return fmt.Sprintf("%s_%s_%s",d.Language, d.Country, d.Realm)
}

