// Package server provides a GUI for doxago.
package app

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	webapi "github.com/liturgiko/doxa/pkg/server/api"
	"log"
	"net/http"
	"time"
)

var router *mux.Router

func Serve(db string, port string) {
	go webapi.Serve(db, "8090")
	if err := run(port); err != nil {
		log.Fatal(err)
	}
}
func run(port string) error {
	http.Handle("/", http.FileServer(rice.MustFindBox("public").HTTPBox()))
	fmt.Printf("doxahugo running on port %s\n",port)
	return http.ListenAndServe(":" + port, nil)
}
func ServeWithMux(port string) {
	router = mux.NewRouter()
	router.HandleFunc("/", handleHome()).Methods("GET")
	http := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + port ,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("APP Server running on %s\n", http.Addr)
	log.Fatal(http.ListenAndServe())
}
func handleHome() http.HandlerFunc {
	return http.FileServer(rice.MustFindBox("public").HTTPBox()).ServeHTTP
}
