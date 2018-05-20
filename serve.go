package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	port      = kingpin.Flag("port", "The port of the HTTP server.").Default("3000").Short('p').Int()
	directory = kingpin.Flag("directory", "The directory to serve").Default(".").Short('d').ExistingDir()
)

func main() {
	kingpin.Parse()
	fmt.Printf("Serving %s on http://%s:%s \n", *directory, "0.0.0.0", strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*directory))))
}
