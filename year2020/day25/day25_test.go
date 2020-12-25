package day25

import (
	"testing"
	"year2020/helpers"
)

func TestRotate(t *testing.T) {
	got := rotate(20152380, 7)
	want := 19859298

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindLoopSize(t *testing.T) {
	got1, got2 := findLoopSize(5764801, 7), findLoopSize(17807724, 7)
	want1, want2 := 8, 11

	if got1 != want1 || got2 != want2 {
		t.Errorf("got %d/%d, want %d/%d", got1, got2, want1, want2)
	}
}

func TestFindEncryptionKey(t *testing.T) {
	got1, got2 := findEncrypionKey(5764801, 11), findEncrypionKey(17807724, 8)
	want1, want2 := 14897079, 14897079

	if got1 != want1 || got2 != want2 {
		t.Errorf("got %d/%d, want %d/%d", got1, got2, want1, want2)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day25.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
