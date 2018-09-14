package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	//Open html file
	dataByte, err := ioutil.ReadFile("html/ex4.html")
	if err != nil {
		log.Fatalln(err)
	}
	//Parse the html byte data
	doc, err := html.Parse(bytes.NewReader(dataByte))
	if err != nil {
		log.Fatalln(err)
	}
	GetLinks(doc)
	links := GetLinks(doc)
	for _, link := range links {
		fmt.Printf("Href: %s\nText: %s \n", link.Href, link.Text)
	}
}

func GetLinks(doc *html.Node) []Link {
	//Get an array of links
	links := []Link{}
	var findHref func(*html.Node)
	//Write function to look through the html nodes
	findHref = func(n *html.Node) {
		//Look for a tags
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				//If it is a a href
				if a.Key == "href" {
					//Loop through its children and add the data into text
					var text string
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type != html.ElementNode && c.Type != html.CommentNode {
							text = text + strings.TrimSpace(c.Data)
						}
						//Children's children
						for d := c.FirstChild; d != nil; d = d.NextSibling {
							fmt.Println(d.Type)
							text = text + " " + d.Data
						}
					}
					link := Link{a.Val, text}
					//Append Href value to slice
					links = append(links, link)

				}
			}
		}
		// recursive to loop througgh every html element
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findHref(c)
		}
	}
	findHref(doc)
	return links
}
