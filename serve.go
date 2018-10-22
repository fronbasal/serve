package main

import (
	"net/http"
	"path/filepath"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"net"
	"os"
	"strings"
)

var (
	app       = kingpin.New("serve", "Serve is a simple utility to serve a directory via HTTP")
	port      = app.Flag("port", "The port of the HTTP server.").Default("3000").Short('p').Int()
	directory = app.Arg("directory", "The directory to serve").Default(".").ExistingDir()
	verbose   = app.Flag("verbose", "Enable verbose output").Default("true").Short('v').Bool()
)

func main() {
	app.Parse(os.Args[1:])
	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	absDir, err := filepath.Abs(*directory)
	relDir := strings.Replace(absDir, os.Getenv("HOME"), "~", 1) + "/"
	if err != nil {
		log.Fatal("Failed to get directory: ", err)
	}
	log.Infof("Serving %s on http://%s:%s \n", relDir, "0.0.0.0", strconv.Itoa(*port))
	server := &http.Server{Addr: ":" + strconv.Itoa(*port), Handler: http.FileServer(http.Dir(absDir))}
	server.ConnState = func(conn net.Conn, state http.ConnState) {
		if state == http.StateActive {
			log.Debug("Serving client ", conn.LocalAddr().String())
		}
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
