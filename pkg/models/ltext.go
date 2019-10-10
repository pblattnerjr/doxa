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

// Prefix for generating SQL for db read
var ReadPrefix = `PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
`
var ReadSuffix = `
COMMIT;`

// Schema to create Ltext
var LtextSchema = `CREATE TABLE IF NOT EXISTS ltext (
    id       TEXT PRIMARY KEY,
    topic    TEXT,
    key      TEXT,
    value    TEXT,
    nnp      TEXT,
    nwp      TEXT,
    comment  TEXT,
    redirect TEXT,
    FOREIGN KEY(redirect) REFERENCES ltext(id));`

// SQL to insert Ltext into database.
//  This is used for insertions in an existing database.
var LtextSQLInsert = `INSERT OR REPLACE INTO ltext (id, topic, key, value, nnp, nwp, comment, redirect) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

// SQL for load of db via .read
// This is used when we are creating a database by reading ares files.
var ReadSQLInsert = "INSERT INTO ltext VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');\n"

// SQL to find Ltext by ID.
// Because sometimes the value is empty and instead there is a redirect,
// We have to do an inner join to do a lookup.  Without the b.id = case in the
// where clause, we would only get a value back if there was a redirect.
var LtextSQLReadByID = `SELECT a.id as id, 
   CASE WHEN LENGTH(a.redirect) = 0 THEN a.value 
       ELSE b.value 
   END as value
FROM ltext a INNER JOIN ltext b 
WHERE a.id  = ?
AND b.id =
    CASE WHEN LENGTH(a.redirect) > 0 THEN a.redirect
    ELSE a.id
    END
`

// SQL to find Ltext records where ID is like a supplied pattern
// Because sometimes the value is empty and instead there is a redirect,
// We have to do an inner join to do a lookup.  Without the b.id = case in the
// where clause, we would only get a value back if there was a redirect.
var LtextSQLReadWhereIdLike = `SELECT a.id as id, 
   CASE WHEN LENGTH(a.redirect) = 0 THEN a.value 
       ELSE b.value 
   END as value
FROM ltext a INNER JOIN ltext b 
WHERE a.id  like ?
AND b.id =
    CASE WHEN LENGTH(a.redirect) > 0 THEN a.redirect
    ELSE a.id
    END
ORDER BY id
;
`

// SQL to find Ltext records where ID is like a supplied pattern and value not blank
// Because sometimes the value is empty and instead there is a redirect,
// We have to do an inner join to do a lookup.  Without the b.id = case in the
// where clause, we would only get a value back if there was a redirect.
var LtextSQLReadWhereIdLikeValueNotBlank = `SELECT a.id as id, 
   CASE WHEN LENGTH(a.redirect) = 0 THEN a.value 
       ELSE b.value 
   END as value
FROM ltext a INNER JOIN ltext b 
WHERE a.id  like ?
AND (length(a.value) > 0 or length(a.redirect) >0 )
AND b.id =
    CASE WHEN LENGTH(a.redirect) > 0 THEN a.redirect
    ELSE a.id
    END
ORDER BY id;`

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
	row := db.QueryRow(LtextSQLReadWhereIdLike, l.ID)
	err := row.Scan(&l.ID, &l.Value)
	if err != nil {
		return err
	}
	return err
}

// UpdateRecord updates the record with the values contained in the Struct.
// If the record has a redirect, update the value of the redirect instead
func (l *Ltext) UpdateRecord(db *sqlx.DB) error {
	return errors.New("not implemented")
}

func (l *Ltext) DeleteRecord(db *sqlx.DB) error {
	return errors.New("not implemented")
}

func (l *Ltext) CreateRecord(db *sqlx.DB) error {
	_, err := db.Exec(LtextSQLInsert, l.ID, l.Topic, l.Key, l.Value, l.NNP, l.NWP, l.Comment, l.Redirect)
	return err
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

