package ltx2sql

import (
	"database/sql"
	"fmt"
	"github.com/liturgiko/doxa/pkg/models"
	"os"
	"testing"
)

var mapper *Mapper

func TestMain(m *testing.M) {
	var err error
	var theDb *sql.DB
	mapper = &Mapper{}
	theDb, err = sql.Open("sqlite3", "test.db")
	if err!= nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer theDb.Close()
	err = theDb.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	_, err = theDb.Exec("DROP TABLE IF EXISTS ltx")
	if err!= nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	_, err = theDb.Exec(SQLCreateTable)
	if err!= nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	mapper.DB = theDb
	os.Exit(m.Run())
}
func TestMapper_Create(t *testing.T) {
	library := "gr_gr_cog"
	topic := "actors"
	key := "Priest"
	l := models.NewLtx(library, topic, key, "Priest", "", "")
	err := mapper.Create(l)
	if err != nil {
		t.Error(fmt.Sprintf("Create %s: %v", l.ID, err))
	}
	r, err := mapper.ReadById(l.ID)
	if err != nil {
		t.Error(fmt.Sprintf("Read %s: %v", l.ID, err))
	}
	if r.ID != l.ID {
		t.Error(fmt.Sprintf("Read %s, %s: do not match", l.ID, r.ID))
	}
	s := models.NewLtx(library, topic, "deacon", "Deacon", "", "")
	err = mapper.Create(s)
	if err != nil {
		t.Error(fmt.Sprintf("Create %s: %v", s.ID, err))
	}
	records, err := mapper.ReadByLT(library, topic)
	if err != nil {
		t.Error(fmt.Sprintf("ReadByLT %s~%s: %v", library, topic, err))
	}
	if records == nil {
		t.Error(fmt.Sprintf("ReadByLT %s~%s: %v", library, topic, err))
	}
	if len(records) != 2 {
		t.Error(fmt.Sprintf("ReadByLT %s~%s, expected 2, got: %v", library, topic, len(records)))
	}
	records, err = mapper.ReadByTK(topic, key)
	if err != nil {
		t.Error(fmt.Sprintf("ReadByTK %s~%s: %v", topic, key, err))
	}
	if records == nil {
		t.Error(fmt.Sprintf("ReadByTK %s~%s: %v", topic, key, err))
	}
	if len(records) != 1 {
		t.Error(fmt.Sprintf("ReadByTK %s~%s, expected 1, got: %v", topic, key, len(records)))
	}
	u := models.NewLtx("en_us_dedes", topic, "deacon", "Deacon", "", "")
	err = mapper.Create(u)
	if err != nil {
		t.Error(fmt.Sprintf("Create %s: %v", u.ID, err))
	}
	v := "Deacon"
	records, err = mapper.ReadByValue(v)
	if err != nil {
		t.Error(fmt.Sprintf("ReadByValue %s: %v", v, err))
	}
	if records == nil {
		t.Error(fmt.Sprintf("ReadByValue %s: %v", v, err))
	}
	if len(records) != 2 {
		t.Error(fmt.Sprintf("ReadByValue %s, expected 2, got: %v", v, len(records)))
	}
}