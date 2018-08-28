package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//struct for the Chapters inside the story
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
		Link string
	} `json:"options"`
}

//List of keys in the story map
/*
new-york
debate
sean-kelly
mark-bates
denver
home
intro
*/

func main() {
	//Decode json
	dataBytes, err := ioutil.ReadFile("story.json")
	if err != nil {
		log.Fatalln(err)
	}
	//Umarshal JSON
	story := make(map[string]Chapter) // Map to hold all the chapters
	err = json.Unmarshal(dataBytes, &story)
	if err != nil {
		log.Fatalln(err)
	}
	//Print out intro story
	PrintChapterText("intro", story)
}

//Get the text of the by chapter name which is the key to the map
func PrintChapterText(name string, story map[string]Chapter) {
	chapter := story[name]
	for _, story := range chapter.Story {
		fmt.Println(story)
	}
}
