package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	//Flag
	urlFlag := flag.String("url", "https://calhoun.io", "url that you want to build a sitemap for")
	flag.Parse()
	//get data from url
	byteData, reqURL, err := getDataFromURL(*urlFlag)
	if err != nil {
		log.Fatalln(err)
	}
	//Parse the byte data into doc
	doc, err := parseURLIntoHTMLNode(byteData)
	if err != nil {
		log.Fatalln(err)
	}
	links := getAllHrefFromURL(doc)
	links = parseURLs(links, reqURL)
	for _, link := range links {
		fmt.Println(link)
	}
}

//Parse response body into html node
func parseURLIntoHTMLNode(data []byte) (*html.Node, error) {
	htmlNode, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return htmlNode, nil
}

//Get byte data from the URL
func getDataFromURL(link string) ([]byte, string, error) {
	//Get response from URL
	resp, err := http.Get(link)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	//Read the body and put it into a slice of byte
	byteData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	return byteData, resp.Request.URL.String(), nil
}

//Get all hrefs from URL
func getAllHrefFromURL(doc *html.Node) []string {
	//Slice of urls
	urls := []string{}
	//function to get all hrefs from URL
	var getHrefs func(*html.Node)
	getHrefs = func(n *html.Node) {
		//check if it is a element and a node
		if n.Type == html.ElementNode && n.Data == "a" {
			//check for hrefs`
			for _, a := range n.Attr {
				if a.Key == "href" {
					urls = append(urls, a.Val)
				}
			}
		}
		//Check for all children
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			getHrefs(c)
		}
	}
	getHrefs(doc)
	return urls
}

//Parse URls into gettables
func parseURLs(urls []string, baseURL string) []string {
	urlsParsed := []string{}
	for _, link := range urls {
		switch {
		case strings.HasPrefix(link, "/"):
			urlsParsed = append(urlsParsed, baseURL+link)
		case strings.HasPrefix(link, baseURL):
			urlsParsed = append(urlsParsed, link)
		}
	}
	return urlsParsed
}
