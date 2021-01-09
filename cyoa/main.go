package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	data := readFile(filename)

	parsedStory, err := parseStory(data)
	if err != nil {
		log.Fatalf("Could not parse json: %v", err)
	}

	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, parsedStory["intro"])
	})

	http.HandleFunc("/new-york", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, parsedStory["new-york"])
	})

	http.HandleFunc("/denver", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, parsedStory["denver"])
	})

	// for arc, storyArc := range parsedStory {
	// 	http.HandleFunc("/"+arc, func(w http.ResponseWriter, r *http.Request) {
	// 		tmpl.Execute(w, storyArc)
	// 	})
	// }

	// mux.HandleFunc("/", storyHandler(parsedStory["intro"]))
	// mux.HandleFunc("/new-york", storyHandler(parsedStory["new-york"]))

	log.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func storyHandler(a arc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, a.Title)
	}
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
