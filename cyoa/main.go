package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type story map[string]arc

type arc struct {
	Title   string
	Story   []string
	Options []arcOption
}

type arcOption struct {
	Text string
	Arc  string
}

func main() {
	var filename string
	flag.StringVar(&filename, "file", "story.json", "a json file with the story")
	flag.Parse()

	storyJSON := readFile(filename)
	parsedStory, err := parseStory(storyJSON)
	if err != nil {
		log.Fatalf("Could not parse json: %v", err)
	}

	tmpl := template.Must(template.ParseFiles("layout.html"))

	for arcName, arc := range parsedStory {
		http.Handle("/"+arcName, logHandler(arcHandler(arc, tmpl)))
	}

	log.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}

	return data
}

func parseStory(data []byte) (story, error) {
	var s story
	err := json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func arcHandler(a arc, tmpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, a)
	})
}

func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method=%s, path=%s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
