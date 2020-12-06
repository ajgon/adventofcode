package day03

import (
	"strings"
	"testing"
	"year2020/helpers"
)

const testData = `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestCountTrees01(t *testing.T) {
	want := uint16(2)
	got := countTrees(strings.Split(strings.TrimSpace(testData), "\n"), 1, 1)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountTrees02(t *testing.T) {
	want := uint16(7)
	got := countTrees(strings.Split(strings.TrimSpace(testData), "\n"), 3, 1)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountTrees03(t *testing.T) {
	want := uint16(3)
	got := countTrees(strings.Split(strings.TrimSpace(testData), "\n"), 5, 1)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountTrees04(t *testing.T) {
	want := uint16(4)
	got := countTrees(strings.Split(strings.TrimSpace(testData), "\n"), 7, 1)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountTrees05(t *testing.T) {
	want := uint16(2)
	got := countTrees(strings.Split(strings.TrimSpace(testData), "\n"), 1, 2)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day03.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day03.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
