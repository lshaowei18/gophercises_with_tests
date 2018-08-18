package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Sleeper interface {
	Sleep()
}

func OpenCSV(fileName string) [][]string {
	//Open file for read access
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Can't open file", fileName)
	}

	reader := csv.NewReader(file) //New reader with CSV

	//Read all from file, gives a [][]string
	questions, err := reader.ReadAll()
	if err != nil {
		log.Fatal("err")
	}
	return questions
}

func InputReader(r io.Reader) string {
	scanner := bufio.NewScanner(r) //create a new scanner
	fmt.Print("Enter Text: ")      // Ask for input
	var input string
	//Scan for input and store it in a variable
	for scanner.Scan() {
		input = scanner.Text()
		break //break out of for loop the moment you get the input
	}
	return input
}

func AskAQuestionAndCheck(slice []string, r io.Reader, score *int) {
	//Ask user what is the answer
	fmt.Printf("What is %s?\n", slice[0])
	input := InputReader(r) //Get user answer
	//add to score if user is correct
	if input == slice[1] {
		*score++
	}
}

func CreateResults(length, score int) string {
	return fmt.Sprintf("Total number of questions: %d, total number of correct answers: %d", length, score)
}

func SleepAndTerminate(seconds int, sleeper Sleeper) {
	for i := seconds; i > 0; i-- {
		sleeper.Sleep()
	}
}
