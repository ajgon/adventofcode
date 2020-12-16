package day15

import (
	"testing"
	"year2020/helpers"
)

func TestCalculateNthIterationSimple01(t *testing.T) {
	got := calculateNthIteration([]int{0, 3, 6}, 2020)
	want := 436

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationSimple02(t *testing.T) {
	got := calculateNthIteration([]int{1, 3, 2}, 2020)
	want := 1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationSimple03(t *testing.T) {
	got := calculateNthIteration([]int{2, 1, 3}, 2020)
	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationSimple04(t *testing.T) {
	got := calculateNthIteration([]int{1, 2, 3}, 2020)
	want := 27

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationSimple05(t *testing.T) {
	got := calculateNthIteration([]int{2, 3, 1}, 2020)
	want := 78

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationSimple06(t *testing.T) {
	got := calculateNthIteration([]int{3, 2, 1}, 2020)
	want := 438

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationSimple07(t *testing.T) {
	got := calculateNthIteration([]int{3, 1, 2}, 2020)
	want := 1836

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationAdvanced01(t *testing.T) {
	got := calculateNthIteration([]int{0, 3, 6}, 30000000)
	want := 175594

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationAdvanced02(t *testing.T) {
	got := calculateNthIteration([]int{1, 3, 2}, 30000000)
	want := 2578

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationAdvanced03(t *testing.T) {
	got := calculateNthIteration([]int{2, 1, 3}, 30000000)
	want := 3544142

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationAdvanced04(t *testing.T) {
	got := calculateNthIteration([]int{1, 2, 3}, 30000000)
	want := 261214

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationAdvanced05(t *testing.T) {
	got := calculateNthIteration([]int{2, 3, 1}, 30000000)
	want := 6895259

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationAdvanced06(t *testing.T) {
	got := calculateNthIteration([]int{3, 2, 1}, 30000000)
	want := 18

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNthIterationAdvanced07(t *testing.T) {
	got := calculateNthIteration([]int{3, 1, 2}, 30000000)
	want := 362

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day15.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day15.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
