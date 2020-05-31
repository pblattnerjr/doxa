/**
Package models provides the structs that are serialized to datastores.
The serialization and retrieval is handled by data mapper packages.
For example, models.ltx (the struct for liturgical text), has a
data mapper package titled ltx2sql for serialization to sqlite3.
 */
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liturgiko/doxa/pkg/ages/ares"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
	"log"
	"strings"
)

const idDelimiter = "/"

// Struct for liturgical text
// This struct deliberately breaks the rules of
// normalization by redundantly including the
// library, topic, and key as separate fields, and in the
// ID field. This is for convenience when creating
// queries.  The Active Record Pattern is used for CRUD operations
// on this struct via receiver functions.
type Ltx struct {
	ID           string `json:"id"`
	Library      string  `json:"library"`
	Topic        string  `json:"topic"`
	Key          string  `json:"key"`
	Value        string  `json:"value"`
	NNP          string  `json:"nnp"`
	NWP          string  `json:"nwp"`
	Comment      string  `json:"comment"`
	Redirect     string  `json:"redirect"`
	CreatedWhen  string  `json:"createdWhen"`
	ModifiedWhen string  `json:"modifiedWhen"`
}
// An array of liturgical text records
type LtxArray []Ltx

// Prefix for generating SQL for db read
var ReadPrefix = `PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
`
var ReadSuffix = `
COMMIT;`

// Schema to create Ltx
var LtxSchema = `CREATE TABLE IF NOT EXISTS ltx (
    id            TEXT PRIMARY KEY,
    library       TEXT,
    topic         TEXT,
    key           TEXT,
    value         TEXT,
    nnp           TEXT,
    nwp           TEXT,
    comment       TEXT,
    redirect      TEXT,
    createdWhen   TEXT,
    modifiedWhen  TEXT,
    FOREIGN KEY(redirect) REFERENCES ltx(id));`

// SQL to insert Ltx into database.
//  This is used for insertions in an existing database.
var LtxSQLInsert = `INSERT OR REPLACE INTO ltx (id, topic, key, value, nnp, nwp, comment, redirect) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

// SQL for load of db via .read
// This is used when we are creating a database by reading ares files.
var ReadSQLInsert = "INSERT INTO ltx VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');\n"

// SQL to find Ltx by ID.
// Because sometimes the value is empty and instead there is a redirect,
// We have to do an inner join to do a lookup.  Without the b.id = case in the
// where clause, we would only get a value back if there was a redirect.
var LtxSQLReadByID = `SELECT a.id as id, 
   CASE WHEN LENGTH(a.redirect) = 0 THEN a.value 
       ELSE b.value 
   END as value
FROM ltx a INNER JOIN ltx b 
WHERE a.id  = ?
AND b.id =
    CASE WHEN LENGTH(a.redirect) > 0 THEN a.redirect
    ELSE a.id
    END
`

// SQL to find Ltx records where ID is like a supplied pattern
// Because sometimes the value is empty and instead there is a redirect,
// We have to do an inner join to do a lookup.  Without the b.id = case in the
// where clause, we would only get a value back if there was a redirect.
var LtxSQLReadWhereIdLike = `SELECT a.id as id, 
   CASE WHEN LENGTH(a.redirect) = 0 THEN a.value 
       ELSE b.value 
   END as value
FROM ltx a INNER JOIN ltx b 
WHERE a.id  like ?
AND b.id =
    CASE WHEN LENGTH(a.redirect) > 0 THEN a.redirect
    ELSE a.id
    END
ORDER BY id
;
`

// SQL to find Ltx records where ID is like a supplied pattern and value not blank
// Because sometimes the value is empty and instead there is a redirect,
// We have to do an inner join to do a lookup.  Without the b.id = case in the
// where clause, we would only get a value back if there was a redirect.
var LtxSQLReadWhereIdLikeValueNotBlank = `SELECT a.id as id, 
   CASE WHEN LENGTH(a.redirect) = 0 THEN a.value 
       ELSE b.value 
   END as value
FROM ltx a INNER JOIN ltx b 
WHERE a.id  like ?
AND (length(a.value) > 0 or length(a.redirect) >0 )
AND b.id =
    CASE WHEN LENGTH(a.redirect) > 0 THEN a.redirect
    ELSE a.id
    END
ORDER BY id;`

// Parses the ID field of the Ltx struct
// and returns an Id struct with the parts
// that make up an ID:
// ISO language code
// ISO country code
// Realm
// Topic
// Key
func (l Ltx) ToId() (Id, error) {
	var result Id
	if len(l.ID) == 0 {
		return result, errors.New("error: ID is empty")
	}
	parts := strings.Split(l.ID, idDelimiter)
	if len(parts) != 3 {
		return result, errors.New("error: ID does not have two delimiters")
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
func (l *LtxArray) ToJson() (string, error) {
	json, err := json.MarshalIndent(l,"","  ")
	if err != nil {return "", err}
	return fmt.Sprintf("%s", json), err
}
// Convert a liturgical text record to a Json string.
// This is useful for sending a response over http(s).
func (l *Ltx) ToJson() (string, error) {
	json, err := json.MarshalIndent(l,"","  ")
	if err != nil {return "", err}
	return fmt.Sprintf("%s", json), err
}

func LineParts2Ltx (out chan<- Ltx, in <-chan ares.LineParts) {

	for lineParts := range in {

		var ltx Ltx

		ltx.ID = ltstring.ToId(
			lineParts.Language,
			lineParts.Country,
			lineParts.Realm,
			lineParts.Topic,
			lineParts.Key,
		)
		ltx.Topic = lineParts.Topic
		ltx.Key = lineParts.Key
		if lineParts.HasValue {
			ltx.Value = lineParts.Value
			ltx.NWP = ltstring.ToNwp(lineParts.Value)
			ltx.NNP = ltstring.ToNnp(ltx.NWP)
		}
		if lineParts.HasComment {
			ltx.Comment = lineParts.Comment
		}
		if lineParts.IsRedirect {
			ltx.Redirect = lineParts.Redirect
		}
		out <- ltx
	}

}

// SetValue sets the value, and the Normalized With Punctuation (NWP) and Normalized without Punctuation (NWP) properties.
func (l *Ltx) SetValue(value string) {
	l.Value = value
	l.NWP = ltstring.ToNwp(value)
	l.NNP = ltstring.ToNnp(value)
	if len(value) > 0 { // value and redirect are mutually exclusive.
		l.Redirect = ""
	}
}
// SetRedirect sets the redirect.  If length > 0, calls SetValue, setting it to ""
// because value and redirect are mutually exclusive.
func (l *Ltx) SetRedirect(to string) {
	l.Redirect = to
	if len(to) > 0 { // value and redirect are mutually exclusive.
		l.SetValue("")
	}
}
// nnp and nwp properties are set correctly.
func NewLtx(library, topic, key, value, comment, redirect string) *Ltx {
	d := Domain{}
	err := d.Parse(library)
	if err != nil {
		log.Println(err.Error())
	}
	id := &Id{}
	id.Domain = d
	id.Topic = topic
	id.Key = key
	l := Ltx{}
	l.ID = id.ToNeoId()
	l.Topic = topic
	l.Key = key
	l.Comment = comment
	l.Redirect = redirect
	l.SetValue(value)
	return &l
}

func (l *LtxArray) Append(i *Ltx) {
	*l = append(*l, *i)
}
func NewLtxArray() *LtxArray {
	var l LtxArray
	return &l
}

