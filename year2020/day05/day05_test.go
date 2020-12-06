package day05

import (
	"testing"
	"year2020/helpers"
)

func TestCalculateID01(t *testing.T) {
	want := 357
	got := calculateID("FBFBBFFRLR")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateID02(t *testing.T) {
	want := 567
	got := calculateID("BFFFBBFRRR")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateID03(t *testing.T) {
	want := 119
	got := calculateID("FFFBBBFRRR")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateID04(t *testing.T) {
	want := 820
	got := calculateID("BBFFBBFRLL")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day05.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day05.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
