package main

import (
	"flag"
	"fmt"
	"os"
)

const fileName = "problems.csv"

var fileFlagPtr string

func init() {
	flag.StringVar(&fileFlagPtr, "filename", fileName, "CSV file for problems")
}
func main() {
	//Parse flag
	flag.Parse()
	//open file and get a [][]strings
	problems := OpenCSV(fileFlagPtr)
	//record score
	score := 0
	//Loop through each question and ask the user.
	for _, problem := range problems {
		AskAQuestionAndCheck(problem, os.Stdin, &score)
	}
	//Return results
	fmt.Println(CreateResults(len(problems), score))
}
