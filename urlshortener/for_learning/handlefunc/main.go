package main

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog dog")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", d)
	fmt.Println(mux)
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", d)
	fmt.Println(mux1)
	x := reflect.ValueOf(mux)
	y := reflect.ValueOf(mux1)
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Println(reflect.DeepEqual(x, y))
	http.ListenAndServe(":8080", mux)
}
