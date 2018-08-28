package main

import (
	"fmt"
	"net/http"
)

const PortNumber = 8080
const Host = "localhost:8080"

func main() {
	//Get story from json file
	story := make(map[string]Chapter)
	GetStoryFromFile("story.json", &story)
	//Create serve mux
	mux := NewStoryMux(story)
	//create string to listen and serve
	port := fmt.Sprintf(":%d", PortNumber)
	http.ListenAndServe(port, mux)
}
