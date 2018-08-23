package main

import (
	"fmt"
	"net/http"
)

func main() {
	mapHandler := MapHandler()
	http.Handle("/", mapHandler)
	http.ListenAndServe(":8080", nil)
}

func MapHandler() http.HandlerFunc {
	a := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello There")
	}
	return a
}
