package day09

import (
	"strings"
	"testing"
	"year2020/helpers"
)

var testData = strings.Split(
	`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`, "\n")

func TestContainsSum(t *testing.T) {
	got := containsSum([]int{8, 14, 25}, 39)
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestDoesNotContainSum(t *testing.T) {
	got := containsSum([]int{8, 14, 25}, 29)
	want := false

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestFindMissingNumber(t *testing.T) {
	got := findMissingNumber(testData, 5)
	want := 127

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindEncryptionWeakness(t *testing.T) {
	got := findEncryptionWeakness(testData, 127)
	want := 62

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day09.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSimpleAdvanced(t *testing.T) {
	want := helpers.Answers.Day09.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
