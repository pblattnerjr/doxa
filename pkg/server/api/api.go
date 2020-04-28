// Package server provides a REST API for liturgical text searching and processing.
// TODO: api versioning based on https://dev.to/geosoft1/versioning-your-api-in-go-1g4h
// 23-11-2019 BigMac: I stubbed out an example of versioning.  It works fine, but I
// am not happy with the current approach of naming handlexV1 vs handlexV2 as handler names.
// TODO: look into gofiber and see if is a better solution than current: https://github.com/gofiber
package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/liturgiko/doxa/pkg/db/ltx2sql"
	"html/template"
	"log"
	"net/http"
	"time"
)

type server struct {
	ltxMapper *ltx2sql.LtxMapper // there should be a map of mappers here
	router    *mux.Router
	api       *mux.Router
	api1      *mux.Router
	api2      *mux.Router
	http      *http.Server
}

var t *template.Template

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
			  <td style="padding: 1em; border: 1px solid #cccccc;">{{.Left}}</td>
			</tr>
		{{end}}
      <tbody>
    <table>
  </body>
</html>`

func init() {
	t, _ = template.New("webpage").Parse(HtmlTemplate)
}
func Serve(dbname, port string) {
	var err error

	// open the database
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
	mapper := ltx2sql.LtxMapper{}
	mapper.DB = db
	srv := server{}
	srv.ltxMapper = &mapper
	srv.router = mux.NewRouter()
	srv.router.Headers().HeadersRegexp("Content-Type", "text/css")
	//	srv.router.Headers().HeadersRegexp("Content-Type", "text/(html|csv|javascript|plain)")
	srv.api = srv.router.PathPrefix("/api").Subrouter()
	srv.api1 = srv.api.PathPrefix("/v1").Subrouter()
	// we don't actually have a version 2 yet, but this is how it will be handled
	srv.api2 = srv.api.PathPrefix("/v2").Subrouter()
	srv.routes() // set the routes for the router

	srv.http = &http.Server{
		Handler:      srv.router,
		Addr:         "127.0.0.1:" + port ,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("doxa api is listening on http://127.0.0.1:%s\n", port)
	log.Printf("doxa db is at %s\n", dbname)
	log.Fatal(srv.http.ListenAndServe())
}
