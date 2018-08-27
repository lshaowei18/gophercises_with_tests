package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// type Story struct {
// 	Intro     Chapter
// 	NewYork   Chapter
// 	Debate    Chapter
// 	SeanKelly Chapter
// 	MarkBates Chapter
// 	Denver    Chapter
// 	Home      Chapter
// }

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func GetStoryFromFile(file string, story *map[string]Chapter) error {
	var returnErr error
	dataByte, err := ioutil.ReadFile(file)
	if err != nil {
		returnErr = err
	}
	err = json.Unmarshal(dataByte, story)
	if err != nil {
		returnErr = err
	}
	return returnErr
}

func (chapter Chapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tpl.html"))
	err := tpl.Execute(w, chapter)
	if err != nil {
		log.Fatalln(err)
	}
}

func NewStoryMux(story map[string]Chapter) *http.ServeMux {
	mux := http.NewServeMux()
	a := story["home"]
	fmt.Println(a)
	mux.Handle("/", a)
	return mux
}
