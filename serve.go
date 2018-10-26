package main

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"
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
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	// Resolve the absolute path from the given directory
	absDir, err := filepath.Abs(*directory)
	// Produce a relative directory ($HOME -> ~)
	relDir := strings.Replace(absDir, os.Getenv("HOME"), "~", 1) + "/"
	if err != nil {
		logrus.Fatal("Failed to get directory: ", err)
	}
	logrus.Infof("Serving %s on http://%s:%s \n", relDir, "0.0.0.0", strconv.Itoa(*port))

	// Create HTTP server
	server := &http.Server{Addr: ":" + strconv.Itoa(*port), Handler: http.FileServer(http.Dir(absDir)), ConnState: func(conn net.Conn, state http.ConnState) {
		if state == http.StateActive {
			logrus.Debug("Serving client %s", conn.RemoteAddr())
		}
	}}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal("Failed to start server: ", err)
	}
}
