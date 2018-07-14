package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
)

var (
	port      = kingpin.Flag("port", "The port of the HTTP server.").Default("3000").Short('p').Int()
	directory = kingpin.Arg("directory", "The directory to serve").Default(".").ExistingDir()
)

func main() {
	kingpin.Parse()
	*directory += "/"
	absDir, err := filepath.Abs(filepath.Dir(*directory))
	relDir := strings.Replace(absDir, os.Getenv("HOME"), "~", 1) + "/"
	if err != nil {
		log.Fatal("Failed to get directory: " + err.Error())
	}
	fmt.Printf("Serving %s on http://%s:%s \n", relDir, "0.0.0.0", strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(absDir))))
}
