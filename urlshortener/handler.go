package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-yaml/yaml"
)

func MapHandler(paths map[string]string, fallback http.Handler) http.HandlerFunc {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, path := range paths {
			if key == r.URL.Path {
				http.Redirect(w, r, path, http.StatusSeeOther)
			}
		}
		fallback.ServeHTTP(w, r)
	})
	return handlerFunc
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var returnErr error
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		datas := []Data{}
		err := yaml.Unmarshal(yml, &datas)
		if err != nil {
			returnErr = err
		}
		for _, data := range datas {
			if data.Path == r.URL.Path {
				http.Redirect(w, r, data.URL, http.StatusSeeOther)
			}
		}
		fallback.ServeHTTP(w, r)
	})
	return handlerFunc, returnErr
}

func JSONHandler(fallback http.Handler) (http.HandlerFunc, error) {
	var returnErr error
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Read json file
		dataByte, err := ioutil.ReadFile("paths.json")
		if err != nil {
			returnErr = err
		}
		//Unmarshal
		var datas []Data
		err = json.Unmarshal(dataByte, &datas)
		if err != nil {
			returnErr = err
		}
		//Handle paths
		for _, data := range datas {
			if data.Path == r.URL.Path {
				http.Redirect(w, r, data.URL, http.StatusSeeOther)
			}
		}
		fallback.ServeHTTP(w, r)
	})
	return handlerFunc, returnErr
}

func OpenYAMLFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

//function to create default mux
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
