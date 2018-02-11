package main

import (
	"net/http"
	"encoding/json"
	"os"
	"fmt"
	"flag"

	"github.com/viveksyngh/gophercises/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file with all CYOA stroy")
	flag.Parse()

	fmt.Println("Using JSON file : ", *filename)
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(file)
	var story cyoa.Story
	err = d.Decode(&story); if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v\n", story)

	storyhandler, err := cyoa.StoryHandler(story)
	http.ListenAndServe(":8081", storyhandler)

}