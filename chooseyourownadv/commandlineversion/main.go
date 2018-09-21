package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//struct for the Chapters inside the story
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
		Link string
	} `json:"options"`
}

//List of keys in the story map
/*
new-york
debate
sean-kelly
mark-bates
denver
home
intro
*/

func main() {
	//Decode json
	dataBytes, err := ioutil.ReadFile("story.json")
	if err != nil {
		log.Fatalln(err)
	}
	//Umarshal JSON
	story := make(map[string]Chapter) // Map to hold all the chapters
	err = json.Unmarshal(dataBytes, &story)
	if err != nil {
		log.Fatalln(err)
	}
	//Print out intro story
	PrintChapter("intro", story)
}

//Get the text of the by chapter name which is the key to the map
func PrintChapter(name string, story map[string]Chapter) {
	chapter := story[name]
	//Loop through story and print out each item
	for _, story := range chapter.Story {
		fmt.Printf("\n%s\n", story)
	}
	//Print out option text
	fmt.Printf("\nYou have %d option(s):\n", len(chapter.Options))
	for idx, option := range chapter.Options {
		fmt.Printf("\n%d) %s\n", idx+1, option.Text)
	}

	//Ask for user input
	input, err := UserInput(len(chapter.Options))
	if err == nil {
		PrintChapter(chapter.Options[input-1].Arc, story)
	}
}

//Ask for the user input, takes in length of options
func UserInput(numberOfOptions int) (int, error) {
	//Check for zero options
	if numberOfOptions == 0 {
		return 0, fmt.Errorf("There is no options.")
	}
	//Get user input
	userChoice := 0
	scanner := bufio.NewScanner(os.Stdin)
	//Scan till there is valid input
	fmt.Printf("Please key in a number that reflects your choice!\n")
	for scanner.Scan() && userChoice < 1 {
		text := scanner.Text()
		//Make sure input is a number
		if num, err := strconv.Atoi(text); err == nil {
			//Number must be more than 0 and less than the total number of inputs.
			if num > 0 && num <= numberOfOptions {
				userChoice = num
				break
			}
		}
		fmt.Printf("Please enter a valid input!\n")
	}
	return userChoice, nil
}
