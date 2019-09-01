package main

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"github.com/takama/daemon"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

const (

	// name of the service
	name        = "doxagos"
	description = "Liturgical Text Server"

	// port which daemon should be listen
	port = ":9977"
)

var HOMEPATH string
var home *template.Template

//    dependencies that are NOT required by the service, but might be used
var dependencies = []string{"doxaserve.service"}

const homeHtm = `{{define "home"}}
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <title>LIML</title>
        <link rel="stylesheet" href="/static/libs/bootstrap/css/bootstrap.min.css">
        <link rel="stylesheet" href="/static/css/doxagos.css">
    </head>
    <body>
    <div class="container-fluid content">
        <h1>The Liturgy in My Language</h1>
    </div>
    <!-- jquery & Bootstrap JS -->
    <script type="text/javascript" src="/static/libs/jquery/jquery.min.js"></script>
    <script type="text/javascript" src="/static/libs/bootstrap/js/bootstrap.min.js"></script>
    </body>
    </html>
{{end}}
`
var fileServer http.Handler
var agesServer http.Handler
var stdlog, errlog *log.Logger

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}
// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {

	usage := "Usage: myservice install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}
	// Do something, call your goroutines, etc
	var err error
	port := os.Getenv("DOXAPORT")
	if port == "" {
		fmt.Println("Please set the DOXAPORT (server port) environment variable, e.g. DOXAGOSPORT:8080")
		os.Exit(1)
	}
	HOMEPATH = os.Getenv("DOXAPATH")
	if HOMEPATH == "" {
		fmt.Println("Please set the DOXAGOPATH environment variable, e.g. $HOME/doxago/http")
		os.Exit(1)
	}
	homeTemplateFile := filepath.Join(HOMEPATH, "http", "index.gohtml")
	if ! ltfile.FileExists(homeTemplateFile) {

		err = ltfile.WriteFile(homeTemplateFile, homeHtm)
		if err != nil {
			fmt.Println(fmt.Sprintf("create home page template: %s", err.Error()))
			os.Exit(1)
		}
	}
	home, err = template.ParseFiles(homeTemplateFile)
	if err != nil {
		fmt.Println(fmt.Sprintf("parse home page template: %s", err.Error()))
		os.Exit(1)
	}
	agesServer = http.FileServer(http.Dir("/Users/mac002/ages/dcs/public"))
	http.Handle("/ages/", http.StripPrefix("/ages/", agesServer))
	fileServer = http.FileServer(http.Dir(filepath.Join(HOMEPATH, "http")))
	http.HandleFunc("/static/", staticHandler)
	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+port, nil)

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Set up listener for defined host and port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return "Possibly was a problem with the port binding", err
	}

	// set up channel on which to send accepted connections
	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case conn := <-listen:
			go handleClient(conn)
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			stdlog.Println("Stoping listening on ", listener.Addr())
			listener.Close()
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}

	// never happen, but need to complete code
	return usage, nil
}
// Accept a client connection and collect it in a channel
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		client.Write(buf[:numbytes])
	}
}

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func main() {
	srv, err := daemon.New(name, description, dependencies...)
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
func handler(w http.ResponseWriter, r *http.Request) {
	home.ExecuteTemplate(w, "home", nil)
}
func staticHandler(w http.ResponseWriter, r *http.Request) {
	ruri := r.RequestURI
	if strings.HasSuffix(ruri,".css") {
		w.Header().Set("Content-Type", "text/css")
	} else if strings.HasSuffix(ruri,".js") {
		w.Header().Set("Content-Type", "application/javascript")
	}
	fileServer.ServeHTTP(w, r)
}