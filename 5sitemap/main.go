package main

import (
	"flag"
	"fmt"
	"log"
)

//Flag variable for URL link
var urlFlag string

func init() {
	flag.StringVar(&urlFlag, "urlOpt", "https://www.calhoun.io", "url for site")
}

func main() {
	flag.Parse()
	baseURL, err := getBaseURL(urlFlag)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(baseURL)
	byteData, err := getDataFromURL(urlFlag)
	if err != nil {
		log.Fatalln(err)
	}
	doc, err := parseURLIntoHTMLNode(byteData)
	if err != nil {
		log.Fatalln(err)
	}
	urls := getAllHrefFromURL(doc)
	fmt.Println(urls)
}
