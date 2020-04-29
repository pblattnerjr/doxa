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
	"os"
	"os/exec"
	"runtime"
	"time"
)

var ctx context.Context
var cancel context.CancelFunc

/**
ServeLocal runs an http server for the Doxa web app.
dbpath is the file path to the sqlite3 database.
sitepath is the file path to the local generated website.
apport is the port to be used for the web app.
apiport is the port for the REST api.

Because many users are not use to working with URLs that have an explicit port,
and because they might forget to issue the quit command from the web app,
two things are done:
1. When the server starts, it automatically opens a browser using the local
   address and port.
2. Before the server starts, it issues a quit via http.Get() to shutdown any
   locally running instance.  Then starts this instance up.
 */
func ServeLocal(dbpath, sitepath, appport, apiport string) {
		log.Println("Shutting down any previous local instance...")
		_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/quit",appport))
		if err != nil {
			log.Println("No local instance running.")
		} else {
			log.Println("Local instance found and stopped.")
		}
		time.Sleep(2 * time.Second)
	ctx, cancel = context.WithCancel(context.Background())
	go webapi.Serve(dbpath, apiport)
	srv := &http.Server{Addr: ":" + appport}
	http.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body><h2>Doxa has stopped. You may close this tab. Glory to God for all things!</h2></body></html>")
	cancel()
	})
	http.Handle("/", http.FileServer(rice.MustFindBox("public").HTTPBox()))
	// set up handler for local generated website
	if _, err := os.Stat(sitepath); os.IsNotExist(err) {
		log.Printf("Local website path does not exist: %s\n", sitepath)
	} else {
		fs := http.FileServer(http.Dir(sitepath))
		http.Handle("/site/", http.StripPrefix( "/site", fs))
		log.Printf("generated site is at http://127.0.0.1:%s/site/",appport)
	}
	log.Printf("doxa app is running at http://127.0.0.1:%s\n", appport)
	openBrowser(fmt.Sprintf("http://127.0.0.1:%s",appport))

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
func ServeGeneratedSite(path, port string) {
	// add a handler for the local generated web site
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Local website path does not exist: %s\n",path)
	} else {
		fs := http.FileServer(http.Dir(path))
		http.Handle("/", fs)
		log.Printf("generated site is at http://127.0.0.1:%s",port)
		err := http.ListenAndServe(":" + port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
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