package models

import (
	"errors"
	"fmt"
	"strings"
)

// The unique identifier for a record
// which consists of a domain (aka library), topic, and key.
type Id struct {
	Domain
	Topic string
	Key string
}
const IDDelimiter = "/"
func (i Id) IDDelimiter() string {
	return IDDelimiter
}
func (i Id) IsEmpty() bool {
	return len(i.Domain.Country) == 0
}
func (i Id) HasValues() bool {
	return ! i.IsEmpty()
}
// Convert the ID to the format used in the OLW Neo4j database
func (i Id) ToNeoId() string {
	return fmt.Sprintf("%s~%s~%s", i.Domain.ToNeo(), i.Topic, i.Key)
}
// Convert the ID to the format used in the OLW Neo4j database
func (i Id) ToSQLId() string {
	return fmt.Sprintf("%s%s%s%s%s", i.Domain.ToNeo(), IDDelimiter, i.Topic, IDDelimiter, i.Key)
}
// Merge an Ares filename from the ID parts,
// e.g. actors_gr_GR_cog.ares
func (i Id) ToNeoAresFilename() string {
	return fmt.Sprintf("%s_%s.ares", i.Topic, i.Domain.ToAres())
}
func (i Id) ToPath() string {
	return fmt.Sprintf("%s/%s/%s",i.Domain.ToNeo(),i.Topic, i.Key)
}
// Sets ID to specified library, topic, key
func (i *Id) Set(library, topic, key string) error {
	var d Domain
	err := d.Parse(library)
	if err != nil {
		return err
	}
	i.Domain = d
	i.Topic = topic
	i.Key = key
	return nil
}
// Parses a delimited id into its parts: domain (aka library), topic, and key
func (i *Id) Parse(id string) error {
	var err error
	parts := strings.Split(id, IDDelimiter)
	if len(parts) == 3 {
		var d Domain
		d.Parse(parts[0])
		i.Domain = d
		i.Topic = parts[1]
		i.Key = parts[2]
	} else {
		return errors.New(fmt.Sprintf("Parse %s: %v", id, err.Error()))
	}
	return err
}
