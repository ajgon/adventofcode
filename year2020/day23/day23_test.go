package day23

import (
	"reflect"
	"testing"
	"year2020/helpers"
)

func TestPickCups(t *testing.T) {
	boardA := []int{1, 2, 3, 4, 5}
	boardB := []int{1, 2, 3, 4, 5}
	gotA := pickCups(&boardA, 1)
	gotB := pickCups(&boardB, 3)
	wantA := []int{2, 3, 4}
	wantB := []int{4, 5, 1}
	wantBoardA := []int{1, 5}
	wantBoardB := []int{2, 3}

	if !reflect.DeepEqual(gotA, wantA) || !reflect.DeepEqual(gotB, wantB) {
		t.Errorf("got %v, %v - want %v, %v", gotA, gotB, wantA, wantB)
	}

	if !reflect.DeepEqual(boardA, wantBoardA) || !reflect.DeepEqual(boardB, wantBoardB) {
		t.Errorf("got %v, %v - want %v, %v", boardA, boardB, wantBoardA, wantBoardB)
	}
}

func TestPlaceCups(t *testing.T) {
	got := []int{1, 5}
	placeCups(&got, []int{2, 3, 4}, 1)
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v - want %v", got, want)
	}
}

func TestPlayMove(t *testing.T) {
	gotA, gotB := playMove([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 3)
	wantA, wantB := []int{3, 2, 8, 9, 1, 5, 4, 6, 7}, 2

	if !reflect.DeepEqual(gotA, wantA) || gotB != wantB {
		t.Errorf("got %v, %d - want %v, %d", gotA, gotB, wantA, wantB)
	}
}

func TestExampleGame(t *testing.T) {
	gotA1, gotB1 := playMove([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 3)
	gotA2, gotB2 := playMove(gotA1, gotB1)
	gotA3, gotB3 := playMove(gotA2, gotB2)
	gotA4, gotB4 := playMove(gotA3, gotB3)
	gotA5, gotB5 := playMove(gotA4, gotB4)
	gotA6, gotB6 := playMove(gotA5, gotB5)
	gotA7, gotB7 := playMove(gotA6, gotB6)
	gotA8, gotB8 := playMove(gotA7, gotB7)
	gotA9, gotB9 := playMove(gotA8, gotB8)
	gotA10, gotB10 := playMove(gotA9, gotB9)
	wantA1, wantB1 := []int{3, 2, 8, 9, 1, 5, 4, 6, 7}, 2
	wantA2, wantB2 := []int{3, 2, 5, 4, 6, 7, 8, 9, 1}, 5
	wantA3, wantB3 := []int{3, 4, 6, 7, 2, 5, 8, 9, 1}, 8
	wantA4, wantB4 := []int{4, 6, 7, 9, 1, 3, 2, 5, 8}, 4
	wantA5, wantB5 := []int{4, 1, 3, 6, 7, 9, 2, 5, 8}, 1
	wantA6, wantB6 := []int{4, 1, 9, 3, 6, 7, 2, 5, 8}, 9
	wantA7, wantB7 := []int{4, 1, 9, 2, 5, 8, 3, 6, 7}, 2
	wantA8, wantB8 := []int{4, 1, 5, 8, 3, 9, 2, 6, 7}, 6
	wantA9, wantB9 := []int{5, 7, 4, 1, 8, 3, 9, 2, 6}, 5
	wantA10, wantB10 := []int{5, 8, 3, 7, 4, 1, 9, 2, 6}, 8

	if !reflect.DeepEqual(gotA1, wantA1) || gotB1 != wantB1 {
		t.Errorf("got %v, %d - want %v, %d", gotA1, gotB1, wantA1, wantB1)
	}

	if !reflect.DeepEqual(gotA2, wantA2) || gotB2 != wantB2 {
		t.Errorf("got %v, %d - want %v, %d", gotA2, gotB2, wantA2, wantB2)
	}

	if !reflect.DeepEqual(gotA3, wantA3) || gotB3 != wantB3 {
		t.Errorf("got %v, %d - want %v, %d", gotA3, gotB3, wantA3, wantB3)
		panic("X")
	}

	if !reflect.DeepEqual(gotA4, wantA4) || gotB4 != wantB4 {
		t.Errorf("got %v, %d - want %v, %d", gotA4, gotB4, wantA4, wantB4)
	}

	if !reflect.DeepEqual(gotA5, wantA5) || gotB5 != wantB5 {
		t.Errorf("got %v, %d - want %v, %d", gotA5, gotB5, wantA5, wantB5)
	}

	if !reflect.DeepEqual(gotA6, wantA6) || gotB6 != wantB6 {
		t.Errorf("got %v, %d - want %v, %d", gotA6, gotB6, wantA6, wantB6)
	}

	if !reflect.DeepEqual(gotA7, wantA7) || gotB7 != wantB7 {
		t.Errorf("got %v, %d - want %v, %d", gotA7, gotB7, wantA7, wantB7)
	}

	if !reflect.DeepEqual(gotA8, wantA8) || gotB8 != wantB8 {
		t.Errorf("got %v, %d - want %v, %d", gotA8, gotB8, wantA8, wantB8)
	}

	if !reflect.DeepEqual(gotA9, wantA9) || gotB9 != wantB9 {
		t.Errorf("got %v, %d - want %v, %d", gotA9, gotB9, wantA9, wantB9)
	}

	if !reflect.DeepEqual(gotA10, wantA10) || gotB10 != wantB10 {
		t.Errorf("got %v, %d - want %v, %d", gotA10, gotB10, wantA10, wantB10)
	}
}

func TestPlayGames(t *testing.T) {
	got := playGames([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 10)
	want := []int{5, 8, 3, 7, 4, 1, 9, 2, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBuildCupList(t *testing.T) {
	cup, cupMap := buildCupList([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 15)

	got := make([]int, 0)

	for i := 0; i < 20; i++ {
		got = append(got, cup.Value)
		cup = cup.Next
	}
	want := []int{3, 8, 9, 1, 2, 5, 4, 6, 7, 10, 11, 12, 13, 14, 15, 3, 8, 9, 1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	gotMap1 := cupMap[9].Value
	wantMap1 := 9

	gotMap2 := cupMap[13].Value
	wantMap2 := 13

	if gotMap1 != wantMap1 || gotMap2 != wantMap2 {
		t.Errorf("got %d/%d, want %d/%d", gotMap1, gotMap2, wantMap1, wantMap2)
	}
}

func TestPlayLinkedGames(t *testing.T) {
	firstCup, cupMap := buildCupList([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 1000000)

	cup := playLinkedGames(firstCup, cupMap, 10000000)

	for {
		if cup.Value == 1 {
			break
		}
		cup = cup.Next
	}

	got := []int{cup.Next.Value, cup.Next.Next.Value}
	want := []int{934001, 159792}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day23.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day23.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
