/**
Package ltx2sql provides a data mapper between the struct Ltx and a Sqlite database.
There are functions to create, read, update, and delete records.
There are four methods for reading by ID:
  one by the entire ID passed as a tilde delimited string.
  one by library + topic + key, which is converted to an properly formatted ID.
  one by the beginning of the ID (i.e. library + topic)
  one by the end of the ID (i.e. topic + key).
The text value is saved in three forms, each with its own property (value, NNP, NWP).
The text value is saved without change in the value property of the Ltx struct.
A normalized version is saved in the NNP property with no punctuation (NP).
Normalization converts the value to lowercase and removes accents.
Another normalized version is saved in the NWP property, with punctuation (WP).
There are three functions for searching the value.
  one by the value property.
  one by the NNP property.
  one by the NWP property.
 */
package ltx2sql

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/liturgiko/doxa/pkg/models"
	"strconv"
	"strings"
)
// TODO: the various functions in ltx and rdb that handle database transactions need to be moved here

type LtxMapper struct {
	DB *sql.DB
}
// SQL to create the table schema for the struct.  If the table exists,
// it will be left untouched.
var SQLCreateTable = `CREATE TABLE IF NOT EXISTS ltx (
    id       TEXT PRIMARY KEY,
    topic    TEXT,
    key      TEXT,
    value    TEXT,
    nnp      TEXT,
    nwp      TEXT,
    comment  TEXT,
    redirect TEXT,
    FOREIGN KEY(redirect) REFERENCES ltx(id));`

// SQL to insert a row for the struct into a database.
var sqlInsert = `INSERT INTO ltx (id, topic, key, value, nnp, nwp, comment, redirect) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

// SQL to insert or replace (merge) a row for the struct into database.
//  This is used for insertions in an existing database where you want a row to be created if it does not exist, or to be replaced if it does.
var SQLMerge = `INSERT OR REPLACE INTO ltx (id, topic, key, value, nnp, nwp, comment, redirect) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

// SQL to delete a record by id
var SQLDelete = `DELETE ltx WHERE id = $1`

// SQL to record count like id
var SQLCountIDLike = `SELECT COUNT(*) FROM ltx WHERE id like $1`

// SQL to record count like id
var SQLCountTopicsForLibrary = `SELECT COUNT(DISTINCT topic) FROM ltx WHERE id like $1`

// SQL to get distinct topics for a library
var SQLTopicsForLibrary = `SELECT DISTINCT topic FROM ltx WHERE id LIKE $1 ORDER BY topic  COLLATE NOCASE ASC`

// SQL to get keys for a library and topic
var SQLKeysForTopic = `SELECT DISTINCT key FROM ltx WHERE id LIKE $1 ORDER BY key COLLATE NOCASE ASC`

// Database column names for the struct.  Must correspond to the Fields() interface below
func (m *LtxMapper) Columns() string {
	return "id, topic, key, value, nnp, nwp, comment, redirect"
}
// Struct properties (fields).  Must correspond to the Columns() above
func (m *LtxMapper) Fields(l *models.Ltx) []interface{} {
	return []interface{}{&l.ID, &l.Topic, &l.Key, &l.Value, &l.NNP, &l.NWP, &l.Comment, &l.Redirect}
}
// TODO: test this function
// BulkInsert uses a db transaction and commit to insert a stream of ltx into a database
func (m *LtxMapper) BulkInsert(in <-chan *models.Ltx) error {
	var err error
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	go func() error {
		for l := range in {
			_, err = tx.Exec(SQLMerge, l.ID, l.Topic, l.Key, l.Value, l.NNP, l.NWP, l.Comment, l.Redirect)
			if err != nil {
				// TODO: do we need to abort for all types of errors?
				return err
			}
		}
		return err
	}()
	err = tx.Commit()
	return err
}
// Creates or updates a row from the struct in the database table using SQL MERGE
func (m *LtxMapper) Merge(l *models.Ltx) error {
	_, err := m.DB.Exec(SQLMerge, l.ID, l.Topic, l.Key, l.Value, l.NNP, l.NWP, l.Comment, l.Redirect)
	return err
}
// Read (by id) returns a struct populated by reading the database table for the specified id
func (m *LtxMapper) ReadById(id string) (*models.Ltx, error) {
	return m.QueryRow("id = $1", id)
}
// Read (by library, topic, and key) returns a struct populated by reading the database table for the specified library, topic, and key
func (m *LtxMapper) ReadByLTK(library, topic, key string) (*models.Ltx, error) {
	recs, err := m.Query("id = $1", true, fmt.Sprintf("%s~%s~%s", library, topic, key))
	if len(recs) > 0 {
		return recs[0], err
	} else {
		return nil, err
	}

}
// Read (by library and topic) returns a struct populated by reading the database table for the specified library and topic
func (m *LtxMapper) ReadByLT(library, topic string, returnEmpty bool) ([]*models.Ltx, error) {
	return m.Query("id like $1", returnEmpty, fmt.Sprintf("%s~%s~%%", library, topic))
}
// Read (by topic and key) returns a struct populated by reading the database table for the specified topic and key
func (m *LtxMapper) ReadByTK(topic, key string, returnEmpty bool) ([]*models.Ltx, error) {
	return m.Query("id like $1", returnEmpty, fmt.Sprintf("%%~%s~%s", topic, key))
}
// Read (by value) returns a struct populated by reading the database table for the specified substring in the value column.
// Note that the value property stores a text value unchanged.  This means it preserves punctuation and accents.
// There are other methods (ReadByNNP and ReadByNWP) that provide a search against a normalized version of the value.
// id is the record id and can be empty or use %, e.g. gr_gr_cog% instead of a full id
// substring is what is being searched for.
func (m *LtxMapper) ReadByValue(id, substring string) ([]*models.Ltx, error) {
	if len(id) > 0 {
		return m.Query("id LIKE $1 and value LIKE $2 ESCAPE '\\'", true, id, fmt.Sprintf("%%%s%%", substring))
	} else {
		return m.Query("value like $1 ESCAPE '\\'", true, fmt.Sprintf("%%%s%%", substring))
	}
}
// Read (by NNP) returns a struct populated by reading the database table for the specified substring in the nnp column.
// The NNP column is a normalized version of the value, converted to lower case, with accents removed. It has no punctuation (NP)
// id is the record id and can be empty or use %, e.g. gr_gr_cog% instead of a full id
// substring is what is being searched for.
func (m *LtxMapper) ReadByNNP(id, substring string) ([]*models.Ltx, error) {
	if len(id) > 0 {
		return m.Query("id LIKE $1 and nnp LIKE $2 ESCAPE '\\'", true, id, fmt.Sprintf("%%%s%%", substring))
	} else {
		return m.Query("nnp like $1 ESCAPE '\\'", true, fmt.Sprintf("%%%s%%", substring))
	}
}
// Read (by NWP) returns a struct populated by reading the database table for the specified substring in the nnp column.
// The NWP column is a normalized version of the value, converted to lower case, with accents removed. But, it is with punctuation (WP).
// id is the record id and can be empty or use %, e.g. gr_gr_cog% instead of a full id
// substring is what is being searched for.
func (m *LtxMapper) ReadByNWP(id, substring string) ([]*models.Ltx, error) {
	if len(id) > 0 {
		return m.Query("id LIKE $1 and nwp LIKE $2 ESCAPE '\\'", true, id, fmt.Sprintf("%s%%%s%%", id, substring))
	} else {
		return m.Query("nwp like $1 ESCAPE '\\'", true, fmt.Sprintf("%%%s%%", substring))
	}
}
// Returns a struct, if found, populated by reading the database table using the c (condition) and interface values
func (m *LtxMapper) QueryRow(c string, v ...interface{}) (*models.Ltx, error) {
	// TODO: need to modify the SQL to match the one in Ltx that does a join if redirect is populated
	u := &models.Ltx{}
	err := m.DB.QueryRow(
		fmt.Sprintf(`SELECT %s FROM ltx WHERE %s`, m.Columns(), c),
		v...,
	).Scan(m.Fields(u)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
// Returns structs, if found, populated by reading the database table using the c (condition) and interface values
func (m *LtxMapper) Query(c string, returnEmpty bool, v ...interface{}) ([] *models.Ltx, error) {
	var records []*models.Ltx

	// TODO: need to modify the SQL to match the one in Ltx that does a join if redirect is populated
	q := fmt.Sprintf(`SELECT %s FROM ltx WHERE %s`, m.Columns(), c)
	rows, err := m.DB.Query(q,v...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r := &models.Ltx{}
		err = rows.Scan(m.Fields(r)...)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if returnEmpty {
			records = append(records, r)
		} else {
			if len(r.Value) > 0 {
				records = append(records, r)
			}
		}
	}
	err = rows.Err()
	return records, err
}
// Returns an array of the libraries found in the database
func (m *LtxMapper) Libraries() ([] string, error) {
	var records []string
//	rows, err := m.DB.Query("SELECT distinct topic AS Library FROM ltx ORDER BY topic;")
	rows, err := m.DB.Query("SELECT distinct substr(id, 1, pos-1) AS library FROM (SELECT *, instr(id,\"~\") AS pos FROM ltx) ORDER BY id;")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var library string
		err = rows.Scan(&library)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if len(library) > 0 {
			records = append(records, library)
		}
	}
	err = rows.Err()
	return records, err
}
// Returns an array of the topics found in the database for specified library
func (m *LtxMapper) Topics(like string) ([] string, error) {
	var records []string
	rows, err := m.DB.Query(SQLTopicsForLibrary, like+"%")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var topic string
		err = rows.Scan(&topic)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if len(topic) > 0 {
			records = append(records, topic)
		}
	}
	err = rows.Err()
	return records, err
}
// Returns an array of the keys found in the database for specified library and topic
func (m *LtxMapper) Keys(like string) ([] string, error) {
	var records []string
	rows, err := m.DB.Query(SQLKeysForTopic, like+"%")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item string
		err = rows.Scan(&item)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if len(item) > 0 {
			records = append(records, item)
		}
	}
	err = rows.Err()
	return records, err
}
// Returns an array of the distinct values of the specified column
func (m *LtxMapper) Distinct(column, like string) ([] string, error) {
	var records []string
	var q string
	if len(like) > 0 {
		q = fmt.Sprintf(`SELECT distinct %s AS item FROM ltx WHERE id LIKE "%s" ORDER BY %s;`, column, like, column)
	} else {
		q = fmt.Sprintf(`SELECT distinct %s AS item FROM ltx ORDER BY %s;`, column, column)
	}
	rows, err := m.DB.Query(q)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item string
		err = rows.Scan(&item)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if len(item) > 0 {
			records = append(records, item)
		}
	}
	err = rows.Err()
	return records, err
}
// Deletes a row from the struct in the database table for the specified id
func (m *LtxMapper) Delete(id string) error {
	_, err := m.DB.Exec(SQLDelete, id)
	return err
}
// returns the count for id like specified library, topic, key
// Topic and/or key can be empty strings
func (m *LtxMapper) CountTopics(library string) (int, error) {
	var count int
	var sb strings.Builder
	sb.WriteString(library)
	sb.WriteString("~%")
	q := sb.String()
	rows, err := m.DB.Query(SQLCountTopicsForLibrary, q)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		var item string
		err = rows.Scan(&item)
		if err != nil {
			return -1, err
		}
		count, err = strconv.Atoi(item)
		if err != nil {
			return -1, err
		}
	}
	err = rows.Err()
	return count, err
}
// returns the count for id like specified library, topic, key
// Topic and/or key can be empty strings
func (m *LtxMapper) CountKeys(library, topic, key string) (int, error) {
	var count int
	var sb strings.Builder
	if len(key) > 0 {
		sb.WriteString(library)
		sb.WriteString("~")
		sb.WriteString(topic)
		sb.WriteString("~")
		sb.WriteString(key)
	} else if len(topic) > 0 {
		sb.WriteString(library)
		sb.WriteString("~")
		sb.WriteString(topic)
		sb.WriteString("~")
		sb.WriteString("%")
	} else {
		sb.WriteString(library)
		sb.WriteString("~")
		sb.WriteString("%")
	}
	rows, err := m.DB.Query(SQLCountIDLike, sb.String())
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		var item string
		err = rows.Scan(&item)
		if err != nil {
			return -1, err
		}
		count, err = strconv.Atoi(item)
		if err != nil {
			return -1, err
		}
	}
	err = rows.Err()
	return count, err
}
func (m *LtxMapper) CaseSensitiveLike(on bool) error {
	_,err := m.DB.Exec("PRAGMA case_sensitive_like = true;")
	return err
}