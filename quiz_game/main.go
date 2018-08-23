package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	// "os/signal"
	// "syscall"
)

const fileName = "problems.csv"

var fileFlagPtr string
var shuffleFlag bool

func init() {
	flag.StringVar(&fileFlagPtr, "filename", fileName, "CSV file for problems")
	flag.BoolVar(&shuffleFlag, "shuffle", false, "Option to shuffle questions")
}

func main() {
	//Parse flag
	flag.Parse()
	//open file and get a [][]strings
	problems := OpenCSV(fileFlagPtr)
	//Shuffle if user wants it
	if shuffleFlag {
		ShuffleSlice(problems)
	}
	//record score
	score := 0
	//Ask user if he is ready to started
	fmt.Println("Press Enter when you are ready to start the timer.")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	//Create Channel to accept signal and end Loop
	quitChannel := make(chan os.Signal, 1)
	go func() {
		for {
			signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
		}
	}()
	//start timer
	timer1 := time.NewTimer(5 * time.Second)
	//Loop through each question and ask the user.
	go func() {
		for _, problem := range problems {
			AskAQuestionAndCheck(problem, os.Stdin, &score)
		}
	}()
	select {
	case _ = <-quitChannel:

	case _ = <-timer1.C:
	}
	// <-quitChannel
	//Return results
	fmt.Println(CreateResults(len(problems), score))
}
