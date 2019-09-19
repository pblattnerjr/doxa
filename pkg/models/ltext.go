package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
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
// An array of liturgical text records
type LtextArray []Ltext

// Schema to create Ltext
var LtextSchema = `CREATE TABLE ltext (
    id       TEXT PRIMARY KEY,
    topic    TEXT,
    key      TEXT,
    value    TEXT,
    nnp      TEXT,
    nwp      TEXT,
    comment  TEXT,
    redirect TEXT);`

// SQL to insert Ltext into database
var LtextSQLInsert = `INSERT INTO ltext (id, topic, key, value, nnp, nwp, comment, redirect) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

// SQL to find Ltext by ID
var LtextSQLReadByID = `SELECT id, value FROM ltext WHERE id = ?`

// SQL to find Ltext records where ID is like a supplied pattern
var LtextSQLReadWhereIdLike = `SELECT id, value FROM ltext WHERE id LIKE ? ORDER BY id`

// SQL to find Ltext records where ID is like a supplied pattern and value not blank
var LtextSQLReadWhereIdLikeValueNotBlank = `SELECT id, value FROM ltext WHERE id LIKE ? AND LENGTH(value) > 0 ORDER BY id`

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

func (l *Ltext) GetRecord(db *sqlx.DB) error {
	row := db.QueryRow(LtextSQLReadByID, l.ID)
	err := row.Scan(&l.ID, &l.Value)
	if err != nil {
		return err
	}
	return err
}

func (l *Ltext) UpdateRecord(db *sqlx.DB) error {
	return errors.New("not implemented")
}

func (l *Ltext) DeleteRecord(db *sqlx.DB) error {
	return errors.New("not implemented")
}

func (l *Ltext) CreateRecord(db *sqlx.DB) error {
	return errors.New("not implemented")
}
// GetRecordsByTopicKey returns records whose ID ends with the requested topic and key.
// If includeEmptyValue = true, then if the record's value field is empty, it will still be returned.
// Otherwise it will be excluded from the results.
func (l *LtextArray) GetRecordsByTopicKey(db *sqlx.DB, topic string, key string, includeEmptyValue bool) error {
	type Record struct {
		id string `db:"id"`
		value string `db:"value"`
	}
	var id []string
	id = append(id, "%")
	id = append(id, topic)
	id = append(id, key)

	like := strings.Join(id, "~")
	if includeEmptyValue {
		return db.Select(l, LtextSQLReadWhereIdLike, like)
	} else {
		return db.Select(l, LtextSQLReadWhereIdLikeValueNotBlank, like)
	}
}
func (l *LtextArray) GetRecordsByLibraryTopic(db *sqlx.DB, library string, topic string, includeEmptyValue bool) error {
	type Record struct {
		id string `db:"id"`
		value string `db:"value"`
	}
	var id []string
	id = append(id, library)
	id = append(id, topic)
	id = append(id, "%")

	like := strings.Join(id, "~")
	if includeEmptyValue {
		return db.Select(l, LtextSQLReadWhereIdLike, like)
	} else {
		return db.Select(l, LtextSQLReadWhereIdLikeValueNotBlank, like)
	}
}
func (l *LtextArray) Append(i *Ltext) {
	*l = append(*l, *i)
}
func NewLtextArray() *LtextArray {
	var l LtextArray
	return &l
}

