package day01

import (
	"testing"
	"year2020/helpers"
)

func TestFindNEntriesWithSum(t *testing.T) {
	var want uint16 = 2020
	got, _ := findNEntriesWithSum(2, want)

	if got[0]+got[1] != want {
		t.Errorf("got %d and %d which sum up to %d, want %d", got[0], got[1], got[0]+got[1], want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day01.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day01.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
