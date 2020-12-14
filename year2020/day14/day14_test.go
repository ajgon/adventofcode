package day14

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"year2020/helpers"
)

var testSimpleData = strings.Split(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`, "\n")

var testAdvancedData = strings.Split(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`, "\n")

func TestApplyMask01(t *testing.T) {
	got := applyMask(11, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	want := 73

	if got != want {
		t.Errorf("got %b, want %b", got, want)
	}
}

func TestApplyMask02(t *testing.T) {
	got := applyMask(101, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	want := 101

	if got != want {
		t.Errorf("got %b, want %b", got, want)
	}
}

func TestApplyMask03(t *testing.T) {
	got := applyMask(0, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	want := 64

	if got != want {
		t.Errorf("got %b, want %b", got, want)
	}
}

func TestApplyFloatingMask01(t *testing.T) {
	got := applyFloatingMask(42, fmt.Sprintf("%036s", "X1001X"))
	want := []int{26, 27, 58, 59}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestApplyFloatingMask02(t *testing.T) {
	got := applyFloatingMask(26, fmt.Sprintf("%036s", "X0XX"))
	want := []int{16, 17, 18, 19, 24, 25, 26, 27}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestProcessProgramVersion1(t *testing.T) {
	got := processProgram(testSimpleData, 1)
	want := map[int]int{7: 101, 8: 64}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestProcessProgramVersion2(t *testing.T) {
	got := processProgram(testAdvancedData, 2)
	want := map[int]int{16: 1, 17: 1, 18: 1, 19: 1, 24: 1, 25: 1, 26: 1, 27: 1, 58: 100, 59: 100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day14.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day14.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
