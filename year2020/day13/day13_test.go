package day13

import (
	"math/big"
	"reflect"
	"strings"
	"testing"
	"year2020/helpers"
)

var testData = strings.Split(`939
7,13,x,x,59,x,31,19`, "\n")

func TestExtractData(t *testing.T) {
	got1, got2 := extractData(testData)
	want1, want2 := 939, []int{7, 13, 0, 0, 59, 0, 31, 19}

	if got1 != want1 || !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %d/%v, want %d/%v", got1, got2, want1, want2)
	}
}

func TestFindEarliestBus(t *testing.T) {
	timestamp, data := extractData(testData)
	got1, got2 := findEarliestBus(timestamp, data)
	want1, want2 := 59, 944

	if got1 != want1 || got2 != want2 {
		t.Errorf("got %d/%d, want %d/%d", got1, got2, want1, want2)
	}
}

func TestCalculateChineseRemainder01(t *testing.T) {
	got := calculateChineseRemainder(
		[]*big.Int{
			big.NewInt(0),
			big.NewInt(12),
			big.NewInt(55),
			big.NewInt(25),
			big.NewInt(12),
		},
		[]*big.Int{
			big.NewInt(7),
			big.NewInt(13),
			big.NewInt(59),
			big.NewInt(31),
			big.NewInt(19),
		},
	)
	want := 1068781

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateChineseRemainder02(t *testing.T) {
	got := calculateChineseRemainder(
		[]*big.Int{
			big.NewInt(0),
			big.NewInt(11),
			big.NewInt(16),
		},
		[]*big.Int{
			big.NewInt(17),
			big.NewInt(13),
			big.NewInt(19),
		},
	)
	want := 3417

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateChineseRemainder03(t *testing.T) {
	got := calculateChineseRemainder(
		[]*big.Int{
			big.NewInt(0),
			big.NewInt(6),
			big.NewInt(57),
			big.NewInt(58),
		},
		[]*big.Int{
			big.NewInt(67),
			big.NewInt(7),
			big.NewInt(59),
			big.NewInt(61),
		},
	)
	want := 754018

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateChineseRemainder04(t *testing.T) {
	got := calculateChineseRemainder(
		[]*big.Int{
			big.NewInt(0),
			big.NewInt(5),
			big.NewInt(56),
			big.NewInt(57),
		},
		[]*big.Int{
			big.NewInt(67),
			big.NewInt(7),
			big.NewInt(59),
			big.NewInt(61),
		},
	)
	want := 779210

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateChineseRemainder05(t *testing.T) {
	got := calculateChineseRemainder(
		[]*big.Int{
			big.NewInt(0),
			big.NewInt(6),
			big.NewInt(56),
			big.NewInt(57),
		},
		[]*big.Int{
			big.NewInt(67),
			big.NewInt(7),
			big.NewInt(59),
			big.NewInt(61),
		},
	)
	want := 1261476

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateChineseRemainder06(t *testing.T) {
	got := calculateChineseRemainder(
		[]*big.Int{
			big.NewInt(0),
			big.NewInt(36),
			big.NewInt(45),
			big.NewInt(1886),
		},
		[]*big.Int{
			big.NewInt(1789),
			big.NewInt(37),
			big.NewInt(47),
			big.NewInt(1889),
		},
	)
	want := 1202161486

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPrepareRemainderData(t *testing.T) {
	_, data := extractData(testData)

	got := prepareRemainderData(data)
	want := [][]*big.Int{
		[]*big.Int{
			big.NewInt(0),
			big.NewInt(12),
			big.NewInt(55),
			big.NewInt(25),
			big.NewInt(12),
		},
		[]*big.Int{
			big.NewInt(7),
			big.NewInt(13),
			big.NewInt(59),
			big.NewInt(31),
			big.NewInt(19),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day13.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day13.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
