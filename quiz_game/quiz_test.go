package main

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"strings"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestOpenCSV(t *testing.T) {
	in := `5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7`
	got := OpenCSV("problems.csv")
	reader := csv.NewReader(strings.NewReader(in))
	want, _ := reader.ReadAll()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInputReader(t *testing.T) {
	buf := bytes.NewBufferString("Hello")
	got := InputReader(buf)
	want := "Hello"
	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAskAQuestionAndCheck(t *testing.T) {
	slice := []string{"7+3", "10"}
	got := 0
	AskAQuestionAndCheck(slice, strings.NewReader("10"), &got)
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateResults(t *testing.T) {
	got := CreateResults(10, 10)
	want := "Total number of questions: 10, total number of correct answers: 10"
	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSleepAndTerminate(t *testing.T) {
	spySleeper := &SpySleeper{}
	SleepAndTerminate(30, spySleeper)
	got := spySleeper.Calls
	want := 30
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestShuffleSlice(t *testing.T) {
	slice := [][]string{
		{"1"},
		{"2"},
		{"1"},
		{"2"},
		{"1"},
		{"2"},
		{"1"},
		{"2"},
	}
	var want [][]string
	copy(want, slice)
	ShuffleSlice(slice)
	if reflect.DeepEqual(want, slice) {
		t.Error("Both slices should not be the same.")
	}
}
