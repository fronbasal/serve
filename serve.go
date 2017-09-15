// Serve: a simple CLI to serve a folder to a port on the interwebz. Use intended for development purposes only!
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	port = kingpin.Arg("port", "The port to serve the directory at").Default("3000").Int()
)

func main() {
	kingpin.Parse()
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Serving " + path + " on http://localhost:" + strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(path))))
}
