package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := [][]string{
		{"1"},
		{"2"},
		{"1"},
		{"2"},
		{"1"},
		{"2"},
		{"1"},
		{"2"},
	}
	fmt.Println(arr)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(r)
	for n := len(arr); n > 0; n-- {
		randIndex := r.Intn(n)
		arr[n-1], arr[randIndex] = arr[randIndex], arr[n-1]
	}
	fmt.Println(arr)
}
