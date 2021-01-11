package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/victorprb/gophercises/cyoa"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "story.json", "a json file with the story")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}

	jsonDecoder := json.NewDecoder(file)
	var story cyoa.Story
	if err := jsonDecoder.Decode(&story); err != nil {
		log.Fatalf("Could not decode json: %v", err)
	}

	tmpl := template.Must(template.ParseFiles("layout.html"))

	for k, v := range story {
		http.Handle("/"+k, logHandler(chapterHandler(v, tmpl)))
	}

	log.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func chapterHandler(c cyoa.Chapter, tmpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, c)
	})
}

func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method=%s, path=%s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
