package day06

import (
	"testing"
	"year2020/helpers"
)

func TestSimpleCountAnswers01(t *testing.T) {
	want := 3
	got := simpleCountAnswers("abc")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleCountAnswers02(t *testing.T) {
	want := 3
	got := simpleCountAnswers("a\nb\nc")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleCountAnswers03(t *testing.T) {
	want := 3
	got := simpleCountAnswers("ab\nac")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleCountAnswers04(t *testing.T) {
	want := 1
	got := simpleCountAnswers("a\na\na\na")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleCountAnswers05(t *testing.T) {
	want := 1
	got := simpleCountAnswers("b")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAdvancedCountAnswers01(t *testing.T) {
	want := 3
	got := advancedCountAnswers("abc")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAdvancedCountAnswers02(t *testing.T) {
	want := 0
	got := advancedCountAnswers("a\nb\nc")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAdvancedCountAnswers03(t *testing.T) {
	want := 1
	got := advancedCountAnswers("ab\nac")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAdvancedCountAnswers04(t *testing.T) {
	want := 1
	got := advancedCountAnswers("a\na\na\na")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAdvancedCountAnswers05(t *testing.T) {
	want := 1
	got := advancedCountAnswers("b")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day06.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day06.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
