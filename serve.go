package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	port      = kingpin.Flag("port", "The port of the HTTP server.").Default("3000").Short('p').Int()
	directory = kingpin.Flag("directory", "The directory to serve").Default(".").Short('d').ExistingDir()
)

func main() {
	kingpin.Parse()
	*directory += "/"
	absDir, err := filepath.Abs(filepath.Dir(*directory))
	if err != nil {
		log.Fatal("Failed to get directory: " + err.Error())
	}
	fmt.Printf("Serving %s on http://%s:%s \n", absDir, "0.0.0.0", strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(absDir))))
}
