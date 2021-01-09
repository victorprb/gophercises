package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/victorprb/gophercises/urlshort"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "example.yaml", "a file with list of paths to urls (json or yaml)")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLhandler or JSONHandler using the mapHandler as the
	// fallback
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var handler http.Handler
	switch filename {
	case suffix(filename, ".json"):
		handler, err = urlshort.YAMLHandler(file, mapHandler)
	default:
		handler, err = urlshort.JSONHandler(file, mapHandler)
	}
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	log.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func suffix(s string, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		suffixIndex := strings.LastIndex(s, suffix)

		return strings.TrimPrefix(s, s[:suffixIndex])
	}

	return s
}
