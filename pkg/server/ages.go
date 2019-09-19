package server

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jmoiron/sqlx"
	"github.com/liturgiko/doxa/pkg/models"
	"html/template"
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
	r.HandleFunc("/id/{library}/{topic}/{key}", IDHandler)
	r.HandleFunc("/id/{topic}/{key}", TKHandler)
	r.HandleFunc("/id", HomeHandler)

	r.HandleFunc("/topic/{library}/{topic}", TopicHandler)

	r.HandleFunc("/", HomeHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())}

// HomeHandler provides path information in the event that the requester fails
// to provide a library, topic, key or a topic and key.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add /id/library/topic/path to %s. E.g., /gr_gr_cog/actors/Priest", r.URL.Host)
}
// TKHandler returns the liturgical text that matches the requested library, topic, and key.
func IDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("charset", "utf-8")

	vars := mux.Vars(r)

	var id []string
	id = append(id, vars["library"])
	id = append(id, vars["topic"])
	id = append(id, vars["key"])

	var rec models.Ltext
	rec.ID = strings.Join(id, "~")
	err := rec.GetRecord(db)
	if err == nil {
		recs := models.NewLtextArray()
		recs.Append(&rec)
		t, _ := template.New("webpage").Parse(HtmlTemplate)
		t.Execute(w, recs)
	} else {
		fmt.Fprintf(w, "error: %s!", err.Error())
	}

}
// TKHandler returns liturgical texts for all libraries that have the requested topic and key.
func TKHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("charset", "utf-8")
	recs := models.NewLtextArray()
	vars := mux.Vars(r)
	err := recs.GetRecordsByTopicKey(db, vars["topic"], vars["key"], true)
	if err == nil {
		t, _ := template.New("webpage").Parse(HtmlTemplate)
		t.Execute(w, recs)
	} else {
		fmt.Fprintf(w, "error: %s!", err.Error())
	}
}
// TopicHandler returns liturgical texts for the requested library and topic.
func TopicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("charset", "utf-8")
	recs := models.NewLtextArray()
	vars := mux.Vars(r)
	err := recs.GetRecordsByLibraryTopic(db, vars["library"], vars["topic"], true)
	if err == nil {
		t, _ := template.New("webpage").Parse(HtmlTemplate)
		t.Execute(w, recs)
	} else {
		fmt.Fprintf(w, "error: %s!", err.Error())
	}
}
func Elapsed(start time.Time, msg string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", msg, elapsed)
}
const HtmlTemplate = `<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="UTF-8">
    <title>DB Lookup</title>
  </head>
  <body>
    <table>
      <tbody>
		{{range .}}
			<tr>
			  <td style="padding: 1em; border: 1px solid #cccccc;">{{.ID}}</td>
			  <td style="padding: 1em; border: 1px solid #cccccc;">{{.Value}}</td>
			</tr>
		{{end}}
      <tbody>
    <table>
  </body>
</html>`

