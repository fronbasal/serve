package main

import (
	"net/http"
	"path/filepath"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
	"net"
	log "github.com/sirupsen/logrus"
)

var (
	port      = kingpin.Flag("port", "The port of the HTTP server.").Default("3000").Short('p').Int()
	directory = kingpin.Arg("directory", "The directory to serve").Default(".").ExistingDir()
	verbose   = kingpin.Flag("verbose", "Enable verbose output").Default("true").Short('v').Bool()
)

func main() {
	kingpin.CommandLine.Name = "serve"
	kingpin.CommandLine.Author("Daniel Malik <mail@fronbasal.de>")
	kingpin.Parse()
	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	*directory += "/"
	absDir, err := filepath.Abs(filepath.Dir(*directory))
	relDir := strings.Replace(absDir, os.Getenv("HOME"), "~", 1) + "/"
	if err != nil {
		log.Fatal("Failed to get directory: ", err.Error())
	}
	log.Infof("Serving %s on http://%s:%s \n", relDir, "0.0.0.0", strconv.Itoa(*port))
	server := &http.Server{Addr: ":" + strconv.Itoa(*port), Handler: http.FileServer(http.Dir(absDir))}
	server.ConnState = func(conn net.Conn, state http.ConnState) {
		if state == http.StateActive {
			log.Debug("Serving client ", conn.LocalAddr().String())
		}
	}
	if server.ListenAndServe() != nil {
		log.Fatal("Failed to start server")
	}
}
