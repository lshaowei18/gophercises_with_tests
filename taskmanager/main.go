package main

import (
	"log"
	"shaowei/gophercises_with_tests/taskmanager/cmd"
	"shaowei/gophercises_with_tests/taskmanager/db"
)

func main() {
	err := db.SetupDB("test.db")
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Execute()
}
