package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ares"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
	"strings"
)

// Struct for liturgical text
// This struct deliberately breaks the rules of
// normalization by redundantly including the
// topic and key as separate fields, and in the
// ID field. This is for convenience when creating
// queries.
type Ltext struct {
	ID       string
	Topic    string
	Key      string
	Value    string
	NNP      string
	NWP      string
	Comment  string
	Redirect string
}
// Parses the ID field of the Ltext struct
// and returns an Id struct with the parts
// that make up an ID:
// ISO language code
// ISO country code
// Realm
// Topic
// Key
func (l Ltext) ToId() (Id, error) {
	var result Id
	if len(l.ID) == 0 {
		return result, errors.New("error: ID is empty")
	}
	parts := strings.Split(l.ID, "~")
	if len(parts) != 3 {
		return result, errors.New("error: ID does not have two ~ delimiters")
	}
	domainParts := strings.Split(parts[0],"_")
	if len(domainParts) != 3 {
		return result, errors.New("error: Domain does not have two _ delimiters")
	}
	result.Language = domainParts[0]
	result.Country = domainParts[1]
	result.Realm = domainParts[2]
	result.Topic = parts[1]
	result.Key = parts[2]
	return result, nil
}
// An array of liturgical text records
type LtextArray []Ltext

// Convert an array of liturgical text records to a Json string.
// This is useful for sending a response over http(s).
func (l *LtextArray) ToJson() (string, error) {
	json, err := json.MarshalIndent(l,"","  ")
	if err != nil {return "", err}
	return fmt.Sprintf("%s", json), err
}
// Convert a liturgical text record to a Json string.
// This is useful for sending a response over http(s).
func (l *Ltext) ToJson() (string, error) {
	json, err := json.MarshalIndent(l,"","  ")
	if err != nil {return "", err}
	return fmt.Sprintf("%s", json), err
}

func LineParts2Ltext (out chan<- Ltext, in <-chan ares.LineParts) {

	for lineParts := range in {

		var ltext Ltext

		ltext.ID = ltstring.ToId(
			lineParts.Language,
			lineParts.Country,
			lineParts.Realm,
			lineParts.Topic,
			lineParts.Key,
		)
		ltext.Topic = lineParts.Topic
		ltext.Key = lineParts.Key
		if lineParts.HasValue {
			ltext.Value = lineParts.Value
			ltext.NWP = ltstring.ToNwp(lineParts.Value)
			ltext.NNP = ltstring.ToNnp(ltext.NWP)
		}
		if lineParts.HasComment {
			ltext.Comment = lineParts.Comment
		}
		if lineParts.IsRedirect {
			ltext.Redirect = lineParts.Redirect
		}
		out <- ltext
	}

}