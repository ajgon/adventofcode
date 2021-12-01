package day01_test

import (
	"testing"

	"year2021/day01"
	"year2021/helpers"
)

func TestFindIncreasingEntriesCount(t *testing.T) {
	t.Parallel()

	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got, _ := day01.FindIncreasingEntriesCount(data, 1)
	want := 7

	if got != uint16(want) {
		t.Errorf("got %d, want %d", got, want)
	}

	got, _ = day01.FindIncreasingEntriesCount(data, 3)
	want = 5

	if got != uint16(want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

//nolint:ifshort
func TestSimpleSolution(t *testing.T) {
	t.Parallel()

	got := day01.SimpleSolution()
	want := helpers.Answers().Day01.Simple

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	t.Parallel()

	got := day01.AdvancedSolution()
	want := helpers.Answers().Day01.Advanced

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
