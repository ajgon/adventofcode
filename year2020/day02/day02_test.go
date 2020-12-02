package day02

import (
	"reflect"
	"testing"
	"year2020/helpers"
)

func TestExtractData(t *testing.T) {
	want := Data{low: 8, high: 11, letter: 'l', password: "qllllqllklhlvtl"}
	got := extractData("8-11 l: qllllqllklhlvtl")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestSimpleValidateLetterInRange(t *testing.T) {
	want := true
	got := simpleValidateData(Data{low: 8, high: 11, letter: 'l', password: "qllllqllklhlvtl"})

	if got != want {
		t.Errorf("data should be validated as valid, got invalid")
	}
}

func TestSimpleValidateLetterInRangeBottom(t *testing.T) {
	want := true
	got := simpleValidateData(Data{low: 8, high: 11, letter: 'l', password: "llllllll"})

	if got != want {
		t.Errorf("data should be validated as valid, got invalid")
	}
}

func TestSimpleValidateLetterInRangeTop(t *testing.T) {
	want := true
	got := simpleValidateData(Data{low: 8, high: 11, letter: 'l', password: "lllllllllll"})

	if got != want {
		t.Errorf("data should be validated as valid, got invalid")
	}
}

func TestSimpleValidateLetterOverRange(t *testing.T) {
	want := false
	got := simpleValidateData(Data{low: 8, high: 11, letter: 'l', password: "qllllqllklhlvtlllll"})

	if got != want {
		t.Errorf("data should be validated as invalid, got valid")
	}
}

func TestSimpleValidateLetterUnderRange(t *testing.T) {
	want := false
	got := simpleValidateData(Data{low: 8, high: 11, letter: 'l', password: "qlllsllv"})

	if got != want {
		t.Errorf("data should be validated as invalid, got valid")
	}
}

func TestAdvancedValidateLetter1(t *testing.T) {
	want := true
	got := advancedValidateData(Data{low: 1, high: 3, letter: 'a', password: "abcde"})

	if got != want {
		t.Errorf("data should be validated as valid, got invalid")
	}
}

func TestAdvancedValidateLetter2(t *testing.T) {
	want := false
	got := advancedValidateData(Data{low: 1, high: 3, letter: 'b', password: "cdefg"})

	if got != want {
		t.Errorf("data should be validated as invalid, got valid")
	}
}

func TestAdvancedValidateLetter3(t *testing.T) {
	want := false
	got := advancedValidateData(Data{low: 2, high: 9, letter: 'c', password: "ccccccccc"})

	if got != want {
		t.Errorf("data should be validated as invalid, got valid")
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day02.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day02.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
