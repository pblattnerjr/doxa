package server

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jmoiron/sqlx"
	"github.com/liturgiko/doxa/pkg/models"
	"log"
	"net/http"
	"strings"
	"time"
)
var db *sqlx.DB

func Ages(dbname string, port string)  {
	var err error
	// open the database
	db, err = sqlx.Open("sqlite3", dbname)
	if err != nil {
		log.Println(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/id/{domain}/{topic}/{key}", IDHandler)
	r.HandleFunc("/id/{topic}/{key}", TKHandler)
	r.HandleFunc("/", HomeHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add /id/domain/topic/path to %s. E.g., /gr_gr_cog/actors/Priest", r.URL.Host)
}
func IDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id []string
	id = append(id, vars["domain"])
	id = append(id, vars["topic"])
	id = append(id, vars["key"])

	var rec models.Ltext
	rec.ID = strings.Join(id, "~")
	err := rec.GetRecord(db)
	if err == nil {
		fmt.Fprintf(w, "%s | %s", rec.ID, rec.Value)
	} else {
		fmt.Fprintf(w, "error: %s!", err.Error())
	}

}
func TKHandler(w http.ResponseWriter, r *http.Request) {
	recs := models.NewLtextArray()
	vars := mux.Vars(r)
	err := recs.GetRecordsByTopicKey(db, vars["topic"], vars["key"])
	if err == nil {
		fmt.Fprintf(w, "%s", *recs)
	} else {
		fmt.Fprintf(w, "error: %s!", err.Error())
	}
}
