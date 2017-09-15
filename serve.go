// Serve: a simple CLI to serve a folder to a port on the interwebz. Use intended for development purposes only!
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	path = kingpin.Flag("path", "Provide the directory of the folder which you are intending to serve.").Default(".").ExistingDir()
	port = kingpin.Flag("port", "The port to serve the directory at").Default("3000").Int()
)

func main() {
	kingpin.Parse()
	fmt.Println("Serving " + *path + " on http://localhost:" + strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*path))))
}
