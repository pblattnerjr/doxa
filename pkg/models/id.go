package models

import "fmt"

// The unique identifier for a record
// which consists of a domain, topic, and key.
type Id struct {
	Domain
	Topic string
	Key string
}
// Convert the ID to the format used in the OLW Neo4j database
func (i Id) ToNeoId() string {
	return fmt.Sprintf("%s~%s~%s", i.Domain.ToNeo(), i.Topic, i.Key)
}
// Create an Ares filename from the ID parts,
// e.g. actors_gr_GR_cog.ares
func (i Id) ToNeoAresFilename() string {
	return fmt.Sprintf("%s_%s.ares", i.Topic, i.Domain.ToAres())
}


