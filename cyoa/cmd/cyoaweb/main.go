package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/victorprb/gophercises/cyoa"
)

func main() {
	var filename string
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "the server listen addr")
	flag.StringVar(&filename, "file", "story.json", "a json file with the story")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}

	story, err := cyoa.JSONStory(file)
	if err != nil {
		log.Fatalf("Could not decode json: %v", err)
	}

	storyHandler := cyoa.NewHandler(story)
	log.Println("Starting the server")
	log.Fatal(http.ListenAndServe(addr, logHandler(storyHandler)))
}

func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method=%s, path=%s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
