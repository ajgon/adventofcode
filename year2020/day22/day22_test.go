package day22

import (
	"reflect"
	"strings"
	"testing"
	"year2020/helpers"
)

var testData = strings.Split(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`, "\n")

func TestExtractDecks(t *testing.T) {
	gotA, gotB := extractDecks(testData)
	wantA, wantB := []int{9, 2, 6, 3, 1}, []int{5, 8, 4, 7, 10}

	if !reflect.DeepEqual(gotA, wantA) || !reflect.DeepEqual(gotB, wantB) {
		t.Errorf("got %v/%v, want %v/%v", gotA, gotB, wantA, wantB)
	}
}

func TestIncludesDeckSet(t *testing.T) {
	gotA := includesDeckSet(
		[][][]int{
			[][]int{[]int{3, 8}, []int{6, 7, 3}},
			[][]int{[]int{5, 2}, []int{1, 2, 3}},
		}, []int{5, 2}, []int{1, 2, 3})
	gotB := includesDeckSet(
		[][][]int{
			[][]int{[]int{3, 8}, []int{6, 7, 3}},
			[][]int{[]int{3, 2}, []int{1, 2, 3}},
		}, []int{5, 2}, []int{1, 2, 3})
	wantA := true
	wantB := false

	if gotA != wantA || gotB != wantB {
		t.Errorf("got %t/%t, want %t/%t", gotA, gotB, wantA, wantB)
	}
}

func TestPlayRound(t *testing.T) {
	deckA, deckB := extractDecks(testData)

	gotA, gotB := playRound(deckA, deckB)
	wantA, wantB := []int{2, 6, 3, 1, 9, 5}, []int{8, 4, 7, 10}

	if !reflect.DeepEqual(gotA, wantA) || !reflect.DeepEqual(gotB, wantB) {
		t.Errorf("got %v/%v, want %v/%v", gotA, gotB, wantA, wantB)
	}

	gotC, gotD := playRound(gotA, gotB)
	wantC, wantD := []int{6, 3, 1, 9, 5}, []int{4, 7, 10, 8, 2}

	if !reflect.DeepEqual(gotC, wantC) || !reflect.DeepEqual(gotD, wantD) {
		t.Errorf("got %v/%v, want %v/%v", gotC, gotD, wantC, wantD)
	}
}

func TestPlaySimpleGame(t *testing.T) {
	deckA, deckB := extractDecks(testData)

	gotA, gotB := playGame(deckA, deckB, playRound)
	wantA, wantB := []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}, 'B'

	if !reflect.DeepEqual(gotA, wantA) && gotB != wantB {
		t.Errorf("got %v/%c, want %v/%c", gotA, wantA, gotB, wantB)
	}
}

func TestPlayAdvancedGame(t *testing.T) {
	deckA, deckB := extractDecks(testData)

	gotA, gotB := playGame(deckA, deckB, playRecursiveRound)
	wantA, wantB := []int{7, 5, 6, 2, 4, 1, 10, 8, 9, 3}, 'B'

	if !reflect.DeepEqual(gotA, wantA) && gotB != wantB {
		t.Errorf("got %v/%c, want %v/%c", gotA, wantA, gotB, wantB)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day22.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day22.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
