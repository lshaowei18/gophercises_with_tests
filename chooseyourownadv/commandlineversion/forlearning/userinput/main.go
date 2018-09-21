package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// num, err := UserInput(0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(num)

	num, err := UserInput(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}

func UserInput(numberOfOptions int) (int, error) {
	if numberOfOptions == 0 {
		return 0, fmt.Errorf("There is no options.")
	}
	userChoice := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && userChoice < 1 {
		text := scanner.Text()
		if num, err := strconv.Atoi(text); err == nil {
			if num > 0 && num <= numberOfOptions {
				userChoice = num
				break
			}
		}
	}
	return userChoice, nil
}
