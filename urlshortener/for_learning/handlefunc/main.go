package main

//Follow this guide for now https://www.alexedwards.net/blog/making-and-using-middleware

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", middleware(messageHandler("/")))
	http.ListenAndServe(":8080", nil)
}

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Middleware logic goes here
		log.Println("Executing Middleware")
		next.ServeHTTP(w, r)
		log.Println("Executing middleware one again")
	})
}
