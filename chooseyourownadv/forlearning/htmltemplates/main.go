package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl := template.Must(template.ParseFiles("tpl.html"))
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
