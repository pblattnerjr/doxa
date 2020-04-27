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
)

var ctx context.Context
var cancel context.CancelFunc

func Serve(db , appport, apiport string, cloud bool) {
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