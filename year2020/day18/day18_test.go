package day18

import (
	"testing"
	"year2020/helpers"
)

func TestCalculateSimple(t *testing.T) {
	got := calculateSimple("12 + 23 * 34")
	want := "1190"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCalculateSimpleInBrackets01(t *testing.T) {
	got := calculateInBrackets("1 + (2 * 3) + (4 * (5 + 6))", calculateSimple)
	want := 51

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateSimpleInBrackets02(t *testing.T) {
	got := calculateInBrackets("2 * 3 + (4 * 5)", calculateSimple)
	want := 26

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateSimpleInBrackets03(t *testing.T) {
	got := calculateInBrackets("5 + (8 * 3 + 9 + 3 * 4 * 3)", calculateSimple)
	want := 437

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateSimpleInBrackets04(t *testing.T) {
	got := calculateInBrackets("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", calculateSimple)
	want := 12240

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateSimpleInBrackets05(t *testing.T) {
	got := calculateInBrackets("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", calculateSimple)
	want := 13632

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateSumsOnly(t *testing.T) {
	got := calculateSumsOnly("5 * 3 + 2 + 5 * 7 + 3 + 9")
	want := "5 * 10 * 19"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCalculateAdvanced(t *testing.T) {
	got := calculateAdvanced("5 * 3 + 2 + 5 * 7 + 3 + 9")
	want := "950"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCalculateAdvancedInBrackets01(t *testing.T) {
	got := calculateInBrackets("1 + (2 * 3) + (4 * (5 + 6))", calculateAdvanced)
	want := 51

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateAdvancedInBrackets02(t *testing.T) {
	got := calculateInBrackets("2 * 3 + (4 * 5)", calculateAdvanced)
	want := 46

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateAdvancedInBrackets03(t *testing.T) {
	got := calculateInBrackets("5 + (8 * 3 + 9 + 3 * 4 * 3)", calculateAdvanced)
	want := 1445

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateAdvancedInBrackets04(t *testing.T) {
	got := calculateInBrackets("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", calculateAdvanced)
	want := 669060

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateAdvancedInBrackets05(t *testing.T) {
	got := calculateInBrackets("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", calculateAdvanced)
	want := 23340

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day18.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day18.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
