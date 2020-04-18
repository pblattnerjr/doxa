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
// Convert the ID to the format used in the OLW Neo4j database
func (i Id) ToNeoId() string {
	return fmt.Sprintf("%s~%s~%s", i.Domain.ToNeo(), i.Topic, i.Key)
}
// Merge an Ares filename from the ID parts,
// e.g. actors_gr_GR_cog.ares
func (i Id) ToNeoAresFilename() string {
	return fmt.Sprintf("%s_%s.ares", i.Topic, i.Domain.ToAres())
}
// Parses a tilde delimited id into its parts: domain (aka library), topic, and key
func (i Id) Parse(id string) error {
	var err error
	parts := strings.Split(id, "~")
	if len(parts) != 3 {
		return errors.New(fmt.Sprintf("Parse %s: %v", id, err.Error()))
	}

	return err
}

