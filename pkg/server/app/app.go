// Package server provides a GUI for doxago.
package app

import (
	"context"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	webapi "github.com/liturgiko/doxa/pkg/server/api"
	"io"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

var ctx context.Context
var cancel context.CancelFunc

/**
Serve runs an http server for the Doxa web app.
db is the file path to the sqlite3 database.
apport is the port to be used for the web app.
apiport is the port for the REST api.
cloud = true means we are running in a cloud server
      = false means we are running locally

Because many users are not use to working with URLs that have an explicit port,
and because they might forget to issue the quit command from the web app,
two things are done:
1. When the server starts, it automatically opens a browser using the local
   address and port.
2. Before the server starts, it issues a quit via http.Get() to shutdown any
   locally running instance.  Then starts this instance up.
 */
func Serve(db , appport, apiport string, cloud bool) {
	if ! cloud { // close other local instance in case user forgot to.
		log.Println("Shutting down any previous local instance...")
		_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/quit",appport))
		if err != nil {
			log.Println("No local instance running.")
		} else {
			log.Println("Local instance found and stopped.")
		}
		time.Sleep(2 * time.Second)
	}
	ctx, cancel = context.WithCancel(context.Background())
	go webapi.Serve(db, apiport)
	srv := &http.Server{Addr: ":" + appport}
	http.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		if cloud {
			io.WriteString(w, "<html><body><h2>To quit, close this tab or the browser. Glory to God for all things!</h2></body></html>")
		} else {
			io.WriteString(w, "<html><body><h2>Doxa has stopped. You may close this tab. Glory to God for all things!</h2></body></html>")
			cancel()
		}
	})
	http.Handle("/", http.FileServer(rice.MustFindBox("public").HTTPBox()))
	log.Printf("doxa app is running at http://127.0.0.1:%s\n", appport)
	if ! cloud { // so localhost user does not have to type the url and port
		openBrowser(fmt.Sprintf("http://127.0.0.1:%s",appport))
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	<-ctx.Done()
	if err := srv.Shutdown(ctx); err != nil && err != context.Canceled {
		log.Println(err)
	}
	log.Println("done.")
}
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}