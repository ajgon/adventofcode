package day10

import (
	"reflect"
	"strings"
	"testing"
	"year2020/helpers"
)

var testData = strings.Split(
	`16
10
15
5
1
11
7
19
6
12
4`, "\n")

var otherTestData = strings.Split(
	`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`, "\n")

func TestFindDifferences01(t *testing.T) {
	got := findDifferences(testData)
	want := []int{1, 3, 1, 1, 1, 3, 1, 1, 3, 1, 3, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindDifferences02(t *testing.T) {
	got := findDifferences(otherTestData)
	want := []int{1, 1, 1, 1, 3, 1, 1, 1, 1, 3, 3, 1, 1, 1, 3, 1, 1, 3, 3, 1, 1, 1, 1, 3, 1, 3, 3, 1, 1, 1, 1, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCalculatePermutations01(t *testing.T) {
	got := calculatePermutations(testData)
	want := 8

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculatePermutations02(t *testing.T) {
	got := calculatePermutations(otherTestData)
	want := 19208

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day10.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day10.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
