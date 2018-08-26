package main

import (
	"fmt"
	"log"

	"github.com/go-yaml/yaml"
)

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
	data := []Data{}
	err := yaml.Unmarshal([]byte(yamlData), &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(data)
}
