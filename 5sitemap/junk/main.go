package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	html "github.com/levigross/exp-html"
)

func main() {
	url := "http://tour.golang.org/welcome/1"
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
}

func getResponseBody(url string) ([]byte, error) {
	//Get the html page
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//Read the response body into a slice of byte
	byteData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return byteData, nil
}

func byteToHTML([]byte) (*html.Node, error) {
	doc, err := html.Parse(bytes.NewReader(htmlData)) // must convert to reader first
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func SliceUniqMap(s []int) []int {
	seen := make(map[int]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {

			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
