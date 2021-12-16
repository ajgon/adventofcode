package day02_test

import (
	"testing"

	"year2021/day02"
	"year2021/helpers"
)

func TestProcessInstruction(t *testing.T) {
	t.Parallel()

	gotCommand, gotValue, gotError := day02.ProcessInstruction("forward 33")
	wantCommand, wantValue := "forward", uint32(33)

	if gotCommand != wantCommand || gotValue != wantValue || gotError != nil {
		t.Errorf(
			"got [%s, %d, %v], want [%s, %d, %v]",
			gotCommand, gotValue, gotError,
			wantCommand, wantValue, nil,
		)
	}
}

func TestProcessInstructionsXY(t *testing.T) {
	t.Parallel()

	instructions := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	gotX, gotY := day02.ProcessInstructionsXY(instructions)
	wantX, wantY := uint32(15), uint32(10)

	if gotX != wantX || gotY != wantY {
		t.Errorf("got [%d, %d], want [%d, %d]", gotX, gotY, wantX, wantY)
	}
}

//nolint:ifshort
func TestSimpleSolution(t *testing.T) {
	t.Parallel()

	got := day02.SimpleSolution()
	want := helpers.Answers().Day02.Simple

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	t.Parallel()

	got := day02.AdvancedSolution()
	want := helpers.Answers().Day02.Advanced

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
