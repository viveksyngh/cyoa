package cyoa

import (
	"log"
	"strings"
	"fmt"
	"net/http"
	"html/template"
)

type Story map[string]Chapter 

type Chapter struct {
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}


func StoryHandler(stories map[string]Chapter) (http.HandlerFunc, error) {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t, err := template.ParseFiles("templates/chapter.html") 
		if(r.URL.Path == "/") {
			chapter := stories["intro"]
			if err != nil {
				log.Fatal(err)
				return
			}
			t.Execute(w, chapter)
			// json.NewEncoder(w).Encode(chapter)
			return
		}

		path := r.URL.Path
		chapterName := strings.Replace(path, "/", "", -1)
		chapter, ok := stories[chapterName]
		if !ok {
			fmt.Println(w, "Chapter not found")
			return
		}
		// json.NewEncoder(w).Encode(chapter)
		t.Execute(w, chapter)
	}, nil
}
