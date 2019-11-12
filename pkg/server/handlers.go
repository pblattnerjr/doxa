package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/liturgiko/doxa/pkg/models"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
)

// handleHome provides path information in the event that the requester fails
// to provide a library, topic, key or a topic and key.
func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Add /id/library/topic/key to %s, e.g., id/gr_gr_cog/actors/Priest to view for specific key.", s.http.Addr)
		fmt.Fprintf(w, "\nor\nAdd /id/topic/key to %s, e.g., id/actors/Priest to view for all libraries.", s.http.Addr)
		fmt.Fprintf(w, "\nor\nAdd /topic/library/topic to %s, e.g., topic/gr_gr_cog/actors to view all keys for that library and topic.", s.http.Addr)
	}
}
// handleID returns the liturgical text that matches the requested library, topic, and key.
func (s *server) handleID(files ...string) http.HandlerFunc {
	var (
		init sync.Once
		tpl  *template.Template
		err  error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("charset", "utf-8")

		init.Do(func(){ // one time only, parse the template files
			tpl, err = template.ParseFiles(files...)
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vars := mux.Vars(r)

		var id []string
		id = append(id, vars["library"])
		id = append(id, vars["topic"])
		id = append(id, vars["key"])

		rec, err := s.mapper.ReadById(strings.Join(id, "~"))
		if err == nil {
			recs := models.NewLtxArray()
			recs.Append(rec)
			tpl.Execute(w, recs)
		} else {
			log.Println(err.Error())
			fmt.Fprintf(w, "%s", "Not found")
		}
	}
}
// TKHandler returns liturgical texts for all libraries that have the requested topic and key.
// If the query ?empty=true is set, then it will include records with empty values.
func (s *server) handleTK() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("charset", "utf-8")
		vars := mux.Vars(r)
		recs, err := s.mapper.ReadByTK(vars["topic"], vars["key"], false)
		if err == nil {
			t.Execute(w, recs)
		} else {
			log.Println(err.Error())
			fmt.Fprintf(w, "%s", "Not found")
		}
	}
}
// TopicHandler returns liturgical texts for the requested library and topic.
func (s *server) handleTopic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("charset", "utf-8")
		vars := mux.Vars(r)
		recs, err := s.mapper.ReadByLT(vars["library"], vars["topic"], false)
		if err == nil {
			t.Execute(w, recs)
		} else {
			log.Println(err.Error())
			fmt.Fprintf(w, "%s", "Not found")
		}
	}
}