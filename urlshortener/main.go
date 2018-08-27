package main

import (
	"flag"
	"net/http"
)

type Data struct {
	Path string
	URL  string
}

var yamlFileFlag string

func init() {
	flag.StringVar(&yamlFileFlag, "file", "paths.yaml", "change yaml file")
}

func main() {
	//Parse flag
	flag.Parse()
	//Default mux
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)
	//Open file that contains yaml data
	yamlData, err := OpenYAMLFile(yamlFileFlag)
	yamlHandler, err := YAMLHandler([]byte(yamlData), mapHandler)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", yamlHandler)
}
