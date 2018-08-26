package main

import "net/http"

var yamlData = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

type Data struct {
	Path string
	URL  string
}

func main() {
	//Default mux
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)
	yamlHandler, err := YAMLHandler([]byte(yamlData), mapHandler)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", yamlHandler)
}
