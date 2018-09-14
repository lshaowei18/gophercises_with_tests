package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	removeSpace := func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}

	fmt.Println(strings.Map(removeSpace, "Hello bitches"))
}
