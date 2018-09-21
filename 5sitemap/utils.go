package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

//Get the url without the path
func getBaseURL(rawURL string) (string, error) {
	path, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return rawURL[:len(rawURL)-len(path.Path)], nil
}

//Get byte data from the URL
func getDataFromURL(link string) ([]byte, error) {
	//Get response from URL
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//Read the body and put it into a slice of byte
	byteData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteData, nil
}

//Parse response body into html node
func parseURLIntoHTMLNode(data []byte) (*html.Node, error) {
	htmlNode, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return htmlNode, nil
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

//Filter urls
func filterURLs(urls []string, baseURL string) []string {
	//remove empty
	var filterFunc func(string) bool
	filterFunc = func(url string) bool {
		//check empty
		if url == "/" {
			return false
		}
		return true
	}
}

//Parse URls into gettables
func parseURLs(urls []string, baseURL string) []string {
	urlsParsed := []string{}
	for _, link := range urls {
		//Check for empty URLs and ones without domain name
		if link[0] == "/" && len(link) > 1 {
			newURL := baseURL + link
			urlsParsed = append(urlsParsed, newURL)
		}
		//Check if it is valid URL and contains baseURL
		if _, err := url.Parse(link); err != nil && strings.Contains(link, baseURL) {
			urlsParsed = append(urlsParsed, newURL)
		}
	}
	return urlsParsed
}
