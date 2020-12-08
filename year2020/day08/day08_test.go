package day08

import (
	"reflect"
	"strings"
	"testing"
	"year2020/helpers"
)

var program = strings.Split(
	`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`, "\n")

func TestProcessAccPositive(t *testing.T) {
	cpu := CPU{}
	cpu.ProcessMnemonic("acc +5")
	got := cpu
	want := CPU{pc: 1, acc: 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestProcessAccNegative(t *testing.T) {
	cpu := CPU{}
	cpu.ProcessMnemonic("acc -8")
	got := cpu
	want := CPU{pc: 1, acc: -8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestProcessJmp(t *testing.T) {
	cpu := CPU{}
	cpu.ProcessMnemonic("jmp +7")
	got := cpu
	want := CPU{pc: 7, acc: 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestProcessNop(t *testing.T) {
	cpu := CPU{}
	cpu.ProcessMnemonic("nop +2")
	got := cpu
	want := CPU{pc: 1, acc: 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestRunUntilLoop(t *testing.T) {
	cpu := CPU{}
	cpu.RunUntilLoop(program)
	got := cpu
	want := CPU{pc: 1, acc: 5, history: []int{0, 1, 2, 6, 7, 3, 4}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day08.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day08.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
