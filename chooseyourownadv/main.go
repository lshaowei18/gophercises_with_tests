package main

import (
	"net/http"
)

func main() {
	//Get story from json file
	story := make(map[string]Chapter)
	GetStoryFromFile("story.json", &story)
	//Create serve mux
	mux := NewStoryMux(story)
	http.ListenAndServe(":8080", mux)
}
